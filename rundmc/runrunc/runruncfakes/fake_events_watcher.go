// Code generated by counterfeiter. DO NOT EDIT.
package runruncfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/runrunc"
	"code.cloudfoundry.org/lager"
)

type FakeEventsWatcher struct {
	WatchEventsStub        func(log lager.Logger, handle string) error
	watchEventsMutex       sync.RWMutex
	watchEventsArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	watchEventsReturns struct {
		result1 error
	}
	watchEventsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventsWatcher) WatchEvents(log lager.Logger, handle string) error {
	fake.watchEventsMutex.Lock()
	ret, specificReturn := fake.watchEventsReturnsOnCall[len(fake.watchEventsArgsForCall)]
	fake.watchEventsArgsForCall = append(fake.watchEventsArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("WatchEvents", []interface{}{log, handle})
	fake.watchEventsMutex.Unlock()
	if fake.WatchEventsStub != nil {
		return fake.WatchEventsStub(log, handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.watchEventsReturns.result1
}

func (fake *FakeEventsWatcher) WatchEventsCallCount() int {
	fake.watchEventsMutex.RLock()
	defer fake.watchEventsMutex.RUnlock()
	return len(fake.watchEventsArgsForCall)
}

func (fake *FakeEventsWatcher) WatchEventsArgsForCall(i int) (lager.Logger, string) {
	fake.watchEventsMutex.RLock()
	defer fake.watchEventsMutex.RUnlock()
	return fake.watchEventsArgsForCall[i].log, fake.watchEventsArgsForCall[i].handle
}

func (fake *FakeEventsWatcher) WatchEventsReturns(result1 error) {
	fake.WatchEventsStub = nil
	fake.watchEventsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventsWatcher) WatchEventsReturnsOnCall(i int, result1 error) {
	fake.WatchEventsStub = nil
	if fake.watchEventsReturnsOnCall == nil {
		fake.watchEventsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.watchEventsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventsWatcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.watchEventsMutex.RLock()
	defer fake.watchEventsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventsWatcher) recordInvocation(key string, args []interface{}) {
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

var _ runrunc.EventsWatcher = new(FakeEventsWatcher)