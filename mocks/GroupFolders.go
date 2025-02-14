// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/avpalienko/gonextcloud"
)

// GroupFolders is an autogenerated mock type for the GroupFolders type
type GroupFolders struct {
	mock.Mock
}

// AddGroup provides a mock function with given fields: folderID, groupName
func (_m *GroupFolders) AddGroup(folderID int, groupName string) error {
	ret := _m.Called(folderID, groupName)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(folderID, groupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: name
func (_m *GroupFolders) Create(name string) (int, error) {
	ret := _m.Called(name)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *GroupFolders) Get(id int) (gonextcloud.GroupFolder, error) {
	ret := _m.Called(id)

	var r0 gonextcloud.GroupFolder
	if rf, ok := ret.Get(0).(func(int) gonextcloud.GroupFolder); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(gonextcloud.GroupFolder)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *GroupFolders) List() (map[int]gonextcloud.GroupFolder, error) {
	ret := _m.Called()

	var r0 map[int]gonextcloud.GroupFolder
	if rf, ok := ret.Get(0).(func() map[int]gonextcloud.GroupFolder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]gonextcloud.GroupFolder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveGroup provides a mock function with given fields: folderID, groupName
func (_m *GroupFolders) RemoveGroup(folderID int, groupName string) error {
	ret := _m.Called(folderID, groupName)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(folderID, groupName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rename provides a mock function with given fields: groupID, name
func (_m *GroupFolders) Rename(groupID int, name string) error {
	ret := _m.Called(groupID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(groupID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGroupPermissions provides a mock function with given fields: folderID, groupName, permission
func (_m *GroupFolders) SetGroupPermissions(folderID int, groupName string, permission gonextcloud.SharePermission) error {
	ret := _m.Called(folderID, groupName, permission)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, gonextcloud.SharePermission) error); ok {
		r0 = rf(folderID, groupName, permission)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetQuota provides a mock function with given fields: folderID, quota
func (_m *GroupFolders) SetQuota(folderID int, quota int) error {
	ret := _m.Called(folderID, quota)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(folderID, quota)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
