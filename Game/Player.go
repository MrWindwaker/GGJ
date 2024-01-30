package game

import (
	"fmt"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	total int
	row   int
}

var Animations map[string]Animation = map[string]Animation{
	"Idle": {row: 0, total: 8},
	"Walk": {row: 1, total: 4},
}

type Player struct {
	pos rl.Vector2

	// width  int32
	// height int32

	txt       rl.Texture2D
	c_frame   int
	t_frame   int
	t_row     int
	direction int

	is_loaded bool
	is_moving bool
}

var pl_lock = &sync.Mutex{}
var pl_insta *Player

func Get_Player() *Player {
	if pl_insta == nil {
		pl_lock.Lock()
		defer pl_lock.Unlock()

		if pl_insta == nil {
			pl_insta = &Player{
				pos:       rl.Vector2{X: 0, Y: 0},
				is_loaded: false,
				c_frame:   0,
				t_frame:   8,
				t_row:     0,
				direction: 1,
				is_moving: false,
			}
		}
	}

	return pl_insta
}

func (p *Player) animate() {
	p.c_frame++

	if p.c_frame >= p.t_frame {
		p.c_frame = 0
	}
}

func (p *Player) change_animation(anim string) {
	an := Animations[anim]

	p.t_frame = an.total
	p.t_row = an.row

	if p.c_frame >= p.t_frame {
		p.c_frame = 0
	}
}

func (p *Player) get_source() rl.Rectangle {
	return rl.NewRectangle(
		float32(p.c_frame)*32,
		float32(p.t_row)*32,
		float32(32*p.direction),
		32,
	)
}

func (p *Player) get_dest() rl.Rectangle {
	return rl.NewRectangle(
		p.pos.X,
		p.pos.Y,
		32*4, 32*4,
	)
}

func (p *Player) load(pos rl.Vector2) {
	p.is_loaded = true
	p.txt = rl.LoadTexture("Assets/Images/Player.png")
	fmt.Println("Player Loaded")
	p.pos = pos
}

func (p *Player) update() {
	dt := rl.GetFrameTime()

	if rl.IsKeyDown(rl.KeyA) {
		p.pos.X -= 150 * dt
		p.change_animation("Walk")
		p.direction = -1

		p.is_moving = true
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.pos.X += 150 * dt
		p.change_animation("Walk")
		p.direction = 1

		p.is_moving = true
	}

	if !p.is_moving {
		p.change_animation("Idle")
	}

	if rl.IsKeyUp(rl.KeyA | rl.KeyD) {
		p.is_moving = false
	}
}

func (p *Player) draw() {
	rl.DrawTexturePro(
		p.txt,
		p.get_source(),
		p.get_dest(),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

func (p *Player) unload() {
	fmt.Println("Player Unloaded")
	p.is_loaded = false
	rl.UnloadTexture(p.txt)
}
