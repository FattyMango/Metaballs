package main

import (
	"image/color"
	"sync"
)

type Group struct {
	balls []Ball
	mutex *sync.RWMutex
}

//	 n
//		∑     =  r(i)^2/ (x-x(i))^2 + (y-y(i))^2
//	 i=0
//
// This is the function of drawing a metaball
func (e *Group) value(x, y float32) float32 {
	// rl := e.mutex.RLocker()
	// rl.Lock()
	// defer rl.Unlock()

	var res float32
	for _, b := range e.balls {
		dx := x - b.pos.x
		dy := y - b.pos.y
		res += square(b.radius) / (square(dx) + square(dy))
	}

	return res
}

func (e *Group) valueV2(x1,x2,y1,y2 float32)(float32,float32,float32,float32) {
	// rl := e.mutex.RLocker()
	// rl.Lock()
	// defer rl.Unlock()

	var a,b,c,d float32
	for _, ball := range e.balls {
		a += circleFunc(ball.radius,ball.pos.x,ball.pos.y,x1, y1)
		b += circleFunc(ball.radius,ball.pos.x,ball.pos.y,x2, y1)
		c += circleFunc(ball.radius,ball.pos.x,ball.pos.y,x2, y2)
		d += circleFunc(ball.radius,ball.pos.x,ball.pos.y,x1, y2)
	}

	return a,b,c,d
}
func circleFunc (r,px,py, x,y float32) float32 {
	return square(r) / (square(x-px) + square(y-py))
}

// get the mix of colors for a given position
func (e *Group) color(x, y float32) color.Color {

	return color.White
}

func newRandomGroup(n int, s BallSpeed, z BallSize) *Group {
	balls := make([]Ball, n)
	for i := 0; i < n; i++ {
		balls[i] = NewRandomBall(s, z)
	}
	return &Group{
		balls: balls,
		mutex: &sync.RWMutex{},
	}
}

func (e *Group) move() {
	// e.mutex.Lock()
	// defer e.mutex.Unlock()
	for i := range e.balls {
		e.balls[i].move()
	}
}

func square(x float32) float32 {
	return x * x
}
