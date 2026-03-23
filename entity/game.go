package entity

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	Players     []Player
	WinnerIDs   uint
}

type Player struct {
	ID     uint
	UserID uint
	GameID uint
	Score  uint
	Answers[]
}
