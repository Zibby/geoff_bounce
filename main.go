package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not init SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("could not init TTF: %v", err)
	}

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window %v", err)
	}
	defer w.Destroy()
	_ = r

	if err := drawTitle(r, "Geoff Bounce", false); err != nil {
		return fmt.Errorf("Could not draw title: %v", err)
	}

	s, err := newScene(r)
	if err != nil {
		return fmt.Errorf("Could not create scene: %v", err)
	}
	defer s.destroy()

	events := make(chan sdl.Event)
	errc := s.run(events, r)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}

func drawTitle(r *sdl.Renderer, text string, dead bool) error {
	r.Clear()

	f, err := ttf.OpenFont("res/fonts/joystix.ttf", 20)
	if err != nil {
		return fmt.Errorf("could not open front: %v", err)
	}
	defer f.Close()

	c := sdl.Color{R: 71, G: 150, B: 235, A: 225}
	s, err := f.RenderUTF8Solid(text, c)
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()

	time.Sleep(4 * time.Second)
	if dead {
		return fmt.Errorf("Thats all folks")
	}
	return nil
}
