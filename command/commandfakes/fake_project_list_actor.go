// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"sync"

	"github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/command"
)

type FakeProjectListActor struct {
	ProjectsStub        func() ([]actor.Project, error)
	projectsMutex       sync.RWMutex
	projectsArgsForCall []struct {
	}
	projectsReturns struct {
		result1 []actor.Project
		result2 error
	}
	projectsReturnsOnCall map[int]struct {
		result1 []actor.Project
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProjectListActor) Projects() ([]actor.Project, error) {
	fake.projectsMutex.Lock()
	ret, specificReturn := fake.projectsReturnsOnCall[len(fake.projectsArgsForCall)]
	fake.projectsArgsForCall = append(fake.projectsArgsForCall, struct {
	}{})
	stub := fake.ProjectsStub
	fakeReturns := fake.projectsReturns
	fake.recordInvocation("Projects", []interface{}{})
	fake.projectsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProjectListActor) ProjectsCallCount() int {
	fake.projectsMutex.RLock()
	defer fake.projectsMutex.RUnlock()
	return len(fake.projectsArgsForCall)
}

func (fake *FakeProjectListActor) ProjectsCalls(stub func() ([]actor.Project, error)) {
	fake.projectsMutex.Lock()
	defer fake.projectsMutex.Unlock()
	fake.ProjectsStub = stub
}

func (fake *FakeProjectListActor) ProjectsReturns(result1 []actor.Project, result2 error) {
	fake.projectsMutex.Lock()
	defer fake.projectsMutex.Unlock()
	fake.ProjectsStub = nil
	fake.projectsReturns = struct {
		result1 []actor.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeProjectListActor) ProjectsReturnsOnCall(i int, result1 []actor.Project, result2 error) {
	fake.projectsMutex.Lock()
	defer fake.projectsMutex.Unlock()
	fake.ProjectsStub = nil
	if fake.projectsReturnsOnCall == nil {
		fake.projectsReturnsOnCall = make(map[int]struct {
			result1 []actor.Project
			result2 error
		})
	}
	fake.projectsReturnsOnCall[i] = struct {
		result1 []actor.Project
		result2 error
	}{result1, result2}
}

func (fake *FakeProjectListActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.projectsMutex.RLock()
	defer fake.projectsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProjectListActor) recordInvocation(key string, args []interface{}) {
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

var _ command.ProjectListActor = new(FakeProjectListActor)
