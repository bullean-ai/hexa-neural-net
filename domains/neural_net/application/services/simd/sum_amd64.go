//go:build !noasm
// +build !noasm

package simd

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/simd/golib/cpu"
)

func init() {
	if cpu.X86.HasAVX2 {
		sumFloat64 = sum_float64_avx2
	} else if cpu.X86.HasSSE42 {
		sumFloat64 = sum_float64_sse4
	} else {
		sumFloat64 = sum_float64_go
	}
}
