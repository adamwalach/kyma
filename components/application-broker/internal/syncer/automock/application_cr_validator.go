// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"

import v1alpha1 "github.com/kyma-project/kyma/components/application-operator/pkg/apis/applicationconnector/v1alpha1"

// ApplicationCRValidator is an autogenerated mock type for the applicationCRValidator type
type ApplicationCRValidator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: dto
func (_m *ApplicationCRValidator) Validate(dto *v1alpha1.Application) error {
	ret := _m.Called(dto)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Application) error); ok {
		r0 = rf(dto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
