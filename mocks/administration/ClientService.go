// Code generated by mockery v2.9.4. DO NOT EDIT.

package administrationmock

import (
	administration "github.com/recolabs/reco/redash-client/gen/client/administration"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// GetPing provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetPing(params *administration.GetPingParams, authInfo runtime.ClientAuthInfoWriter, opts ...administration.ClientOption) (*administration.GetPingOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *administration.GetPingOK
	if rf, ok := ret.Get(0).(func(*administration.GetPingParams, runtime.ClientAuthInfoWriter, ...administration.ClientOption) *administration.GetPingOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*administration.GetPingOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*administration.GetPingParams, runtime.ClientAuthInfoWriter, ...administration.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}