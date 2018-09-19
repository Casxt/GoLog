package GoLog

import (
	"testing"
)

func Test_service(t *testing.T) {
	StartLogger()
	Info("Info")
	Warn("Warn")
	Debug("Debug")
	Error("Error")
	StopLogger()
}

func BenchmarkInfo(b *testing.B) {
	StartLogger()
	b.ResetTimer()
	b.N = 40000
	for i := 0; i < b.N; i++ {
		Info("Info")
	}
	StopLogger()
}
