// Copyright (c) 2021 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Eclipse Public License 2.0 which is available at
// https://www.eclipse.org/legal/epl-2.0, or the Apache License, Version 2.0
// which is available at https://www.apache.org/licenses/LICENSE-2.0.
//
// SPDX-License-Identifier: EPL-2.0 OR Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: containerm/events/events_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	types "github.com/eclipse-kanto/container-management/containerm/containers/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockContainerEventsManager is a mock of ContainerEventsManager interface
type MockContainerEventsManager struct {
	ctrl     *gomock.Controller
	recorder *MockContainerEventsManagerMockRecorder
}

// MockContainerEventsManagerMockRecorder is the mock recorder for MockContainerEventsManager
type MockContainerEventsManagerMockRecorder struct {
	mock *MockContainerEventsManager
}

// NewMockContainerEventsManager creates a new mock instance
func NewMockContainerEventsManager(ctrl *gomock.Controller) *MockContainerEventsManager {
	mock := &MockContainerEventsManager{ctrl: ctrl}
	mock.recorder = &MockContainerEventsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContainerEventsManager) EXPECT() *MockContainerEventsManagerMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockContainerEventsManager) Publish(ctx context.Context, eventType types.EventType, eventAction types.EventAction, source *types.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, eventType, eventAction, source)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockContainerEventsManagerMockRecorder) Publish(ctx, eventType, eventAction, source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockContainerEventsManager)(nil).Publish), ctx, eventType, eventAction, source)
}

// Subscribe mocks base method
func (m *MockContainerEventsManager) Subscribe(ctx context.Context) (<-chan *types.Event, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx)
	ret0, _ := ret[0].(<-chan *types.Event)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockContainerEventsManagerMockRecorder) Subscribe(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockContainerEventsManager)(nil).Subscribe), ctx)
}
