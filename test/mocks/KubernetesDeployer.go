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

// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// KubernetesDeployer is an autogenerated mock type for the KubernetesDeployer type
type KubernetesDeployer struct {
	mock.Mock
}

// CreateNamespace provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) CreateNamespace(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePrometheusConfigMap provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) CreatePrometheusConfigMap(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRole provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) CreateRole(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRoleBinding provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) CreateRoleBinding(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSvcAccount provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) CreateSvcAccount(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployCore provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) DeployCore(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployCoreService provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) DeployCoreService(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployPostgres provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) DeployPostgres(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployPostgresService provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) DeployPostgresService(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployPrometheus provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) DeployPrometheus(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployPrometheusService provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) DeployPrometheusService(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployWorker provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) DeployWorker(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartInitPostgres provides a mock function with given fields: ctx
func (_m *KubernetesDeployer) StartInitPostgres(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithConfig provides a mock function with given fields: config
func (_m *KubernetesDeployer) WithConfig(config string) error {
	ret := _m.Called(config)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewKubernetesDeployer interface {
	mock.TestingT
	Cleanup(func())
}

// NewKubernetesDeployer creates a new instance of KubernetesDeployer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKubernetesDeployer(t mockConstructorTestingTNewKubernetesDeployer) *KubernetesDeployer {
	mock := &KubernetesDeployer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
