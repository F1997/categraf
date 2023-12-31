//go:build !no_logs

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package providers

import (
	"log"

	"github.com/F1997/categraf/logs/util/containers"
)

// ContainerImpl without implementation
// Implementations should call Register() in their init()
var containerImpl containers.ContainerImplementation

// ContainerImpl returns the ContainerImplementation
func ContainerImpl() containers.ContainerImplementation {
	if containerImpl == nil {
		panic("Trying to get nil ContainerInterface")
	}

	return containerImpl
}

// Register allows to set a ContainerImplementation
func Register(impl containers.ContainerImplementation) {
	if containerImpl == nil {
		containerImpl = impl
	} else {
		log.Printf("Trying to set multiple ContainerImplementation")
	}
}

// Deregister allows to unset a ContainerImplementation
// this should only be used in tests to clean the global state
func Deregister() {
	containerImpl = nil
}
