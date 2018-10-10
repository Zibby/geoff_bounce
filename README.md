# Install

git clone

## On Ubuntu 14.04 and above, type:
`apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev golang`

```
go get -v github.com/veandco/go-sdl2/sdl
go get -v github.com/veandco/go-sdl2/img
go get -v github.com/veandco/go-sdl2/mix
go get -v github.com/veandco/go-sdl2/ttf
go get -v github.com/veandco/go-sdl2/gfx
go build .
```
## On osx

`brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config`

```
go get -v github.com/veandco/go-sdl2/sdl
go get -v github.com/veandco/go-sdl2/img
go get -v github.com/veandco/go-sdl2/mix
go get -v github.com/veandco/go-sdl2/ttf
go get -v github.com/veandco/go-sdl2/gfx
go build .
```

