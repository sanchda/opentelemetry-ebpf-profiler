// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host // import "go.opentelemetry.io/ebpf-profiler/host"

import "go.opentelemetry.io/ebpf-profiler/libpf"

// Trace preserves the older exported name expected by downstream consumers.
type Trace = libpf.EbpfTrace
