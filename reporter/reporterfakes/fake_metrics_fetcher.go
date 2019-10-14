// Code generated by counterfeiter. DO NOT EDIT.
package reporterfakes

import (
	"sync"

	"code.cloudfoundry.org/cpu-entitlement-plugin/cf"
	"code.cloudfoundry.org/cpu-entitlement-plugin/fetchers"
	"code.cloudfoundry.org/cpu-entitlement-plugin/reporter"
)

type FakeMetricsFetcher struct {
	FetchInstanceDataStub        func(string, map[int]cf.Instance) (map[int]fetchers.InstanceData, error)
	fetchInstanceDataMutex       sync.RWMutex
	fetchInstanceDataArgsForCall []struct {
		arg1 string
		arg2 map[int]cf.Instance
	}
	fetchInstanceDataReturns struct {
		result1 map[int]fetchers.InstanceData
		result2 error
	}
	fetchInstanceDataReturnsOnCall map[int]struct {
		result1 map[int]fetchers.InstanceData
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetricsFetcher) FetchInstanceData(arg1 string, arg2 map[int]cf.Instance) (map[int]fetchers.InstanceData, error) {
	fake.fetchInstanceDataMutex.Lock()
	ret, specificReturn := fake.fetchInstanceDataReturnsOnCall[len(fake.fetchInstanceDataArgsForCall)]
	fake.fetchInstanceDataArgsForCall = append(fake.fetchInstanceDataArgsForCall, struct {
		arg1 string
		arg2 map[int]cf.Instance
	}{arg1, arg2})
	fake.recordInvocation("FetchInstanceData", []interface{}{arg1, arg2})
	fake.fetchInstanceDataMutex.Unlock()
	if fake.FetchInstanceDataStub != nil {
		return fake.FetchInstanceDataStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchInstanceDataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMetricsFetcher) FetchInstanceDataCallCount() int {
	fake.fetchInstanceDataMutex.RLock()
	defer fake.fetchInstanceDataMutex.RUnlock()
	return len(fake.fetchInstanceDataArgsForCall)
}

func (fake *FakeMetricsFetcher) FetchInstanceDataCalls(stub func(string, map[int]cf.Instance) (map[int]fetchers.InstanceData, error)) {
	fake.fetchInstanceDataMutex.Lock()
	defer fake.fetchInstanceDataMutex.Unlock()
	fake.FetchInstanceDataStub = stub
}

func (fake *FakeMetricsFetcher) FetchInstanceDataArgsForCall(i int) (string, map[int]cf.Instance) {
	fake.fetchInstanceDataMutex.RLock()
	defer fake.fetchInstanceDataMutex.RUnlock()
	argsForCall := fake.fetchInstanceDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMetricsFetcher) FetchInstanceDataReturns(result1 map[int]fetchers.InstanceData, result2 error) {
	fake.fetchInstanceDataMutex.Lock()
	defer fake.fetchInstanceDataMutex.Unlock()
	fake.FetchInstanceDataStub = nil
	fake.fetchInstanceDataReturns = struct {
		result1 map[int]fetchers.InstanceData
		result2 error
	}{result1, result2}
}

func (fake *FakeMetricsFetcher) FetchInstanceDataReturnsOnCall(i int, result1 map[int]fetchers.InstanceData, result2 error) {
	fake.fetchInstanceDataMutex.Lock()
	defer fake.fetchInstanceDataMutex.Unlock()
	fake.FetchInstanceDataStub = nil
	if fake.fetchInstanceDataReturnsOnCall == nil {
		fake.fetchInstanceDataReturnsOnCall = make(map[int]struct {
			result1 map[int]fetchers.InstanceData
			result2 error
		})
	}
	fake.fetchInstanceDataReturnsOnCall[i] = struct {
		result1 map[int]fetchers.InstanceData
		result2 error
	}{result1, result2}
}

func (fake *FakeMetricsFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchInstanceDataMutex.RLock()
	defer fake.fetchInstanceDataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMetricsFetcher) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ reporter.MetricsFetcher = new(FakeMetricsFetcher)
