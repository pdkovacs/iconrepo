// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	authr "iconrepo/internal/app/security/authr"

	domain "iconrepo/internal/app/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// AddIconfile provides a mock function with given fields: ctx, iconName, iconfile, modifiedBy
func (_m *Repository) AddIconfile(ctx context.Context, iconName string, iconfile domain.Iconfile, modifiedBy authr.UserInfo) error {
	ret := _m.Called(ctx, iconName, iconfile, modifiedBy)

	if len(ret) == 0 {
		panic("no return value specified for AddIconfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Iconfile, authr.UserInfo) error); ok {
		r0 = rf(ctx, iconName, iconfile, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_AddIconfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddIconfile'
type Repository_AddIconfile_Call struct {
	*mock.Call
}

// AddIconfile is a helper method to define mock.On call
//   - ctx context.Context
//   - iconName string
//   - iconfile domain.Iconfile
//   - modifiedBy authr.UserInfo
func (_e *Repository_Expecter) AddIconfile(ctx interface{}, iconName interface{}, iconfile interface{}, modifiedBy interface{}) *Repository_AddIconfile_Call {
	return &Repository_AddIconfile_Call{Call: _e.mock.On("AddIconfile", ctx, iconName, iconfile, modifiedBy)}
}

func (_c *Repository_AddIconfile_Call) Run(run func(ctx context.Context, iconName string, iconfile domain.Iconfile, modifiedBy authr.UserInfo)) *Repository_AddIconfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.Iconfile), args[3].(authr.UserInfo))
	})
	return _c
}

func (_c *Repository_AddIconfile_Call) Return(_a0 error) *Repository_AddIconfile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_AddIconfile_Call) RunAndReturn(run func(context.Context, string, domain.Iconfile, authr.UserInfo) error) *Repository_AddIconfile_Call {
	_c.Call.Return(run)
	return _c
}

// AddTag provides a mock function with given fields: ctx, iconName, tag, modifiedBy
func (_m *Repository) AddTag(ctx context.Context, iconName string, tag string, modifiedBy authr.UserInfo) error {
	ret := _m.Called(ctx, iconName, tag, modifiedBy)

	if len(ret) == 0 {
		panic("no return value specified for AddTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, authr.UserInfo) error); ok {
		r0 = rf(ctx, iconName, tag, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_AddTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTag'
type Repository_AddTag_Call struct {
	*mock.Call
}

// AddTag is a helper method to define mock.On call
//   - ctx context.Context
//   - iconName string
//   - tag string
//   - modifiedBy authr.UserInfo
func (_e *Repository_Expecter) AddTag(ctx interface{}, iconName interface{}, tag interface{}, modifiedBy interface{}) *Repository_AddTag_Call {
	return &Repository_AddTag_Call{Call: _e.mock.On("AddTag", ctx, iconName, tag, modifiedBy)}
}

func (_c *Repository_AddTag_Call) Run(run func(ctx context.Context, iconName string, tag string, modifiedBy authr.UserInfo)) *Repository_AddTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(authr.UserInfo))
	})
	return _c
}

func (_c *Repository_AddTag_Call) Return(_a0 error) *Repository_AddTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_AddTag_Call) RunAndReturn(run func(context.Context, string, string, authr.UserInfo) error) *Repository_AddTag_Call {
	_c.Call.Return(run)
	return _c
}

// CreateIcon provides a mock function with given fields: ctx, iconName, iconfile, modifiedBy
func (_m *Repository) CreateIcon(ctx context.Context, iconName string, iconfile domain.Iconfile, modifiedBy authr.UserInfo) error {
	ret := _m.Called(ctx, iconName, iconfile, modifiedBy)

	if len(ret) == 0 {
		panic("no return value specified for CreateIcon")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Iconfile, authr.UserInfo) error); ok {
		r0 = rf(ctx, iconName, iconfile, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_CreateIcon_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateIcon'
type Repository_CreateIcon_Call struct {
	*mock.Call
}

// CreateIcon is a helper method to define mock.On call
//   - ctx context.Context
//   - iconName string
//   - iconfile domain.Iconfile
//   - modifiedBy authr.UserInfo
func (_e *Repository_Expecter) CreateIcon(ctx interface{}, iconName interface{}, iconfile interface{}, modifiedBy interface{}) *Repository_CreateIcon_Call {
	return &Repository_CreateIcon_Call{Call: _e.mock.On("CreateIcon", ctx, iconName, iconfile, modifiedBy)}
}

func (_c *Repository_CreateIcon_Call) Run(run func(ctx context.Context, iconName string, iconfile domain.Iconfile, modifiedBy authr.UserInfo)) *Repository_CreateIcon_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.Iconfile), args[3].(authr.UserInfo))
	})
	return _c
}

func (_c *Repository_CreateIcon_Call) Return(_a0 error) *Repository_CreateIcon_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_CreateIcon_Call) RunAndReturn(run func(context.Context, string, domain.Iconfile, authr.UserInfo) error) *Repository_CreateIcon_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteIcon provides a mock function with given fields: ctx, iconName, modifiedBy
func (_m *Repository) DeleteIcon(ctx context.Context, iconName string, modifiedBy authr.UserInfo) error {
	ret := _m.Called(ctx, iconName, modifiedBy)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIcon")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, authr.UserInfo) error); ok {
		r0 = rf(ctx, iconName, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_DeleteIcon_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteIcon'
type Repository_DeleteIcon_Call struct {
	*mock.Call
}

// DeleteIcon is a helper method to define mock.On call
//   - ctx context.Context
//   - iconName string
//   - modifiedBy authr.UserInfo
func (_e *Repository_Expecter) DeleteIcon(ctx interface{}, iconName interface{}, modifiedBy interface{}) *Repository_DeleteIcon_Call {
	return &Repository_DeleteIcon_Call{Call: _e.mock.On("DeleteIcon", ctx, iconName, modifiedBy)}
}

func (_c *Repository_DeleteIcon_Call) Run(run func(ctx context.Context, iconName string, modifiedBy authr.UserInfo)) *Repository_DeleteIcon_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(authr.UserInfo))
	})
	return _c
}

func (_c *Repository_DeleteIcon_Call) Return(_a0 error) *Repository_DeleteIcon_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_DeleteIcon_Call) RunAndReturn(run func(context.Context, string, authr.UserInfo) error) *Repository_DeleteIcon_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteIconfile provides a mock function with given fields: ctx, iconName, iconfile, modifiedBy
func (_m *Repository) DeleteIconfile(ctx context.Context, iconName string, iconfile domain.IconfileDescriptor, modifiedBy authr.UserInfo) error {
	ret := _m.Called(ctx, iconName, iconfile, modifiedBy)

	if len(ret) == 0 {
		panic("no return value specified for DeleteIconfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.IconfileDescriptor, authr.UserInfo) error); ok {
		r0 = rf(ctx, iconName, iconfile, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_DeleteIconfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteIconfile'
type Repository_DeleteIconfile_Call struct {
	*mock.Call
}

// DeleteIconfile is a helper method to define mock.On call
//   - ctx context.Context
//   - iconName string
//   - iconfile domain.IconfileDescriptor
//   - modifiedBy authr.UserInfo
func (_e *Repository_Expecter) DeleteIconfile(ctx interface{}, iconName interface{}, iconfile interface{}, modifiedBy interface{}) *Repository_DeleteIconfile_Call {
	return &Repository_DeleteIconfile_Call{Call: _e.mock.On("DeleteIconfile", ctx, iconName, iconfile, modifiedBy)}
}

func (_c *Repository_DeleteIconfile_Call) Run(run func(ctx context.Context, iconName string, iconfile domain.IconfileDescriptor, modifiedBy authr.UserInfo)) *Repository_DeleteIconfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.IconfileDescriptor), args[3].(authr.UserInfo))
	})
	return _c
}

func (_c *Repository_DeleteIconfile_Call) Return(_a0 error) *Repository_DeleteIconfile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_DeleteIconfile_Call) RunAndReturn(run func(context.Context, string, domain.IconfileDescriptor, authr.UserInfo) error) *Repository_DeleteIconfile_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeAllIcons provides a mock function with given fields: ctx
func (_m *Repository) DescribeAllIcons(ctx context.Context) ([]domain.IconDescriptor, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAllIcons")
	}

	var r0 []domain.IconDescriptor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.IconDescriptor, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.IconDescriptor); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.IconDescriptor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_DescribeAllIcons_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeAllIcons'
type Repository_DescribeAllIcons_Call struct {
	*mock.Call
}

// DescribeAllIcons is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Repository_Expecter) DescribeAllIcons(ctx interface{}) *Repository_DescribeAllIcons_Call {
	return &Repository_DescribeAllIcons_Call{Call: _e.mock.On("DescribeAllIcons", ctx)}
}

func (_c *Repository_DescribeAllIcons_Call) Run(run func(ctx context.Context)) *Repository_DescribeAllIcons_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Repository_DescribeAllIcons_Call) Return(_a0 []domain.IconDescriptor, _a1 error) *Repository_DescribeAllIcons_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_DescribeAllIcons_Call) RunAndReturn(run func(context.Context) ([]domain.IconDescriptor, error)) *Repository_DescribeAllIcons_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeIcon provides a mock function with given fields: ctx, iconName
func (_m *Repository) DescribeIcon(ctx context.Context, iconName string) (domain.IconDescriptor, error) {
	ret := _m.Called(ctx, iconName)

	if len(ret) == 0 {
		panic("no return value specified for DescribeIcon")
	}

	var r0 domain.IconDescriptor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.IconDescriptor, error)); ok {
		return rf(ctx, iconName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.IconDescriptor); ok {
		r0 = rf(ctx, iconName)
	} else {
		r0 = ret.Get(0).(domain.IconDescriptor)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, iconName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_DescribeIcon_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeIcon'
type Repository_DescribeIcon_Call struct {
	*mock.Call
}

// DescribeIcon is a helper method to define mock.On call
//   - ctx context.Context
//   - iconName string
func (_e *Repository_Expecter) DescribeIcon(ctx interface{}, iconName interface{}) *Repository_DescribeIcon_Call {
	return &Repository_DescribeIcon_Call{Call: _e.mock.On("DescribeIcon", ctx, iconName)}
}

func (_c *Repository_DescribeIcon_Call) Run(run func(ctx context.Context, iconName string)) *Repository_DescribeIcon_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Repository_DescribeIcon_Call) Return(_a0 domain.IconDescriptor, _a1 error) *Repository_DescribeIcon_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_DescribeIcon_Call) RunAndReturn(run func(context.Context, string) (domain.IconDescriptor, error)) *Repository_DescribeIcon_Call {
	_c.Call.Return(run)
	return _c
}

// GetIconfile provides a mock function with given fields: ctx, iconName, iconfile
func (_m *Repository) GetIconfile(ctx context.Context, iconName string, iconfile domain.IconfileDescriptor) ([]byte, error) {
	ret := _m.Called(ctx, iconName, iconfile)

	if len(ret) == 0 {
		panic("no return value specified for GetIconfile")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.IconfileDescriptor) ([]byte, error)); ok {
		return rf(ctx, iconName, iconfile)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.IconfileDescriptor) []byte); ok {
		r0 = rf(ctx, iconName, iconfile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.IconfileDescriptor) error); ok {
		r1 = rf(ctx, iconName, iconfile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetIconfile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIconfile'
type Repository_GetIconfile_Call struct {
	*mock.Call
}

// GetIconfile is a helper method to define mock.On call
//   - ctx context.Context
//   - iconName string
//   - iconfile domain.IconfileDescriptor
func (_e *Repository_Expecter) GetIconfile(ctx interface{}, iconName interface{}, iconfile interface{}) *Repository_GetIconfile_Call {
	return &Repository_GetIconfile_Call{Call: _e.mock.On("GetIconfile", ctx, iconName, iconfile)}
}

func (_c *Repository_GetIconfile_Call) Run(run func(ctx context.Context, iconName string, iconfile domain.IconfileDescriptor)) *Repository_GetIconfile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(domain.IconfileDescriptor))
	})
	return _c
}

func (_c *Repository_GetIconfile_Call) Return(_a0 []byte, _a1 error) *Repository_GetIconfile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetIconfile_Call) RunAndReturn(run func(context.Context, string, domain.IconfileDescriptor) ([]byte, error)) *Repository_GetIconfile_Call {
	_c.Call.Return(run)
	return _c
}

// GetTags provides a mock function with given fields: ctx
func (_m *Repository) GetTags(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetTags")
	}

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

// Repository_GetTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTags'
type Repository_GetTags_Call struct {
	*mock.Call
}

// GetTags is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Repository_Expecter) GetTags(ctx interface{}) *Repository_GetTags_Call {
	return &Repository_GetTags_Call{Call: _e.mock.On("GetTags", ctx)}
}

func (_c *Repository_GetTags_Call) Run(run func(ctx context.Context)) *Repository_GetTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Repository_GetTags_Call) Return(_a0 []string, _a1 error) *Repository_GetTags_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetTags_Call) RunAndReturn(run func(context.Context) ([]string, error)) *Repository_GetTags_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveTag provides a mock function with given fields: ctx, iconName, tag, modifiedBy
func (_m *Repository) RemoveTag(ctx context.Context, iconName string, tag string, modifiedBy authr.UserInfo) error {
	ret := _m.Called(ctx, iconName, tag, modifiedBy)

	if len(ret) == 0 {
		panic("no return value specified for RemoveTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, authr.UserInfo) error); ok {
		r0 = rf(ctx, iconName, tag, modifiedBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_RemoveTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveTag'
type Repository_RemoveTag_Call struct {
	*mock.Call
}

// RemoveTag is a helper method to define mock.On call
//   - ctx context.Context
//   - iconName string
//   - tag string
//   - modifiedBy authr.UserInfo
func (_e *Repository_Expecter) RemoveTag(ctx interface{}, iconName interface{}, tag interface{}, modifiedBy interface{}) *Repository_RemoveTag_Call {
	return &Repository_RemoveTag_Call{Call: _e.mock.On("RemoveTag", ctx, iconName, tag, modifiedBy)}
}

func (_c *Repository_RemoveTag_Call) Run(run func(ctx context.Context, iconName string, tag string, modifiedBy authr.UserInfo)) *Repository_RemoveTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(authr.UserInfo))
	})
	return _c
}

func (_c *Repository_RemoveTag_Call) Return(_a0 error) *Repository_RemoveTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_RemoveTag_Call) RunAndReturn(run func(context.Context, string, string, authr.UserInfo) error) *Repository_RemoveTag_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
