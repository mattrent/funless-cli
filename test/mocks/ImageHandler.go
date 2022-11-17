// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/docker/docker/client"

	mock "github.com/stretchr/testify/mock"
)

// ImageHandler is an autogenerated mock type for the ImageHandler type
type ImageHandler struct {
	mock.Mock
}

// Exists provides a mock function with given fields: ctx, dockerClient, image
func (_m *ImageHandler) Exists(ctx context.Context, dockerClient *client.Client, image string) (bool, error) {
	ret := _m.Called(ctx, dockerClient, image)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *client.Client, string) bool); ok {
		r0 = rf(ctx, dockerClient, image)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.Client, string) error); ok {
		r1 = rf(ctx, dockerClient, image)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Pull provides a mock function with given fields: ctx, dockerClient, image
func (_m *ImageHandler) Pull(ctx context.Context, dockerClient *client.Client, image string) error {
	ret := _m.Called(ctx, dockerClient, image)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *client.Client, string) error); ok {
		r0 = rf(ctx, dockerClient, image)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewImageHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewImageHandler creates a new instance of ImageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImageHandler(t mockConstructorTestingTNewImageHandler) *ImageHandler {
	mock := &ImageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
