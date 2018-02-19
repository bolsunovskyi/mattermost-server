// Code generated by mockery v1.0.0

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/bolsunovskyi/mattermost-server/model"
import store "github.com/bolsunovskyi/mattermost-server/store"

// ComplianceStore is an autogenerated mock type for the ComplianceStore type
type ComplianceStore struct {
	mock.Mock
}

// ComplianceExport provides a mock function with given fields: compliance
func (_m *ComplianceStore) ComplianceExport(compliance *model.Compliance) store.StoreChannel {
	ret := _m.Called(compliance)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Compliance) store.StoreChannel); ok {
		r0 = rf(compliance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ComplianceStore) Get(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields: offset, limit
func (_m *ComplianceStore) GetAll(offset int, limit int) store.StoreChannel {
	ret := _m.Called(offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int) store.StoreChannel); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// MessageExport provides a mock function with given fields: after, limit
func (_m *ComplianceStore) MessageExport(after int64, limit int) store.StoreChannel {
	ret := _m.Called(after, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64, int) store.StoreChannel); ok {
		r0 = rf(after, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: compliance
func (_m *ComplianceStore) Save(compliance *model.Compliance) store.StoreChannel {
	ret := _m.Called(compliance)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Compliance) store.StoreChannel); ok {
		r0 = rf(compliance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: compliance
func (_m *ComplianceStore) Update(compliance *model.Compliance) store.StoreChannel {
	ret := _m.Called(compliance)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Compliance) store.StoreChannel); ok {
		r0 = rf(compliance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
