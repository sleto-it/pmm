// Code generated by mockery. DO NOT EDIT.

package client

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
	mock "github.com/stretchr/testify/mock"

	agentv1 "github.com/percona/pmm/api/agent/v1"
	agentlocalv1 "github.com/percona/pmm/api/agentlocal/v1"
)

// mockSupervisor is an autogenerated mock type for the supervisor type
type mockSupervisor struct {
	mock.Mock
}

// AgentLogByID provides a mock function with given fields: _a0
func (_m *mockSupervisor) AgentLogByID(_a0 string) ([]string, uint) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AgentLogByID")
	}

	var r0 []string
	var r1 uint
	if rf, ok := ret.Get(0).(func(string) ([]string, uint)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) uint); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(uint)
	}

	return r0, r1
}

// AgentsList provides a mock function with given fields:
func (_m *mockSupervisor) AgentsList() []*agentlocalv1.AgentInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AgentsList")
	}

	var r0 []*agentlocalv1.AgentInfo
	if rf, ok := ret.Get(0).(func() []*agentlocalv1.AgentInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*agentlocalv1.AgentInfo)
		}
	}

	return r0
}

// Changes provides a mock function with given fields:
func (_m *mockSupervisor) Changes() <-chan *agentv1.StateChangedRequest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Changes")
	}

	var r0 <-chan *agentv1.StateChangedRequest
	if rf, ok := ret.Get(0).(func() <-chan *agentv1.StateChangedRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *agentv1.StateChangedRequest)
		}
	}

	return r0
}

// ClearChangesChannel provides a mock function with given fields:
func (_m *mockSupervisor) ClearChangesChannel() {
	_m.Called()
}

// Collect provides a mock function with given fields: _a0
func (_m *mockSupervisor) Collect(_a0 chan<- prometheus.Metric) {
	_m.Called(_a0)
}

// Describe provides a mock function with given fields: _a0
func (_m *mockSupervisor) Describe(_a0 chan<- *prometheus.Desc) {
	_m.Called(_a0)
}

// QANRequests provides a mock function with given fields:
func (_m *mockSupervisor) QANRequests() <-chan *agentv1.QANCollectRequest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for QANRequests")
	}

	var r0 <-chan *agentv1.QANCollectRequest
	if rf, ok := ret.Get(0).(func() <-chan *agentv1.QANCollectRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *agentv1.QANCollectRequest)
		}
	}

	return r0
}

// RestartAgents provides a mock function with given fields:
func (_m *mockSupervisor) RestartAgents() {
	_m.Called()
}

// SetState provides a mock function with given fields: _a0
func (_m *mockSupervisor) SetState(_a0 *agentv1.SetStateRequest) {
	_m.Called(_a0)
}

// newMockSupervisor creates a new instance of mockSupervisor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSupervisor(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockSupervisor {
	mock := &mockSupervisor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
