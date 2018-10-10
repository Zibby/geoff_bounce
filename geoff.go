package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gravity   = 0.25
	jumpSpeed = -10
)

type geoff struct {
	time        int
	textures    []*sdl.Texture
	x, y, speed float64
	jumplevel   int
}

func newGeoff(r *sdl.Renderer) (*geoff, error) {
	var textures []*sdl.Texture
	for i := 1; i <= 17; i++ {
		path := fmt.Sprintf("res/images/frame_%02d_delay-0.05s.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return nil, fmt.Errorf("Could not load texture: %v", err)
		}
		textures = append(textures, texture)
	}
	return &geoff{textures: textures, y: 140, speed: 0, jumplevel: 0, x: 300}, nil
}

func (geoff *geoff) paint(r *sdl.Renderer) error {
	geoff.time++
	geoff.y -= geoff.speed
	if geoff.y < 140 {
		geoff.speed = 0
		geoff.jumplevel = 0
		geoff.y = 140
	}

	if geoff.y > 500 {
		geoff.speed = 0
		geoff.y = 500
	}
	geoff.speed += gravity
	rect := &sdl.Rect{X: 300, Y: (600 - int32(geoff.y)) - 80/2, W: 100, H: 100}
	i := geoff.time / 5 % len(geoff.textures)
	if err := r.Copy(geoff.textures[i], nil, rect); err != nil {
		return fmt.Errorf("Could not copy geoff: %v", err)
	}
	return nil
}

func (geoff *geoff) destroy() {
	for _, geoff := range geoff.textures {
		geoff.Destroy()
	}
}

func (geoff *geoff) jump() {
	geoff.jumplevel++
	if geoff.jumplevel <= 3 {
		geoff.speed = jumpSpeed
	}
}
