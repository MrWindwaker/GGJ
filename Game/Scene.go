package game

type Scene_Type int32

const (
	MENU     Scene_Type = 0
	GAMEPLAY Scene_Type = 1
	CREDITS  Scene_Type = 3
)

type Scene struct {
	room []Room
}
