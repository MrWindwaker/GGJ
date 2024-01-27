package game

import rl "github.com/gen2brain/raylib-go/raylib"

type NPC struct {
	active bool

	width  int32
	height int32

	txt      rl.Texture2D
	c_frame  int
	t_frames int

	pos rl.Vector2
}

func (n *NPC) set_values(file string, total_frames int) {
	n.txt = rl.LoadTexture(file)

	n.width = n.txt.Width / 14
	n.height = 32

	n.c_frame = 0
	n.t_frames = total_frames

	n.pos = rl.NewVector2(0, 0)
	n.active = true
}

func (n *NPC) animate() {
	n.c_frame++

	if n.c_frame >= n.t_frames {
		n.c_frame = 0
	}
}

func (n *NPC) get_source() rl.Rectangle {
	return rl.NewRectangle(
		float32(int32(n.c_frame)*n.width),
		0,
		float32(n.width),
		float32(n.height),
	)
}

func (n *NPC) get_dest() rl.Rectangle {
	return rl.NewRectangle(
		n.pos.X,
		n.pos.Y,
		float32(n.width)*4,
		float32(n.height)*4,
	)
}

func (n *NPC) draw() {
	if n.active {
		rl.DrawTexturePro(
			n.txt,
			n.get_source(),
			n.get_dest(),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)
	}
}

func (n *NPC) close() {
	rl.UnloadTexture(n.txt)
}
