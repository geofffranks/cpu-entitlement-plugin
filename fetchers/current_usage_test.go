package fetchers_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"code.cloudfoundry.org/cpu-entitlement-plugin/fetchers"
	"code.cloudfoundry.org/cpu-entitlement-plugin/fetchers/fetchersfakes"
	"code.cloudfoundry.org/cpu-entitlement-plugin/metadata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CurrentUsage", func() {
	var (
		logCacheClient *fetchersfakes.FakeLogCacheClient
		fetcher        *fetchers.CurrentUsageFetcher
		appGuid        string
		appInstances   map[int]metadata.CFAppInstance
		currentUsage   map[int][]fetchers.InstanceData
		fetchErr       error
	)

	BeforeEach(func() {
		logCacheClient = new(fetchersfakes.FakeLogCacheClient)
		fetcher = fetchers.NewCurrentUsageFetcher(logCacheClient)

		appGuid = "foo"

		logCacheClient.PromQLReturns(queryResult(
			sample("0", "abc",
				point("1", 0.2),
			),
			sample("1", "def",
				point("2", 0.4),
			),
			sample("2", "ghi",
				point("4", 0.5),
			),
		), nil)
		appInstances = map[int]metadata.CFAppInstance{
			0: metadata.CFAppInstance{InstanceID: 0, Since: time.Unix(0, 0)},
			1: metadata.CFAppInstance{InstanceID: 1, Since: time.Unix(0, 0)},
			2: metadata.CFAppInstance{InstanceID: 2, Since: time.Unix(0, 0)},
		}
	})

	JustBeforeEach(func() {
		currentUsage, fetchErr = fetcher.FetchInstanceData(appGuid, appInstances)
	})

	It("gets the current usage from the log-cache client", func() {
		Expect(logCacheClient.PromQLCallCount()).To(Equal(1))
		ctx, query, _ := logCacheClient.PromQLArgsForCall(0)
		Expect(ctx).To(Equal(context.Background()))
		Expect(query).To(Equal(fmt.Sprintf(`idelta(absolute_usage{source_id="%s"}[1m]) / idelta(absolute_entitlement{source_id="%s"}[1m])`, appGuid, appGuid)))
	})

	It("returns the correct current usage", func() {
		Expect(fetchErr).NotTo(HaveOccurred())
		Expect(currentUsage).To(Equal(map[int][]fetchers.InstanceData{
			0: {
				{
					Time:       time.Unix(1, 0),
					InstanceID: 0,
					Value:      0.2,
				},
			},
			1: {
				{
					Time:       time.Unix(2, 0),
					InstanceID: 1,
					Value:      0.4,
				},
			},
			2: {
				{
					Time:       time.Unix(4, 0),
					InstanceID: 2,
					Value:      0.5,
				},
			},
		}))
	})

	When("cache returns data for instances that are no longer running (because the app has been scaled down", func() {
		BeforeEach(func() {
			appInstances = map[int]metadata.CFAppInstance{
				0: metadata.CFAppInstance{InstanceID: 0, Since: time.Unix(0, 0)},
			}
		})

		It("returns current usage for running instances only", func() {
			Expect(fetchErr).NotTo(HaveOccurred())
			Expect(currentUsage).To(Equal(map[int][]fetchers.InstanceData{
				0: {
					{
						Time:       time.Unix(1, 0),
						InstanceID: 0,
						Value:      0.2,
					},
				},
			}))
		})
	})

	When("cache returns data before the instance has been created", func() {
		BeforeEach(func() {
			appInstances = map[int]metadata.CFAppInstance{
				0: metadata.CFAppInstance{InstanceID: 0, Since: time.Unix(2, 0)},
				1: metadata.CFAppInstance{InstanceID: 1, Since: time.Unix(2, 0)},
				2: metadata.CFAppInstance{InstanceID: 2, Since: time.Unix(2, 0)},
			}
		})

		It("filters out data before the instance has been created", func() {
			Expect(fetchErr).NotTo(HaveOccurred())
			Expect(currentUsage).To(Equal(map[int][]fetchers.InstanceData{
				1: {
					{
						Time:       time.Unix(2, 0),
						InstanceID: 1,
						Value:      0.4,
					},
				},
				2: {
					{
						Time:       time.Unix(4, 0),
						InstanceID: 2,
						Value:      0.5,
					},
				},
			}))
		})
	})

	When("fetching the current usage fails", func() {
		BeforeEach(func() {
			logCacheClient.PromQLReturns(nil, errors.New("fetch-failed"))
		})

		It("returns the error", func() {
			Expect(fetchErr).To(MatchError("fetch-failed"))
		})
	})

	When("fetched data has corrupt timestamp", func() {
		BeforeEach(func() {
			logCacheClient.PromQLReturns(queryResult(
				sample("0", "abc",
					point("1", 0.2),
				),
				sample("1", "def",
					point("baba", 0.4),
				),
				sample("2", "ghi",
					point("4", 0.5),
				),
			), nil)
		})

		It("ignores the corrupt data point", func() {
			Expect(fetchErr).NotTo(HaveOccurred())
			Expect(currentUsage).To(Equal(map[int][]fetchers.InstanceData{
				0: {
					{
						Time:       time.Unix(1, 0),
						InstanceID: 0,
						Value:      0.2,
					},
				},
				2: {
					{
						Time:       time.Unix(4, 0),
						InstanceID: 2,
						Value:      0.5,
					},
				},
			}))
		})
	})

	When("fetched data has corrupt instance id", func() {
		BeforeEach(func() {
			logCacheClient.PromQLReturns(queryResult(
				sample("0", "abc",
					point("1", 0.2),
				),
				sample("dyado", "def",
					point("2", 0.4),
				),
				sample("2", "ghi",
					point("4", 0.5),
				),
			), nil)
		})

		It("ignores the corrupt data point", func() {
			Expect(fetchErr).NotTo(HaveOccurred())
			Expect(currentUsage).To(Equal(map[int][]fetchers.InstanceData{
				0: {
					{
						Time:       time.Unix(1, 0),
						InstanceID: 0,
						Value:      0.2,
					},
				},
				2: {
					{
						Time:       time.Unix(4, 0),
						InstanceID: 2,
						Value:      0.5,
					},
				},
			}))
		})
	})
})
