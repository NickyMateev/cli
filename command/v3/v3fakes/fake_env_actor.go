// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeEnvActor struct {
	GetEnvironmentVariablesByApplicationNameAndSpaceStub        func(appName string, spaceGUID string) (v3action.EnvironmentVariableGroups, v3action.Warnings, error)
	getEnvironmentVariablesByApplicationNameAndSpaceMutex       sync.RWMutex
	getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getEnvironmentVariablesByApplicationNameAndSpaceReturns struct {
		result1 v3action.EnvironmentVariableGroups
		result2 v3action.Warnings
		result3 error
	}
	getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.EnvironmentVariableGroups
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpace(appName string, spaceGUID string) (v3action.EnvironmentVariableGroups, v3action.Warnings, error) {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall[len(fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall)]
	fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall = append(fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetEnvironmentVariablesByApplicationNameAndSpace", []interface{}{appName, spaceGUID})
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.Unlock()
	if fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub != nil {
		return fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getEnvironmentVariablesByApplicationNameAndSpaceReturns.result1, fake.getEnvironmentVariablesByApplicationNameAndSpaceReturns.result2, fake.getEnvironmentVariablesByApplicationNameAndSpaceReturns.result3
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceCallCount() int {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RLock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RUnlock()
	return len(fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall)
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RLock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RUnlock()
	return fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall[i].appName, fake.getEnvironmentVariablesByApplicationNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceReturns(result1 v3action.EnvironmentVariableGroups, result2 v3action.Warnings, result3 error) {
	fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub = nil
	fake.getEnvironmentVariablesByApplicationNameAndSpaceReturns = struct {
		result1 v3action.EnvironmentVariableGroups
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeEnvActor) GetEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall(i int, result1 v3action.EnvironmentVariableGroups, result2 v3action.Warnings, result3 error) {
	fake.GetEnvironmentVariablesByApplicationNameAndSpaceStub = nil
	if fake.getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall == nil {
		fake.getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.EnvironmentVariableGroups
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getEnvironmentVariablesByApplicationNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.EnvironmentVariableGroups
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeEnvActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RLock()
	defer fake.getEnvironmentVariablesByApplicationNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnvActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.EnvActor = new(FakeEnvActor)