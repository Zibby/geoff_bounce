package main

import (
	"fmt"
	"math"
	"os"
	"time"

	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	time  int
	bg    *sdl.Texture
	geoff *geoff
	sun   *sun
	dead  bool
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/images/back.png")
	if err != nil {
		return nil, fmt.Errorf("Could not drawBackground: %v", err)
	}
	geoff, err := newGeoff(r)
	if err != nil {
		return nil, err
	}

	sun, err := newSun(r)
	if err != nil {
		return nil, err
	}
	return &scene{bg: bg, geoff: geoff, sun: sun}, nil
}

func (s *scene) run(events <-chan sdl.Event, r *sdl.Renderer) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		tick := time.Tick(10 * time.Millisecond)
		for {
			select {
			case e := <-events:
				if done := s.handleEvent(e); done {
					return
				}
				//log.Printf("event: %T", e)
			case <-tick:
				if err := s.paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()
	return errc
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.MouseButtonEvent:
		s.geoff.jump()
	case *sdl.MouseMotionEvent, *sdl.WindowEvent, *sdl.TouchFingerEvent:
	default:
		//log.Printf("unknown event %T", event)
	}
	return false
}

func (s *scene) dieded(r *sdl.Renderer) {
	fmt.Fprintf(os.Stderr, "YOU DIED")
	s.geoff.destroy()
	s.sun.destroy()
	s.bg.Destroy()
	s.dead = true
	drawTitle(r, "Game Over", true)
}

func (s *scene) paint(r *sdl.Renderer) error {
	s.time++
	r.Clear()
	if !s.dead {
		if err := r.Copy(s.bg, nil, nil); err != nil {
			return fmt.Errorf("Could not copy background: %v", err)
		}

		if err := s.sun.paint(r); err != nil {
			return err
		}

		if err := s.geoff.paint(r); err != nil {
			return err
		}

		r.Present()
	}

	go func() {
		//fmt.Printf("%v", s.geoff.y)
		//fmt.Fprintf(os.Stderr, "\n")
		//fmt.Println(s.geoff.x, s.sun.x)
		diff := s.sun.x - s.geoff.x
		magdiff := math.Sqrt(diff * diff)
		//fmt.Println(magdiff)

		if magdiff <= 10 {
			if s.geoff.y <= 150 {
				s.dieded(r)
				return
			}
			return
		}
	}()
	return nil
}

func (s *scene) gameover() {
	s.geoff.destroy()
}

func (s *scene) destroy() {
	s.bg.Destroy()
	s.geoff.destroy()
	s.sun.destroy()
}
