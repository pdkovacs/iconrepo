// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "igo-repo/internal/app/domain"
	authr "igo-repo/internal/app/security/authr"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddIconfile provides a mock function with given fields: iconName, iconfile, modifiedBy
func (_m *Repository) AddIconfile(iconName string, iconfile domain.Iconfile, modifiedBy authr.UserInfo) error {
	ret := _m.Called(iconName, iconfile, modifiedBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.Iconfile, authr.UserInfo) error); ok {
		r0 = rf(iconName, iconfile, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddTag provides a mock function with given fields: iconName, tag, modifiedBy
func (_m *Repository) AddTag(iconName string, tag string, modifiedBy authr.UserInfo) error {
	ret := _m.Called(iconName, tag, modifiedBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, authr.UserInfo) error); ok {
		r0 = rf(iconName, tag, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateIcon provides a mock function with given fields: iconName, iconfile, modifiedBy
func (_m *Repository) CreateIcon(iconName string, iconfile domain.Iconfile, modifiedBy authr.UserInfo) error {
	ret := _m.Called(iconName, iconfile, modifiedBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.Iconfile, authr.UserInfo) error); ok {
		r0 = rf(iconName, iconfile, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIcon provides a mock function with given fields: iconName, modifiedBy
func (_m *Repository) DeleteIcon(iconName string, modifiedBy authr.UserInfo) error {
	ret := _m.Called(iconName, modifiedBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, authr.UserInfo) error); ok {
		r0 = rf(iconName, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteIconfile provides a mock function with given fields: iconName, iconfile, modifiedBy
func (_m *Repository) DeleteIconfile(iconName string, iconfile domain.IconfileDescriptor, modifiedBy authr.UserInfo) error {
	ret := _m.Called(iconName, iconfile, modifiedBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.IconfileDescriptor, authr.UserInfo) error); ok {
		r0 = rf(iconName, iconfile, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeAllIcons provides a mock function with given fields:
func (_m *Repository) DescribeAllIcons() ([]domain.IconDescriptor, error) {
	ret := _m.Called()

	var r0 []domain.IconDescriptor
	if rf, ok := ret.Get(0).(func() []domain.IconDescriptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.IconDescriptor)
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

// DescribeIcon provides a mock function with given fields: iconName
func (_m *Repository) DescribeIcon(iconName string) (domain.IconDescriptor, error) {
	ret := _m.Called(iconName)

	var r0 domain.IconDescriptor
	if rf, ok := ret.Get(0).(func(string) domain.IconDescriptor); ok {
		r0 = rf(iconName)
	} else {
		r0 = ret.Get(0).(domain.IconDescriptor)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(iconName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIconFile provides a mock function with given fields: iconName, iconfile
func (_m *Repository) GetIconFile(iconName string, iconfile domain.IconfileDescriptor) ([]byte, error) {
	ret := _m.Called(iconName, iconfile)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, domain.IconfileDescriptor) []byte); ok {
		r0 = rf(iconName, iconfile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, domain.IconfileDescriptor) error); ok {
		r1 = rf(iconName, iconfile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields:
func (_m *Repository) GetTags() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// RemoveTag provides a mock function with given fields: iconName, tag, modifiedBy
func (_m *Repository) RemoveTag(iconName string, tag string, modifiedBy authr.UserInfo) error {
	ret := _m.Called(iconName, tag, modifiedBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, authr.UserInfo) error); ok {
		r0 = rf(iconName, tag, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}