// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
	mock "github.com/stretchr/testify/mock"

	v1beta1 "github.com/kyma-project/rafter/pkg/apis/rafter/v1beta1"
)

// gqlAssetGroupConverter is an autogenerated mock type for the gqlAssetGroupConverter type
type gqlAssetGroupConverter struct {
	mock.Mock
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlAssetGroupConverter) ToGQL(in *v1beta1.AssetGroup) (*gqlschema.AssetGroup, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.AssetGroup
	if rf, ok := ret.Get(0).(func(*v1beta1.AssetGroup) *gqlschema.AssetGroup); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.AssetGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.AssetGroup) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlAssetGroupConverter) ToGQLs(in []*v1beta1.AssetGroup) ([]gqlschema.AssetGroup, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.AssetGroup
	if rf, ok := ret.Get(0).(func([]*v1beta1.AssetGroup) []gqlschema.AssetGroup); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.AssetGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1beta1.AssetGroup) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
