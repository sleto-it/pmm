// Code generated by mockery v1.0.0. DO NOT EDIT.

package inventory

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	reform "gopkg.in/reform.v1"

	models "github.com/percona/pmm/managed/models"
)

// mockConnectionChecker is an autogenerated mock type for the connectionChecker type
type mockConnectionChecker struct {
	mock.Mock
}

// CheckConnectionToService provides a mock function with given fields: ctx, q, service, agent
func (_m *mockConnectionChecker) CheckConnectionToService(ctx context.Context, q *reform.Querier, service *models.Service, agent *models.Agent) error {
	ret := _m.Called(ctx, q, service, agent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reform.Querier, *models.Service, *models.Agent) error); ok {
		r0 = rf(ctx, q, service, agent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
