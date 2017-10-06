package main

import (
	"fmt"
	"sync"
)

var (
	cm       *ConcurrentMap
	sm       sync.Map
	ip       = "127.0.0.1"
	tasks    = 0
	maxtasks = 3
	port     = 8020
	mem      = 80
	disk     = 30
	cpu      = 20
	dockerv  = "1.7.0"
	gov      = "1.9"
	group    = "default"
	join     = "yobot"
	version  = "1.0.0"
)

type worker struct {
	IP       string
	MaxTasks int
	Tasks    int
	Port     int
	CPU      int
	Mem      int
	Disk     int
	Docker   string
	Go       string
	Group    string
	LastJoin string
	Version  string
}

type ConcurrentMap struct {
	sync.RWMutex
	workerInfo map[string]*worker
}

func (cm *ConcurrentMap) Create(ip string, w worker) {
	cm.Lock()
	defer cm.Unlock()
	cm.workerInfo[ip] = &w
}

func (cm *ConcurrentMap) Delete(ip string) {
	cm.Lock()
	defer cm.Unlock()
	delete(cm.workerInfo, ip)
}

func (cm *ConcurrentMap) Get(ip string) (worker, bool) {
	cm.RLock()
	defer cm.RUnlock()
	val, ok := cm.workerInfo[ip]
	if !ok {
		return worker{}, ok
	}
	return *val, ok
}

func main() {

	cm = &ConcurrentMap{workerInfo: map[string]*worker{}}
	cm.Create(ip, worker{IP: ip, Tasks: tasks, MaxTasks: maxtasks, Port: port, CPU: cpu, Mem: mem, Disk: disk, Docker: dockerv, Go: gov, Group: group, LastJoin: join, Version: version})
	sm.Store(ip, worker{IP: ip, Tasks: tasks, MaxTasks: maxtasks, Port: port, CPU: cpu, Mem: mem, Disk: disk, Docker: dockerv, Go: gov, Group: group, LastJoin: join, Version: version})

	fmt.Println(cm.Get(ip))
	fmt.Println(sm.Load(ip))
}
