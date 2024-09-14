//go:build !noasm
// +build !noasm

package utf8

import (
	"github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services/simd/golib/cpu"
)

func init() {
	if cpu.X86.HasAVX2 {
		validFn = validate_utf8_fast_avx2
	} else if cpu.X86.HasSSE42 {
		validFn = validate_utf8_fast_sse4
	} else {
		validFn = validate_utf8_go
	}
}
