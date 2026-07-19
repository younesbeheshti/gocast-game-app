package entity

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	Players     []Player
}

type Player struct {
	ID     uint
	UserID uint
	GameID uint
	Score  uint
	Answer []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
	Choice     PossibleAnswerChoice
}
