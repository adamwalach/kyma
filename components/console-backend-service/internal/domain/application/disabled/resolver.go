// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import (
	context "context"

	gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
)

// Resolver is an autogenerated failing mock type for the Resolver type
type Resolver struct {
	err error
}

// NewResolver creates a new Resolver type instance
func NewResolver(err error) *Resolver {
	return &Resolver{err: err}
}

// ApplicationEnabledInNamespacesField provides a failing mock function with given fields: ctx, obj
func (_m *Resolver) ApplicationEnabledInNamespacesField(ctx context.Context, obj *gqlschema.Application) ([]string, error) {
	var r0 []string
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ApplicationEnabledMappingServices provides a failing mock function with given fields: ctx, obj
func (_m *Resolver) ApplicationEnabledMappingServices(ctx context.Context, obj *gqlschema.Application) ([]*gqlschema.EnabledMappingService, error) {
	var r0 []*gqlschema.EnabledMappingService
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ApplicationEventSubscription provides a failing mock function with given fields: ctx
func (_m *Resolver) ApplicationEventSubscription(ctx context.Context) (<-chan *gqlschema.ApplicationEvent, error) {
	var r0 <-chan *gqlschema.ApplicationEvent
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ApplicationQuery provides a failing mock function with given fields: ctx, name
func (_m *Resolver) ApplicationQuery(ctx context.Context, name string) (*gqlschema.Application, error) {
	var r0 *gqlschema.Application
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ApplicationStatusField provides a failing mock function with given fields: ctx, app
func (_m *Resolver) ApplicationStatusField(ctx context.Context, app *gqlschema.Application) (gqlschema.ApplicationStatus, error) {
	var r0 gqlschema.ApplicationStatus
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ApplicationsQuery provides a failing mock function with given fields: ctx, namespace, first, offset
func (_m *Resolver) ApplicationsQuery(ctx context.Context, namespace *string, first *int, offset *int) ([]*gqlschema.Application, error) {
	var r0 []*gqlschema.Application
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ConnectorServiceQuery provides a failing mock function with given fields: ctx, _a1
func (_m *Resolver) ConnectorServiceQuery(ctx context.Context, _a1 string) (*gqlschema.ConnectorService, error) {
	var r0 *gqlschema.ConnectorService
	var r1 error
	r1 = _m.err

	return r0, r1
}

// CreateApplication provides a failing mock function with given fields: ctx, name, description, qglLabels
func (_m *Resolver) CreateApplication(ctx context.Context, name string, description *string, qglLabels gqlschema.Labels) (*gqlschema.ApplicationMutationOutput, error) {
	var r0 *gqlschema.ApplicationMutationOutput
	var r1 error
	r1 = _m.err

	return r0, r1
}

// DeleteApplication provides a failing mock function with given fields: ctx, name
func (_m *Resolver) DeleteApplication(ctx context.Context, name string) (*gqlschema.DeleteApplicationOutput, error) {
	var r0 *gqlschema.DeleteApplicationOutput
	var r1 error
	r1 = _m.err

	return r0, r1
}

// DisableApplicationMutation provides a failing mock function with given fields: ctx, _a1, namespace
func (_m *Resolver) DisableApplicationMutation(ctx context.Context, _a1 string, namespace string) (*gqlschema.ApplicationMapping, error) {
	var r0 *gqlschema.ApplicationMapping
	var r1 error
	r1 = _m.err

	return r0, r1
}

// EnableApplicationMutation provides a failing mock function with given fields: ctx, _a1, namespace, allServices, services
func (_m *Resolver) EnableApplicationMutation(ctx context.Context, _a1 string, namespace string, allServices *bool, services []*gqlschema.ApplicationMappingService) (*gqlschema.ApplicationMapping, error) {
	var r0 *gqlschema.ApplicationMapping
	var r1 error
	r1 = _m.err

	return r0, r1
}

// EventActivationEventsField provides a failing mock function with given fields: ctx, eventActivation
func (_m *Resolver) EventActivationEventsField(ctx context.Context, eventActivation *gqlschema.EventActivation) ([]*gqlschema.EventActivationEvent, error) {
	var r0 []*gqlschema.EventActivationEvent
	var r1 error
	r1 = _m.err

	return r0, r1
}

// EventActivationsQuery provides a failing mock function with given fields: ctx, namespace
func (_m *Resolver) EventActivationsQuery(ctx context.Context, namespace string) ([]*gqlschema.EventActivation, error) {
	var r0 []*gqlschema.EventActivation
	var r1 error
	r1 = _m.err

	return r0, r1
}

// OverloadApplicationMutation provides a failing mock function with given fields: ctx, _a1, namespace, allServices, services
func (_m *Resolver) OverloadApplicationMutation(ctx context.Context, _a1 string, namespace string, allServices *bool, services []*gqlschema.ApplicationMappingService) (*gqlschema.ApplicationMapping, error) {
	var r0 *gqlschema.ApplicationMapping
	var r1 error
	r1 = _m.err

	return r0, r1
}

// UpdateApplication provides a failing mock function with given fields: ctx, name, description, qglLabels
func (_m *Resolver) UpdateApplication(ctx context.Context, name string, description *string, qglLabels gqlschema.Labels) (*gqlschema.ApplicationMutationOutput, error) {
	var r0 *gqlschema.ApplicationMutationOutput
	var r1 error
	r1 = _m.err

	return r0, r1
}
