package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

var (
	globalResult     int
	globalResultChan = make(chan int, 100)
)

func srand(n int) []string {
	s := make([]string, n)
	for ind := range s {
		b := make([]byte, 16)
		_, err := rand.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		s[ind] = base64.URLEncoding.EncodeToString(b)
	}
	return s
}

func nrand(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func createCM(n int, cm *ConcurrentMap) []string {
	nums := nrand(n)
	strs := srand(n)
	for i, v := range nums {
		cm.Create(strs[i], worker{IP: strs[i], Tasks: v, MaxTasks: maxtasks, Port: port, CPU: cpu, Mem: mem, Disk: disk, Docker: dockerv, Go: gov, Group: group, LastJoin: join, Version: version})
	}
	return strs
}

func createSM(n int, sm *sync.Map) []string {
	nums := nrand(n)
	strs := srand(n)
	for i, v := range nums {
		sm.Store(strs[i], worker{IP: strs[i], Tasks: v, MaxTasks: maxtasks, Port: port, CPU: cpu, Mem: mem, Disk: disk, Docker: dockerv, Go: gov, Group: group, LastJoin: join, Version: version})
	}
	return strs
}

func BenchmarkCMGet_1(b *testing.B) {
	benchmarkCMGet(b, 1)
}

func BenchmarkCMGet_2(b *testing.B) {
	benchmarkCMGet(b, 2)
}

func BenchmarkCMGet_4(b *testing.B) {
	benchmarkCMGet(b, 4)
}

func BenchmarkCMGet_8(b *testing.B) {
	benchmarkCMGet(b, 8)
}

func BenchmarkCMGet_16(b *testing.B) {
	benchmarkCMGet(b, 16)
}

func BenchmarkCMGet_32(b *testing.B) {
	benchmarkCMGet(b, 32)
}

func BenchmarkCMGet_64(b *testing.B) {
	benchmarkCMGet(b, 64)
}

func benchmarkCMGet(b *testing.B, workerCount int) {
	numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(workerCount)

	cm = &ConcurrentMap{workerInfo: map[string]*worker{}}
	strs := createCM(b.N, cm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				if _, ok := cm.Get(strs[i]); ok {
					currentResult++
				}
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
	b.StopTimer()
	runtime.GOMAXPROCS(numcpus)
}

func BenchmarkSMLoad_1(b *testing.B) {
	benchmarkSMLoad(b, 1)
}

func BenchmarkSMLoad_2(b *testing.B) {
	benchmarkSMLoad(b, 2)
}

func BenchmarkSMLoad_4(b *testing.B) {
	benchmarkSMLoad(b, 4)
}

func BenchmarkSMLoad_8(b *testing.B) {
	benchmarkSMLoad(b, 8)
}

func BenchmarkSMLoad_16(b *testing.B) {
	benchmarkSMLoad(b, 16)
}

func BenchmarkSMLoad_32(b *testing.B) {
	benchmarkSMLoad(b, 32)
}

func BenchmarkSMLoad_64(b *testing.B) {
	benchmarkSMLoad(b, 64)
}

func benchmarkSMLoad(b *testing.B, workerCount int) {
	numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(workerCount)

	var sm sync.Map
	strs := createSM(b.N, &sm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				if _, ok := sm.Load(strs[i]); ok {
					currentResult++
				}
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
	b.StopTimer()
	runtime.GOMAXPROCS(numcpus)
}

func BenchmarkCMCreate_1(b *testing.B) {
	benchmarkCMCreate(b, 1)
}

func BenchmarkCMCreate_2(b *testing.B) {
	benchmarkCMCreate(b, 2)
}

func BenchmarkCMCreate_4(b *testing.B) {
	benchmarkCMCreate(b, 4)
}

func BenchmarkCMCreate_8(b *testing.B) {
	benchmarkCMCreate(b, 8)
}

func BenchmarkCMCreate_16(b *testing.B) {
	benchmarkCMCreate(b, 16)
}

func BenchmarkCMCreate_32(b *testing.B) {
	benchmarkCMCreate(b, 32)
}

func BenchmarkCMCreate_64(b *testing.B) {
	benchmarkCMCreate(b, 64)
}

func benchmarkCMCreate(b *testing.B, workerCount int) {
	numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(workerCount)

	cm = &ConcurrentMap{workerInfo: map[string]*worker{}}
	strs := srand(b.N)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				cm.Create(strs[i], worker{IP: strs[i], Tasks: i, MaxTasks: maxtasks, Port: port, CPU: cpu, Mem: mem, Disk: disk, Docker: dockerv, Go: gov, Group: group, LastJoin: join, Version: version})
				currentResult++
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
	b.StopTimer()
	runtime.GOMAXPROCS(numcpus)
}

func BenchmarkSMStore_1(b *testing.B) {
	benchmarkSMStore(b, 1)
}

func BenchmarkSMStore_2(b *testing.B) {
	benchmarkSMStore(b, 2)
}

func BenchmarkSMStore_4(b *testing.B) {
	benchmarkSMStore(b, 4)
}

func BenchmarkSMStore_8(b *testing.B) {
	benchmarkSMStore(b, 8)
}

func BenchmarkSMStore_16(b *testing.B) {
	benchmarkSMStore(b, 16)
}

func BenchmarkSMStore_32(b *testing.B) {
	benchmarkSMStore(b, 32)
}

func BenchmarkSMStore_64(b *testing.B) {
	benchmarkSMStore(b, 64)
}

func benchmarkSMStore(b *testing.B, workerCount int) {
	numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(workerCount)

	var sm sync.Map
	strs := srand(b.N)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				sm.Store(strs[i], worker{IP: strs[i], Tasks: i, MaxTasks: maxtasks, Port: port, CPU: cpu, Mem: mem, Disk: disk, Docker: dockerv, Go: gov, Group: group, LastJoin: join, Version: version})
				currentResult++
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
	b.StopTimer()
	runtime.GOMAXPROCS(numcpus)
}

func BenchmarkCMGetCreate_1(b *testing.B) {
	benchmarkCMGetCreate(b, 1)
}

func BenchmarkCMGetCreate_2(b *testing.B) {
	benchmarkCMGetCreate(b, 2)
}

func BenchmarkCMGetCreate_4(b *testing.B) {
	benchmarkCMGetCreate(b, 4)
}

func BenchmarkCMGetCreate_8(b *testing.B) {
	benchmarkCMGetCreate(b, 8)
}

func BenchmarkCMGetCreate_16(b *testing.B) {
	benchmarkCMGetCreate(b, 16)
}

func BenchmarkCMGetCreate_32(b *testing.B) {
	benchmarkCMGetCreate(b, 32)
}

func BenchmarkCMGetCreate_64(b *testing.B) {
	benchmarkCMGetCreate(b, 64)
}

func benchmarkCMGetCreate(b *testing.B, workerCount int) {
	numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(workerCount)

	cm = &ConcurrentMap{workerInfo: map[string]*worker{}}
	strs := createCM(b.N, cm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				if v, ok := cm.Get(strs[i]); ok {
					v.Tasks = i
					cm.Create(strs[i], v)
					currentResult++
				}
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
	b.StopTimer()
	runtime.GOMAXPROCS(numcpus)
}

func BenchmarkSMLoadStore_1(b *testing.B) {
	benchmarkSMLoadStore(b, 1)
}

func BenchmarkSMLoadStore_2(b *testing.B) {
	benchmarkSMLoadStore(b, 2)
}

func BenchmarkSMLoadStore_4(b *testing.B) {
	benchmarkSMLoadStore(b, 4)
}

func BenchmarkSMLoadStore_8(b *testing.B) {
	benchmarkSMLoadStore(b, 8)
}

func BenchmarkSMLoadStore_16(b *testing.B) {
	benchmarkSMLoadStore(b, 16)
}

func BenchmarkSMLoadStore_32(b *testing.B) {
	benchmarkSMLoadStore(b, 32)
}

func BenchmarkSMLoadStore_64(b *testing.B) {
	benchmarkSMLoadStore(b, 64)
}

func benchmarkSMLoadStore(b *testing.B, workerCount int) {
	numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(workerCount)

	var sm sync.Map
	strs := createSM(b.N, &sm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				sm.LoadOrStore(strs[i], worker{Tasks: i})
				currentResult++
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
	b.StopTimer()
	runtime.GOMAXPROCS(numcpus)
}

func BenchmarkCMDelete_1(b *testing.B) {
	benchmarkCMDelete(b, 1)
}

func BenchmarkCMDelete_2(b *testing.B) {
	benchmarkCMDelete(b, 2)
}

func BenchmarkCMDelete_4(b *testing.B) {
	benchmarkCMDelete(b, 4)
}

func BenchmarkCMDelete_8(b *testing.B) {
	benchmarkCMDelete(b, 8)
}

func BenchmarkCMDelete_16(b *testing.B) {
	benchmarkCMDelete(b, 16)
}

func BenchmarkCMDelete_32(b *testing.B) {
	benchmarkCMDelete(b, 32)
}

func BenchmarkCMDelete_64(b *testing.B) {
	benchmarkCMDelete(b, 64)
}

func benchmarkCMDelete(b *testing.B, workerCount int) {
	numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(workerCount)

	cm = &ConcurrentMap{workerInfo: map[string]*worker{}}
	strs := createCM(b.N, cm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				cm.Delete(strs[i])
				currentResult++
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
	b.StopTimer()
	runtime.GOMAXPROCS(numcpus)
}

func BenchmarkSMDelete_1(b *testing.B) {
	benchmarkSMDelete(b, 1)
}

func BenchmarkSMDelete_2(b *testing.B) {
	benchmarkSMDelete(b, 2)
}

func BenchmarkSMDelete_4(b *testing.B) {
	benchmarkSMDelete(b, 4)
}

func BenchmarkSMDelete_8(b *testing.B) {
	benchmarkSMDelete(b, 8)
}

func BenchmarkSMDelete_16(b *testing.B) {
	benchmarkSMDelete(b, 16)
}

func BenchmarkSMDelete_32(b *testing.B) {
	benchmarkSMDelete(b, 32)
}

func BenchmarkSMDelete_64(b *testing.B) {
	benchmarkSMDelete(b, 64)
}

func benchmarkSMDelete(b *testing.B, workerCount int) {
	numcpus := runtime.NumCPU()
	runtime.GOMAXPROCS(workerCount)

	var sm sync.Map
	strs := createSM(b.N, &sm)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	globalResultChan = make(chan int, workerCount)

	b.ResetTimer()

	for wc := 0; wc < workerCount; wc++ {
		go func(n int) {
			currentResult := 0
			for i := 0; i < n; i++ {
				sm.Delete(strs[i])
				currentResult++
			}
			globalResultChan <- currentResult
			wg.Done()
		}(b.N)
	}

	wg.Wait()
	b.StopTimer()
	runtime.GOMAXPROCS(numcpus)
}
