// Copyright 2022 Giuseppe De Palma, Matteo Trentin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DockerShell is an autogenerated mock type for the DockerShell type
type DockerShell struct {
	mock.Mock
}

// ComposeDown provides a mock function with given fields: ctx, composeFilePath
func (_m *DockerShell) ComposeDown(ctx context.Context, composeFilePath string) error {
	ret := _m.Called(ctx, composeFilePath)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, composeFilePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ComposeList provides a mock function with given fields: ctx
func (_m *DockerShell) ComposeList(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ComposeUp provides a mock function with given fields: ctx, composeFilePath
func (_m *DockerShell) ComposeUp(ctx context.Context, composeFilePath string) error {
	ret := _m.Called(ctx, composeFilePath)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, composeFilePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogTokens provides a mock function with given fields: ctx
func (_m *DockerShell) LogTokens(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDockerShell interface {
	mock.TestingT
	Cleanup(func())
}

// NewDockerShell creates a new instance of DockerShell. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDockerShell(t mockConstructorTestingTNewDockerShell) *DockerShell {
	mock := &DockerShell{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
