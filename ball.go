package main

import "math/rand"

type Position struct {
	x float32
	y float32
}
func NewrandomPosition() Position {
	return Position{
		x: rand.Float32()*0.8 + 0.1,
		y: rand.Float32()*0.8 + 0.1,
	}
}

type Velocity struct {
	x float32
	y float32
}
func NewVelocity() Velocity {
	return Velocity{
		x: (rand.Float32() - 0.5) * 0.01,
		y: (rand.Float32() - 0.5) * 0.01,
	}
}
type Ball struct {
	pos Position
	vel Velocity
	radius  float32

}

func NewRandomBall() Ball {
	return Ball{
		pos: NewrandomPosition(),
		radius:  rand.Float32()*0.1 + 0.025,
		vel: NewVelocity(),
	}
}

func (b *Ball) move() {
	b.pos.x += b.vel.x
	if b.pos.x <= 0.1 || b.pos.x >= 0.85 {
		b.vel.x = -b.vel.x
	}
	b.pos.y  += b.vel.y
	if b.pos.y  <= 0.1 || b.pos.y   >= 0.85 {
		b.vel.y  = -b.vel.y 
	}
}
