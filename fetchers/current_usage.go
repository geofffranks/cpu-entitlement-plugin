package fetchers

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"code.cloudfoundry.org/cpu-entitlement-plugin/cf"
	"code.cloudfoundry.org/go-log-cache/rpc/logcache_v1"
	"code.cloudfoundry.org/lager"
)

type CurrentInstanceData struct {
	InstanceID int
	Usage      float64
}

//go:generate counterfeiter . Fetcher
type Fetcher interface {
	FetchInstanceData(logger lager.Logger, appGUID string, appInstances map[int]cf.Instance) (map[int]interface{}, error)
}

type CurrentUsageFetcher struct {
	client          LogCacheClient
	fallbackFetcher Fetcher
}

func NewCurrentUsageFetcher(client LogCacheClient) CurrentUsageFetcher {
	return CurrentUsageFetcher{
		client:          client,
		fallbackFetcher: NewCumulativeUsageFetcher(client),
	}
}

func NewCurrentUsageFetcherWithFallbackFetcher(client LogCacheClient, fallbackFetcher Fetcher) CurrentUsageFetcher {
	return CurrentUsageFetcher{
		client:          client,
		fallbackFetcher: fallbackFetcher,
	}
}

func (f CurrentUsageFetcher) FetchInstanceData(logger lager.Logger, appGUID string, appInstances map[int]cf.Instance) (map[int]interface{}, error) {
	logger = logger.Session("current-usage-fetcher", lager.Data{"app-guid": appGUID})
	logger.Info("start")
	defer logger.Info("end")

	query := fmt.Sprintf(`idelta(absolute_usage{source_id="%s"}[1m]) / idelta(absolute_entitlement{source_id="%s"}[1m])`, appGUID, appGUID)
	logger.Info("Query string ", lager.Data{"query": query})
	res, err := f.client.PromQL(context.Background(), query)
	if err != nil {
		logger.Error("promql-failed", err, lager.Data{"query": query})
		return nil, err
	}
	logger.Info("Query result ", lager.Data{"result": res})

	currentUsage := parseCurrentUsage(logger, res, appInstances)
	logger.Info("currentUsage ", lager.Data{"usage": currentUsage})
	if len(currentUsage) == len(appInstances) {
		return currentUsage, nil
	}

	logger.Info("falling-back-to-cumulative-fetcher")

	cumulativeResult, err := f.fallbackFetcher.FetchInstanceData(logger, appGUID, appInstances)
	logger.Info("cumulativeResult ", lager.Data{"cumulativeResult": cumulativeResult})
	if err != nil {
		logger.Info("fallback-fetcher-failed")
		return nil, err
	}

	for instanceID := range appInstances {
		if _, ok := currentUsage[instanceID]; ok {
			continue
		}

		result, ok := cumulativeResult[instanceID]
		if !ok {
			continue
		}

		cumulativeData, ok := result.(CumulativeInstanceData)
		if !ok {
			err = errors.New("cumulative fetcher returned result in unexpected struct")
			logger.Error("unexpected-type-from-fetcher", err, lager.Data{"result": result})
			return map[int]interface{}{}, err
		}

		currentUsage[instanceID] = CurrentInstanceData{
			InstanceID: cumulativeData.InstanceID,
			Usage:      cumulativeData.Usage,
		}
		logger.Info("currentUsage ", lager.Data{"currentUsage": currentUsage})
	}

	return currentUsage, nil
}

func parseCurrentUsage(logger lager.Logger, res *logcache_v1.PromQL_InstantQueryResult, appInstances map[int]cf.Instance) map[int]interface{} {
	usagePerInstance := make(map[int]interface{})
	for _, sample := range res.GetVector().GetSamples() {
		instanceID, err := strconv.Atoi(sample.GetMetric()["instance_id"])
		if err != nil {
			logger.Info("ignoring-corrupt-instance-id", lager.Data{"instance-id": sample.GetMetric()["instance_id"]})
			continue
		}

		processInstanceID := sample.GetMetric()["process_instance_id"]
		if processInstanceID != appInstances[instanceID].ProcessInstanceID {
			continue
		}

		dataPoint := CurrentInstanceData{
			InstanceID: instanceID,
			Usage:      sample.GetPoint().GetValue(),
		}
		usagePerInstance[instanceID] = dataPoint
	}

	return usagePerInstance
}
