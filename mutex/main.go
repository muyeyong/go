package main

import (
	"sync"
	"time"
)

type SafeNum struct {
	v map[string]int
	m sync.Mutex
}

func (s *SafeNum) Inc(key string) {
	// s.m.Lock()
	s.v[key] ++
	// s.m.Unlock()
}

func (s *SafeNum) Value(key string) int {
	// s.m.Lock()
	// defer s.m.Unlock()
	return s.v[key]
}

func main() {
	num := SafeNum{v: make(map[string]int)}
	for i :=0; i < 1000; i++ {
		go num.Inc("a")
	}
	time.Sleep(time.Second)
	// fmt.Println(num.Value("a"))
}