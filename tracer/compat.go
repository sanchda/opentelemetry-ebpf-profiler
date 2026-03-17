// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracer // import "go.opentelemetry.io/ebpf-profiler/tracer"

import "go.opentelemetry.io/ebpf-profiler/internal/linux"

// ProbeBPFSyscall preserves the older exported package API expected by downstream consumers.
func ProbeBPFSyscall() error {
	return linux.ProbeBPFSyscall()
}

// GetCurrentKernelVersion preserves the older exported package API expected by downstream consumers.
func GetCurrentKernelVersion() (major, minor, patch uint32, err error) {
	return linux.GetCurrentKernelVersion()
}
