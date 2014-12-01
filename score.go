package montecarlo

import "sync"

type Score struct {
  m sync.Mutex
  x int64
}

func (s *Score) Add(x int64) {
  s.m.Lock()
  defer s.m.Unlock()
  s.x += x
}

func (s *Score) Value() (x int64) {
    s.m.Lock()
    defer s.m.Unlock()
    x = s.x
    return
}