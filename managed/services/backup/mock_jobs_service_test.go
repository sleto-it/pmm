// Code generated by mockery v2.39.1. DO NOT EDIT.

package backup

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
)

// mockJobsService is an autogenerated mock type for the jobsService type
type mockJobsService struct {
	mock.Mock
}

// StartMongoDBBackupJob provides a mock function with given fields: service, jobID, pmmAgentID, timeout, name, dbConfig, mode, dataModel, locationConfig, folder
func (_m *mockJobsService) StartMongoDBBackupJob(service *models.Service, jobID string, pmmAgentID string, timeout time.Duration, name string, dbConfig *models.DBConfig, mode models.BackupMode, dataModel models.DataModel, locationConfig *models.BackupLocationConfig, folder string) error {
	ret := _m.Called(service, jobID, pmmAgentID, timeout, name, dbConfig, mode, dataModel, locationConfig, folder)

	if len(ret) == 0 {
		panic("no return value specified for StartMongoDBBackupJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Service, string, string, time.Duration, string, *models.DBConfig, models.BackupMode, models.DataModel, *models.BackupLocationConfig, string) error); ok {
		r0 = rf(service, jobID, pmmAgentID, timeout, name, dbConfig, mode, dataModel, locationConfig, folder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMongoDBRestoreBackupJob provides a mock function with given fields: service, jobID, pmmAgentID, timeout, name, pbmBackupName, dbConfig, dataModel, locationConfig, pitrTimestamp, folder
func (_m *mockJobsService) StartMongoDBRestoreBackupJob(service *models.Service, jobID string, pmmAgentID string, timeout time.Duration, name string, pbmBackupName string, dbConfig *models.DBConfig, dataModel models.DataModel, locationConfig *models.BackupLocationConfig, pitrTimestamp time.Time, folder string) error {
	ret := _m.Called(service, jobID, pmmAgentID, timeout, name, pbmBackupName, dbConfig, dataModel, locationConfig, pitrTimestamp, folder)

	if len(ret) == 0 {
		panic("no return value specified for StartMongoDBRestoreBackupJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Service, string, string, time.Duration, string, string, *models.DBConfig, models.DataModel, *models.BackupLocationConfig, time.Time, string) error); ok {
		r0 = rf(service, jobID, pmmAgentID, timeout, name, pbmBackupName, dbConfig, dataModel, locationConfig, pitrTimestamp, folder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMySQLBackupJob provides a mock function with given fields: jobID, pmmAgentID, timeout, name, dbConfig, locationConfig, folder
func (_m *mockJobsService) StartMySQLBackupJob(jobID string, pmmAgentID string, timeout time.Duration, name string, dbConfig *models.DBConfig, locationConfig *models.BackupLocationConfig, folder string) error {
	ret := _m.Called(jobID, pmmAgentID, timeout, name, dbConfig, locationConfig, folder)

	if len(ret) == 0 {
		panic("no return value specified for StartMySQLBackupJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, time.Duration, string, *models.DBConfig, *models.BackupLocationConfig, string) error); ok {
		r0 = rf(jobID, pmmAgentID, timeout, name, dbConfig, locationConfig, folder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartMySQLRestoreBackupJob provides a mock function with given fields: jobID, pmmAgentID, serviceID, timeout, name, locationConfig, folder
func (_m *mockJobsService) StartMySQLRestoreBackupJob(jobID string, pmmAgentID string, serviceID string, timeout time.Duration, name string, locationConfig *models.BackupLocationConfig, folder string) error {
	ret := _m.Called(jobID, pmmAgentID, serviceID, timeout, name, locationConfig, folder)

	if len(ret) == 0 {
		panic("no return value specified for StartMySQLRestoreBackupJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, time.Duration, string, *models.BackupLocationConfig, string) error); ok {
		r0 = rf(jobID, pmmAgentID, serviceID, timeout, name, locationConfig, folder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopJob provides a mock function with given fields: jobID
func (_m *mockJobsService) StopJob(jobID string) error {
	ret := _m.Called(jobID)

	if len(ret) == 0 {
		panic("no return value specified for StopJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockJobsService creates a new instance of mockJobsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockJobsService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockJobsService {
	mock := &mockJobsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
