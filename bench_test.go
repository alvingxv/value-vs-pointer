package main

import (
	"go-mem-profile/repo"
	"go-mem-profile/service"
	"os"
	"runtime/pprof"
	"testing"
)

func BenchmarkGetDataPointer(b *testing.B) {
	repository := repo.NewRepository(nil)
	svc := service.NewService(repository)

	// Start CPU profiling
	f, err := os.Create("mem1.prof")
	if err != nil {
		b.Error(err)
	}
	defer f.Close()
	defer pprof.WriteHeapProfile(f)
	cpuFile, err := os.Create("cpu1.prof")
	if err != nil {
		b.Error(err)
	}

	defer cpuFile.Close()
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		svc.GetDataPointer()
	}
}

func BenchmarkGetDataValue(b *testing.B) {
	repository := repo.NewRepository(nil)
	svc := service.NewService(repository)

	// Start CPU profiling
	f, err := os.Create("mem2.prof")
	if err != nil {
		b.Error(err)
	}
	defer f.Close()
	defer pprof.WriteHeapProfile(f)

	// Start CPU profiling
	cpuFile, err := os.Create("cpu2.prof")
	if err != nil {
		b.Error(err)
	}
	defer cpuFile.Close()
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		svc.GetDataValue()
	}
}
