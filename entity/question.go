package entity

type Question struct {
	ID              uint
	Question        string
	PossibleAnswers []PossibleAnswer
	CorrectAnswer   string
	Difficulty      string
	CategoryID      uint
}

type PossibleAnswer struct {
	ID      uint
	Content string
	Choice  uint8
}
type Answer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
}
