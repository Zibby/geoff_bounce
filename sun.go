package main

import (
	"fmt"
	"math/rand"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	sunSpeedMax = 5
	sunSpeedMin = 1
)

func sunSpeed() float64 {
	speed := (sunSpeedMin + rand.Float64()*(sunSpeedMax-sunSpeedMin))
	fmt.Println(speed)
	return speed
}

type sun struct {
	time     int
	textures []*sdl.Texture
	x, speed float64
}

func newSun(r *sdl.Renderer) (*sun, error) {
	var textures []*sdl.Texture
	for i := 2; i <= 2; i++ {
		path := fmt.Sprintf("res/images/sun-%d.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return nil, fmt.Errorf("Could not load texture: %v", err)
		}
		textures = append(textures, texture)
	}
	return &sun{textures: textures, x: 760, speed: sunSpeed()}, nil
}

func (sun *sun) paint(r *sdl.Renderer) error {
	sun.time++
	sun.x -= float64(sun.speed)
	if sun.x < 0 {
		sun.x = 760
		sun.speed = sunSpeed()
	}
	rect := &sdl.Rect{X: int32(sun.x), Y: 430, W: 80, H: 80}
	if err := r.Copy(sun.textures[0], nil, rect); err != nil {
		return fmt.Errorf("Could not copy sun %v", err)
	}
	return nil
}

func (sun *sun) destroy() {
	for _, sun := range sun.textures {
		sun.Destroy()
	}
}
