// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"sync"

	"github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/command"
)

type FakeUserAddActor struct {
	AddUserStub        func() (actor.User, error)
	addUserMutex       sync.RWMutex
	addUserArgsForCall []struct {
	}
	addUserReturns struct {
		result1 actor.User
		result2 error
	}
	addUserReturnsOnCall map[int]struct {
		result1 actor.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserAddActor) AddUser() (actor.User, error) {
	fake.addUserMutex.Lock()
	ret, specificReturn := fake.addUserReturnsOnCall[len(fake.addUserArgsForCall)]
	fake.addUserArgsForCall = append(fake.addUserArgsForCall, struct {
	}{})
	fake.recordInvocation("AddUser", []interface{}{})
	fake.addUserMutex.Unlock()
	if fake.AddUserStub != nil {
		return fake.AddUserStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.addUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserAddActor) AddUserCallCount() int {
	fake.addUserMutex.RLock()
	defer fake.addUserMutex.RUnlock()
	return len(fake.addUserArgsForCall)
}

func (fake *FakeUserAddActor) AddUserCalls(stub func() (actor.User, error)) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = stub
}

func (fake *FakeUserAddActor) AddUserReturns(result1 actor.User, result2 error) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = nil
	fake.addUserReturns = struct {
		result1 actor.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserAddActor) AddUserReturnsOnCall(i int, result1 actor.User, result2 error) {
	fake.addUserMutex.Lock()
	defer fake.addUserMutex.Unlock()
	fake.AddUserStub = nil
	if fake.addUserReturnsOnCall == nil {
		fake.addUserReturnsOnCall = make(map[int]struct {
			result1 actor.User
			result2 error
		})
	}
	fake.addUserReturnsOnCall[i] = struct {
		result1 actor.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserAddActor) Invocations() map[string][][]interface{} {
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

func (fake *FakeUserAddActor) recordInvocation(key string, args []interface{}) {
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

var _ command.UserAddActor = new(FakeUserAddActor)
