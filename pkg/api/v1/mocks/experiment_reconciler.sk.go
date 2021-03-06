// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/v1/experiment_reconciler.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/glooshot/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// MockExperimentReconciler is a mock of ExperimentReconciler interface
type MockExperimentReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockExperimentReconcilerMockRecorder
}

// MockExperimentReconcilerMockRecorder is the mock recorder for MockExperimentReconciler
type MockExperimentReconcilerMockRecorder struct {
	mock *MockExperimentReconciler
}

// NewMockExperimentReconciler creates a new mock instance
func NewMockExperimentReconciler(ctrl *gomock.Controller) *MockExperimentReconciler {
	mock := &MockExperimentReconciler{ctrl: ctrl}
	mock.recorder = &MockExperimentReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExperimentReconciler) EXPECT() *MockExperimentReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method
func (m *MockExperimentReconciler) Reconcile(namespace string, desiredResources v1.ExperimentList, transition v1.TransitionExperimentFunc, opts clients.ListOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", namespace, desiredResources, transition, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockExperimentReconcilerMockRecorder) Reconcile(namespace, desiredResources, transition, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockExperimentReconciler)(nil).Reconcile), namespace, desiredResources, transition, opts)
}
