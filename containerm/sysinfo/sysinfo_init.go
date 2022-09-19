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

package sysinfo

import (
	"github.com/eclipse-kanto/container-management/containerm/registry"
	"github.com/eclipse-kanto/container-management/containerm/sysinfo/types"
	"github.com/eclipse-kanto/container-management/containerm/version"
)

const (
	// SystemInfoServiceLocalID is the ID of the local service system info
	SystemInfoServiceLocalID = "container-management.service.local.v1.service-system-info"
)

func init() {
	registry.Register(&registry.Registration{
		ID:       SystemInfoServiceLocalID,
		Type:     registry.SystemInfoService,
		InitFunc: registryInit,
	})
}

func registryInit(registryCtx *registry.ServiceRegistryContext) (interface{}, error) {
	//create system info service instance
	return newSystemInfoMgr(types.ProjectInfo{
		ProjectVersion: version.ProjectVersion,
		BuildTime:      version.BuildTime,
		APIVersion:     version.APIVersion,
		GitCommit:      version.GitCommit,
	}), nil
}
