package fetchers

import (
	"context"
	"strconv"
	"time"

	"code.cloudfoundry.org/cpu-entitlement-plugin/cf"
	logcache "code.cloudfoundry.org/go-log-cache"
	"code.cloudfoundry.org/go-log-cache/rpc/logcache_v1"
	"code.cloudfoundry.org/go-loggregator/v9/rpc/loggregator_v2"
	"code.cloudfoundry.org/lager"
)

type LastSpikeInstanceData struct {
	InstanceID int
	From       time.Time
	To         time.Time
}

type LastSpikeFetcher struct {
	client LogCacheClient
	since  time.Time
}

func NewLastSpikeFetcher(client LogCacheClient, since time.Time) *LastSpikeFetcher {
	return &LastSpikeFetcher{client: client, since: since}
}

func (f LastSpikeFetcher) FetchInstanceData(logger lager.Logger, appGUID string, appInstances map[int]cf.Instance) (map[int]interface{}, error) {
	logger = logger.Session("last-spike-fetcher", lager.Data{"app-guid": appGUID})
	logger.Info("start")
	defer logger.Info("end")

	res, err := f.client.Read(context.Background(), appGUID, f.since,
		logcache.WithEnvelopeTypes(logcache_v1.EnvelopeType_TIMER),
		logcache.WithDescending(),
		logcache.WithNameFilter("spike"),
	)
	if err != nil {
		logger.Error("logcache-client-read-failed", err)
		return nil, err
	}

	return parseLastSpike(logger, res, appInstances)
}

func parseLastSpike(logger lager.Logger, res []*loggregator_v2.Envelope, appInstances map[int]cf.Instance) (map[int]interface{}, error) {
	logger = logger.Session("parse-last-spike")
	logger.Info("start")
	defer logger.Info("end")

	lastSpikePerInstance := make(map[int]interface{})
	for _, envelope := range res {
		instanceID, err := strconv.Atoi(envelope.InstanceId)
		if err != nil {
			logger.Info("ignoring-corrupt-instance-id", lager.Data{"instance-id": envelope.InstanceId, "envelope": envelope})
			continue
		}

		if _, alreadySet := lastSpikePerInstance[instanceID]; alreadySet {
			logger.Info("already-set", lager.Data{"instanceID": instanceID, "lastSpike": lastSpikePerInstance[instanceID]})
			continue
		}

		envelopeGauge, ok := envelope.Message.(*loggregator_v2.Envelope_Timer)
		if !ok {
			logger.Info("ignoring-non-gauge-message", lager.Data{"gauge-message": envelope.Message})
			continue
		}

		logger.Info("Tags: ", lager.Data{"tags": envelope.Tags, "ProcessInstanceID": appInstances[instanceID].ProcessInstanceID})

		processInstanceID := envelope.Tags["process_instance_id"]
		if appInstances[instanceID].ProcessInstanceID != processInstanceID {
			continue
		}

		if envelopeGauge.Timer != nil {
			lastSpikePerInstance[instanceID] = LastSpikeInstanceData{
				InstanceID: instanceID,
				From:       time.Unix(0, envelopeGauge.Timer.Start),
				To:         time.Unix(0, envelopeGauge.Timer.Stop),
			}
		}
	}

	return lastSpikePerInstance, nil
}
