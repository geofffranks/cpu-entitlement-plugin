// Code generated by counterfeiter. DO NOT EDIT.
package fetchersfakes

import (
	"context"
	"sync"
	"time"

	"code.cloudfoundry.org/cpu-entitlement-plugin/fetchers"
	client "code.cloudfoundry.org/go-log-cache"
	"code.cloudfoundry.org/go-log-cache/rpc/logcache_v1"
	"code.cloudfoundry.org/go-loggregator/v9/rpc/loggregator_v2"
)

type FakeLogCacheClient struct {
	PromQLStub        func(context.Context, string, ...client.PromQLOption) (*logcache_v1.PromQL_InstantQueryResult, error)
	promQLMutex       sync.RWMutex
	promQLArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []client.PromQLOption
	}
	promQLReturns struct {
		result1 *logcache_v1.PromQL_InstantQueryResult
		result2 error
	}
	promQLReturnsOnCall map[int]struct {
		result1 *logcache_v1.PromQL_InstantQueryResult
		result2 error
	}
	PromQLRangeStub        func(context.Context, string, ...client.PromQLOption) (*logcache_v1.PromQL_RangeQueryResult, error)
	promQLRangeMutex       sync.RWMutex
	promQLRangeArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []client.PromQLOption
	}
	promQLRangeReturns struct {
		result1 *logcache_v1.PromQL_RangeQueryResult
		result2 error
	}
	promQLRangeReturnsOnCall map[int]struct {
		result1 *logcache_v1.PromQL_RangeQueryResult
		result2 error
	}
	ReadStub        func(context.Context, string, time.Time, ...client.ReadOption) ([]*loggregator_v2.Envelope, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 time.Time
		arg4 []client.ReadOption
	}
	readReturns struct {
		result1 []*loggregator_v2.Envelope
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 []*loggregator_v2.Envelope
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogCacheClient) PromQL(arg1 context.Context, arg2 string, arg3 ...client.PromQLOption) (*logcache_v1.PromQL_InstantQueryResult, error) {
	fake.promQLMutex.Lock()
	ret, specificReturn := fake.promQLReturnsOnCall[len(fake.promQLArgsForCall)]
	fake.promQLArgsForCall = append(fake.promQLArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []client.PromQLOption
	}{arg1, arg2, arg3})
	stub := fake.PromQLStub
	fakeReturns := fake.promQLReturns
	fake.recordInvocation("PromQL", []interface{}{arg1, arg2, arg3})
	fake.promQLMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLogCacheClient) PromQLCallCount() int {
	fake.promQLMutex.RLock()
	defer fake.promQLMutex.RUnlock()
	return len(fake.promQLArgsForCall)
}

func (fake *FakeLogCacheClient) PromQLCalls(stub func(context.Context, string, ...client.PromQLOption) (*logcache_v1.PromQL_InstantQueryResult, error)) {
	fake.promQLMutex.Lock()
	defer fake.promQLMutex.Unlock()
	fake.PromQLStub = stub
}

func (fake *FakeLogCacheClient) PromQLArgsForCall(i int) (context.Context, string, []client.PromQLOption) {
	fake.promQLMutex.RLock()
	defer fake.promQLMutex.RUnlock()
	argsForCall := fake.promQLArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLogCacheClient) PromQLReturns(result1 *logcache_v1.PromQL_InstantQueryResult, result2 error) {
	fake.promQLMutex.Lock()
	defer fake.promQLMutex.Unlock()
	fake.PromQLStub = nil
	fake.promQLReturns = struct {
		result1 *logcache_v1.PromQL_InstantQueryResult
		result2 error
	}{result1, result2}
}

func (fake *FakeLogCacheClient) PromQLReturnsOnCall(i int, result1 *logcache_v1.PromQL_InstantQueryResult, result2 error) {
	fake.promQLMutex.Lock()
	defer fake.promQLMutex.Unlock()
	fake.PromQLStub = nil
	if fake.promQLReturnsOnCall == nil {
		fake.promQLReturnsOnCall = make(map[int]struct {
			result1 *logcache_v1.PromQL_InstantQueryResult
			result2 error
		})
	}
	fake.promQLReturnsOnCall[i] = struct {
		result1 *logcache_v1.PromQL_InstantQueryResult
		result2 error
	}{result1, result2}
}

func (fake *FakeLogCacheClient) PromQLRange(arg1 context.Context, arg2 string, arg3 ...client.PromQLOption) (*logcache_v1.PromQL_RangeQueryResult, error) {
	fake.promQLRangeMutex.Lock()
	ret, specificReturn := fake.promQLRangeReturnsOnCall[len(fake.promQLRangeArgsForCall)]
	fake.promQLRangeArgsForCall = append(fake.promQLRangeArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []client.PromQLOption
	}{arg1, arg2, arg3})
	stub := fake.PromQLRangeStub
	fakeReturns := fake.promQLRangeReturns
	fake.recordInvocation("PromQLRange", []interface{}{arg1, arg2, arg3})
	fake.promQLRangeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLogCacheClient) PromQLRangeCallCount() int {
	fake.promQLRangeMutex.RLock()
	defer fake.promQLRangeMutex.RUnlock()
	return len(fake.promQLRangeArgsForCall)
}

func (fake *FakeLogCacheClient) PromQLRangeCalls(stub func(context.Context, string, ...client.PromQLOption) (*logcache_v1.PromQL_RangeQueryResult, error)) {
	fake.promQLRangeMutex.Lock()
	defer fake.promQLRangeMutex.Unlock()
	fake.PromQLRangeStub = stub
}

func (fake *FakeLogCacheClient) PromQLRangeArgsForCall(i int) (context.Context, string, []client.PromQLOption) {
	fake.promQLRangeMutex.RLock()
	defer fake.promQLRangeMutex.RUnlock()
	argsForCall := fake.promQLRangeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLogCacheClient) PromQLRangeReturns(result1 *logcache_v1.PromQL_RangeQueryResult, result2 error) {
	fake.promQLRangeMutex.Lock()
	defer fake.promQLRangeMutex.Unlock()
	fake.PromQLRangeStub = nil
	fake.promQLRangeReturns = struct {
		result1 *logcache_v1.PromQL_RangeQueryResult
		result2 error
	}{result1, result2}
}

func (fake *FakeLogCacheClient) PromQLRangeReturnsOnCall(i int, result1 *logcache_v1.PromQL_RangeQueryResult, result2 error) {
	fake.promQLRangeMutex.Lock()
	defer fake.promQLRangeMutex.Unlock()
	fake.PromQLRangeStub = nil
	if fake.promQLRangeReturnsOnCall == nil {
		fake.promQLRangeReturnsOnCall = make(map[int]struct {
			result1 *logcache_v1.PromQL_RangeQueryResult
			result2 error
		})
	}
	fake.promQLRangeReturnsOnCall[i] = struct {
		result1 *logcache_v1.PromQL_RangeQueryResult
		result2 error
	}{result1, result2}
}

func (fake *FakeLogCacheClient) Read(arg1 context.Context, arg2 string, arg3 time.Time, arg4 ...client.ReadOption) ([]*loggregator_v2.Envelope, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 time.Time
		arg4 []client.ReadOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.ReadStub
	fakeReturns := fake.readReturns
	fake.recordInvocation("Read", []interface{}{arg1, arg2, arg3, arg4})
	fake.readMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLogCacheClient) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeLogCacheClient) ReadCalls(stub func(context.Context, string, time.Time, ...client.ReadOption) ([]*loggregator_v2.Envelope, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeLogCacheClient) ReadArgsForCall(i int) (context.Context, string, time.Time, []client.ReadOption) {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeLogCacheClient) ReadReturns(result1 []*loggregator_v2.Envelope, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 []*loggregator_v2.Envelope
		result2 error
	}{result1, result2}
}

func (fake *FakeLogCacheClient) ReadReturnsOnCall(i int, result1 []*loggregator_v2.Envelope, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 []*loggregator_v2.Envelope
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 []*loggregator_v2.Envelope
		result2 error
	}{result1, result2}
}

func (fake *FakeLogCacheClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.promQLMutex.RLock()
	defer fake.promQLMutex.RUnlock()
	fake.promQLRangeMutex.RLock()
	defer fake.promQLRangeMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogCacheClient) recordInvocation(key string, args []interface{}) {
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

var _ fetchers.LogCacheClient = new(FakeLogCacheClient)
