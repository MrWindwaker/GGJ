package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Room struct {
	gag GAG
	npc NPC
}

func (r *Room) draw() {
	rl.DrawRectangle(0, 200, 500, 300, rl.Green)
}
