package game

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ANIMATION_TIME float32 = 1.0 / 12.0
var W_WIDTH = 0
var W_HEIGHT = 0
var SHOULD_CLOSE = false

type GameEngine struct {
	sm *SceneManager
}

var G_FLAGS []uint32 = []uint32{
	rl.FlagMsaa4xHint,
	rl.FlagWindowHighdpi,
}

var g_lock = &sync.Mutex{}
var g_instance *GameEngine

func Get_Game() *GameEngine {
	if g_instance == nil {
		g_lock.Lock()
		defer g_lock.Unlock()

		if g_instance == nil {
			g_instance = &GameEngine{
				sm: Get_SceneManager(),
			}
		}
	}

	return g_instance
}

func (g *GameEngine) Run() {
	g.init()

	for !SHOULD_CLOSE {
		g.game_loop()
	}

	g.close()
}

func (g *GameEngine) init() {

	for _, f := range G_FLAGS {
		rl.SetConfigFlags(f)
	}

	rl.InitWindow(20, 20, "Title")
	rl.InitAudioDevice()

	W_WIDTH = rl.GetMonitorWidth(0)
	W_HEIGHT = rl.GetMonitorHeight(0)

	rl.SetWindowSize(W_WIDTH, W_HEIGHT)
	rl.SetWindowPosition(0, 0)
	rl.ToggleFullscreen()

	rl.SetTargetFPS(30)
	// rl.SetExitKey(0)
}

func (g *GameEngine) game_loop() {
	g.sm.draw()
}

func (g *GameEngine) close() {
	g.sm.Close()

	rl.CloseAudioDevice()
	rl.CloseWindow()
}
