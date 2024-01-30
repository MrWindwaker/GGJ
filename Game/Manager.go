package game

import (
	"sync"

	rl_gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneManager struct {
	c_state   Scene_Type
	anim_time float32

	pl *Player

	exp   Explotion_GAG
	hood  NPC
	inter Interactable

	loaded bool

	main_theme   rl.Music
	is_mt_loaded bool
}

var sm_lock = &sync.Mutex{}
var sm_insta *SceneManager

func Get_SceneManager() *SceneManager {

	if sm_insta == nil {
		sm_lock.Lock()
		defer sm_lock.Unlock()

		if sm_insta == nil {
			sm_insta = &SceneManager{
				c_state:      MENU,
				loaded:       false,
				pl:           Get_Player(),
				is_mt_loaded: false,
				inter: Interactable{
					active:     false,
					active_key: rl.KeyE,
					width:      128,
					height:     128,
					pos:        rl.NewVector2(1000, 700),
					spent:      false,
				},
			}
		}
	}

	return sm_insta
}

func (sm *SceneManager) draw_menu() {
	if !sm.is_mt_loaded {
		sm.main_theme = rl.LoadMusicStream("Assets/Sounds/Theme.mp3")
		sm.is_mt_loaded = true
		rl.PlayAudioStream(sm.main_theme.Stream)
	}

	if rl_gui.Button(
		rl.NewRectangle(0, 0, 200, 100),
		"Start Game",
	) {
		sm.c_state = GAMEPLAY
	}

	if rl_gui.Button(
		rl.NewRectangle(0, 100, 200, 100),
		"Credits",
	) {
		sm.c_state = CREDITS
	}

	if rl_gui.Button(
		rl.NewRectangle(0, 200, 200, 100),
		"Exit",
	) {
		SHOULD_CLOSE = true
	}
}

func (sm *SceneManager) load_assets() {

	sm.hood.set_values("Assets/Images/NPCs/Fox.png", 4, rl.NewVector2(1000, 700))
	sm.hood.change_direction(-1)

	sm.exp.load()
	sm.loaded = true

	if !sm.pl.is_loaded {
		sm.pl.load(rl.NewVector2(200, 700))
	}

	sm.inter.script = func() {
		sm.inter.spent = true
	}
}

func (sm *SceneManager) draw_gameplay() {

	rl.StopMusicStream(sm.main_theme)

	if !sm.loaded {
		sm.load_assets()
	}

	if sm.pl.is_loaded {
		sm.pl.draw()
	}

	sm.inter.draw()

	sm.hood.draw()
	sm.exp.draw()
}

func (sm *SceneManager) draw_credits() {
	rl.DrawText("MRWW", 200, 200, 20, rl.RayWhite)

	if rl_gui.Button(
		rl.NewRectangle(150, 250, 200, 100),
		"Back",
	) {
		sm.c_state = MENU
	}
}

func (sm *SceneManager) draw() {

	// Update
	SHOULD_CLOSE = rl.WindowShouldClose()

	if sm.is_mt_loaded && (sm.c_state == MENU || sm.c_state == CREDITS) {
		rl.UpdateMusicStream(sm.main_theme)
	}

	sm.anim_time += rl.GetFrameTime()

	if sm.anim_time >= ANIMATION_TIME {
		sm.hood.animate()
		sm.exp.animate()
		sm.anim_time = 0

		if sm.pl.is_loaded {
			sm.pl.animate()
		}
	}

	if rl.IsKeyPressed(rl.KeyE) && sm.hood.active {
		sm.hood.active = false
		sm.exp.play(sm.hood.pos)
	}

	if rl.IsKeyPressed(rl.KeyR) {
		sm.hood.active = true
	}

	if sm.c_state == GAMEPLAY {
		sm.pl.update()
		sm.inter.update(*sm.pl)
	}

	rl.BeginDrawing()

	rl.ClearBackground(rl.GetColor(0x333333FF))

	switch sm.c_state {
	case MENU:
		sm.draw_menu()
	case GAMEPLAY:
		sm.draw_gameplay()
	case CREDITS:
		sm.draw_credits()
	default:
		break
	}

	rl.EndDrawing()
}

func (sm *SceneManager) Close() {

	if sm.is_mt_loaded {
		rl.UnloadMusicStream(sm.main_theme)
	}

	if sm.loaded {
		sm.hood.close()
		sm.exp.unload()
	}

	if sm.pl.is_loaded {
		sm.pl.unload()
	}
}
