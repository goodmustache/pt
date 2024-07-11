// Code generated by counterfeiter. DO NOT EDIT.
package actorfakes

import (
	"sync"

	"github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/config"
)

type FakeConfig struct {
	AddUserStub        func(config.User) error
	addUserMutex       sync.RWMutex
	addUserArgsForCall []struct {
		arg1 config.User
	}
	addUserReturns struct {
		result1 error
	}
	addUserReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) AddUser(arg1 config.User) error {
	fake.addUserMutex.Lock()
	ret, specificReturn := fake.addUserReturnsOnCall[len(fake.addUserArgsForCall)]
	fake.addUserArgsForCall = append(fake.addUserArgsForCall, struct {
		arg1 config.User
	}{arg1})
	stub := fake.AddUserStub
	fakeReturns := fake.addUserReturns
	fake.recordInvocation("AddUser", []interface{}{arg1})
	fake.addUserMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) AddUserCallCount() int {
	fake.addUserMutex.RLock()
	defer fake.addUserMutex.RUnlock()
	return len(fake.addUserArgsForCall)
}

func (fake *FakeConfig) AddUserCalls(stub func(config.User) error) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = stub
}

func (fake *FakeConfig) AddUserArgsForCall(i int) config.User {
	fake.addUserMutex.RLock()
	defer fake.addUserMutex.RUnlock()
	argsForCall := fake.addUserArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfig) AddUserReturns(result1 error) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = nil
	fake.addUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfig) AddUserReturnsOnCall(i int, result1 error) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = nil
	if fake.addUserReturnsOnCall == nil {
		fake.addUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addUserMutex.RLock()
	defer fake.addUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
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

var _ actor.Config = new(FakeConfig)
