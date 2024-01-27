package game

import rl "github.com/gen2brain/raylib-go/raylib"

type GAG interface {
	load()
	animate()
	play()
	unload()
	draw()
}

type Explotion_GAG struct {
	txt    rl.Texture2D
	width  int32
	height int32
	pos    rl.Vector2

	t_frames int
	c_frame  int

	sfx rl.Sound

	active bool
}

func (ex *Explotion_GAG) load() {
	ex.txt = rl.LoadTexture("Assets/Images/VFXs/character_explotion.png")
	ex.sfx = rl.LoadSound("Assets/Sounds/VFXs/explosion.mp3")

	ex.t_frames = 12
	ex.width = ex.txt.Width / int32(ex.t_frames)
	ex.height = ex.txt.Height

	ex.active = false
}

func (ex *Explotion_GAG) animate() {
	if ex.active {
		ex.c_frame++
		if ex.c_frame >= ex.t_frames {
			ex.active = false
			ex.c_frame = 0
		}
	}
}

func (ex *Explotion_GAG) play(pos rl.Vector2) {
	ex.active = true
	ex.pos = pos
	rl.PlaySound(ex.sfx)
}

func (ex *Explotion_GAG) unload() {
	rl.UnloadTexture(ex.txt)
	rl.UnloadSound(ex.sfx)
}

func (ex *Explotion_GAG) draw() {
	if ex.active {
		rl.DrawTexturePro(
			ex.txt,
			rl.NewRectangle(float32(int32(ex.c_frame)*ex.width), 0, float32(ex.width), float32(ex.height)),
			rl.NewRectangle(ex.pos.X, ex.pos.Y, float32(ex.width), float32(ex.height)),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)
	}
}
