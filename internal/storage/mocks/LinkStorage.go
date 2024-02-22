// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	storage "github.com/KKitsun/link-shortener-svc/internal/storage"
	mock "github.com/stretchr/testify/mock"
)

// LinkStorage is an autogenerated mock type for the LinkStorage type
type LinkStorage struct {
	mock.Mock
}

// Link provides a mock function with given fields:
func (_m *LinkStorage) Link() storage.LinkQ {
	ret := _m.Called()

	var r0 storage.LinkQ
	if rf, ok := ret.Get(0).(func() storage.LinkQ); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.LinkQ)
		}
	}

	return r0
}

// New provides a mock function with given fields:
func (_m *LinkStorage) New() storage.LinkStorage {
	ret := _m.Called()

	var r0 storage.LinkStorage
	if rf, ok := ret.Get(0).(func() storage.LinkStorage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.LinkStorage)
		}
	}

	return r0
}

type mockConstructorTestingTNewLinkStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewLinkStorage creates a new instance of LinkStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLinkStorage(t mockConstructorTestingTNewLinkStorage) *LinkStorage {
	mock := &LinkStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
