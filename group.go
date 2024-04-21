package main

import "sync"

type Group struct {
	balls []Ball
	mutex sync.RWMutex
}

func (e *Group) value(x, y float32) float32 {
	rl := e.mutex.RLocker()
	rl.Lock()
	defer rl.Unlock()
	var res float32
	for _, b := range e.balls {
		dx := x - b.pos.x
		dy := y - b.pos.y
		res += b.radius * b.radius / (dx*dx + dy*dy)
	}
	return res
}

func newRandomEnsemble(n int) *Group {
	balls := make([]Ball, n)
	for i := 0; i < n; i++ {
		balls[i] = NewRandomBall()
	}
	return &Group{balls: balls}
}

func (e *Group) move() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for i := range e.balls {
		e.balls[i].move()
	}
}
