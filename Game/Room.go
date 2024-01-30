package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Room struct {
	gag GAG
	npc NPC

	floor rl.Rectangle
}

func (r *Room) draw() {
	rl.DrawRectangleRec(r.floor, rl.Green)
}

func (r *Room) update() {

}

func (r *Room) load() {

}

func (r *Room) unload() {

}
