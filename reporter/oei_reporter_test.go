package reporter_test

import (
	"errors"

	"code.cloudfoundry.org/cpu-entitlement-plugin/cf"
	"code.cloudfoundry.org/cpu-entitlement-plugin/fetchers"
	"code.cloudfoundry.org/cpu-entitlement-plugin/reporter"
	"code.cloudfoundry.org/cpu-entitlement-plugin/reporter/reporterfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Over-entitlement Instances Reporter", func() {
	var (
		oeiReporter        reporter.OverEntitlementInstances
		fakeCfClient       *reporterfakes.FakeCloudFoundryClient
		fakeMetricsFetcher *reporterfakes.FakeMetricsFetcher
		report             reporter.OEIReport
		err                error
	)

	BeforeEach(func() {
		fakeCfClient = new(reporterfakes.FakeCloudFoundryClient)
		fakeMetricsFetcher = new(reporterfakes.FakeMetricsFetcher)

		fakeCfClient.GetCurrentOrgReturns("org", nil)
		fakeCfClient.UsernameReturns("user", nil)
		fakeCfClient.GetSpacesReturns([]cf.Space{
			{
				Name: "space1",
				Applications: []cf.Application{
					{Name: "app1", Guid: "space1-app1-guid"},
					{Name: "app2", Guid: "space1-app2-guid"},
				},
			},
			{
				Name: "space2",
				Applications: []cf.Application{
					{Name: "app1", Guid: "space2-app1-guid"},
				},
			},
		}, nil)

		fakeMetricsFetcher.FetchInstanceDataStub = func(appGuid string, appInstances map[int]cf.Instance) (map[int]fetchers.InstanceData, error) {
			switch appGuid {
			case "space1-app1-guid":
				return map[int]fetchers.InstanceData{
					0: {Value: 1.5},
					1: {Value: 0.5},
				}, nil
			case "space1-app2-guid":
				return map[int]fetchers.InstanceData{
					0: {Value: 0.3},
				}, nil
			case "space2-app1-guid":
				return map[int]fetchers.InstanceData{
					0: {Value: 0.2},
				}, nil
			}

			return nil, nil
		}

		oeiReporter = reporter.NewOverEntitlementInstances(fakeCfClient, fakeMetricsFetcher)
	})

	JustBeforeEach(func() {
		report, err = oeiReporter.OverEntitlementInstances()
	})

	It("succeeds", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns all instances that are over entitlement", func() {
		Expect(report).To(Equal(reporter.OEIReport{
			Org:      "org",
			Username: "user",
			SpaceReports: []reporter.SpaceReport{
				reporter.SpaceReport{
					SpaceName: "space1",
					Apps: []string{
						"app1",
					},
				},
			},
		}))
	})

	When("fetching the list of apps fails", func() {
		BeforeEach(func() {
			fakeCfClient.GetSpacesReturns(nil, errors.New("get-space-error"))
		})

		It("returns the error", func() {
			Expect(err).To(MatchError("get-space-error"))
		})
	})

	When("getting the entitlement usage for an app fails", func() {
		BeforeEach(func() {
			fakeMetricsFetcher.FetchInstanceDataReturns(nil, errors.New("fetch-error"))
		})

		It("returns the error", func() {
			Expect(err).To(MatchError("fetch-error"))
		})
	})

	When("getting the current org fails", func() {
		BeforeEach(func() {
			fakeCfClient.GetCurrentOrgReturns("", errors.New("get-org-error"))
		})

		It("returns the error", func() {
			Expect(err).To(MatchError("get-org-error"))
		})
	})

	When("getting the username fails", func() {
		BeforeEach(func() {
			fakeCfClient.UsernameReturns("", errors.New("get-user-error"))
		})

		It("returns the error", func() {
			Expect(err).To(MatchError("get-user-error"))
		})
	})

	When("spaces are not sorted alphabetically", func() {
		BeforeEach(func() {
			fakeCfClient.GetSpacesReturns([]cf.Space{
				{
					Name: "space2",
					Applications: []cf.Application{
						{Name: "app1", Guid: "space2-app1-guid"},
					},
				},
				{
					Name: "space1",
					Applications: []cf.Application{
						{Name: "app1", Guid: "space1-app1-guid"},
					},
				},
			}, nil)
			fakeMetricsFetcher.FetchInstanceDataReturns(
				map[int]fetchers.InstanceData{
					0: {Value: 1.5},
				}, nil)
		})

		It("reports sorted spaces", func() {
			Expect(len(report.SpaceReports)).To(Equal(2))
			Expect(report.SpaceReports[0].SpaceName).To(Equal("space1"))
			Expect(report.SpaceReports[1].SpaceName).To(Equal("space2"))
		})
	})

	When("apps in a single space are not sorted alphabetically", func() {
		BeforeEach(func() {
			fakeCfClient.GetSpacesReturns([]cf.Space{
				{
					Name: "space1",
					Applications: []cf.Application{
						{Name: "app2", Guid: "space1-app2-guid"},
						{Name: "app1", Guid: "space1-app1-guid"},
					},
				},
			}, nil)
			fakeMetricsFetcher.FetchInstanceDataReturns(
				map[int]fetchers.InstanceData{
					0: {Value: 1.5},
				}, nil)
		})

		It("reports sorted apps in the report", func() {
			Expect(len(report.SpaceReports)).To(Equal(1))
			Expect(len(report.SpaceReports[0].Apps)).To(Equal(2))
			Expect(report.SpaceReports[0].Apps[0]).To(Equal("app1"))
			Expect(report.SpaceReports[0].Apps[1]).To(Equal("app2"))
		})
	})
})
