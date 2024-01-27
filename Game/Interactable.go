package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Interactable struct {
	active bool

	active_key int

	width  int
	height int
	pos    rl.Vector2

	script func()
}

func (i *Interactable) run() {
	if i.active {
		i.script()
	}
}

func (i *Interactable) update(pl Player) {
	if rl.CheckCollisionRecs(i.get_rec(), pl.get_dest()) {
		i.active = true
	} else {
		i.active = false
	}

	if i.active && rl.IsKeyPressed(int32(i.active_key)) {
		i.run()
	}
}

func (i *Interactable) get_rec() rl.Rectangle {
	return rl.NewRectangle(
		i.pos.X,
		i.pos.Y,
		float32(i.width),
		float32(i.height),
	)
}
