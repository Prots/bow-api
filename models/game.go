package models

//UserScore use it for recording user's score for each particular frame
type UserScore struct {
	FrameID  int    `json:"frameID"  valid:"required"`
	UserName string `json:"userName" valid:"required"`
	Score    int    `json:"score"    valid:"required"`
}

//GamePersistence interface with methods for Game persistence
type GamePersistence interface {
	SaveUserScore(UserScore) error
	Display() (map[int]map[string]int, error)
}

//GameRepoMap primitive implementation of the GamePersistence based on in memory map
type GameRepoMap struct {
	frames map[int]map[string]int
}

var (
	//GamePersistenceInstance instance of the GameRepoMap
	GamePersistenceInstance GamePersistence = GameRepoMap{
		frames: map[int]map[string]int{},
	}
)

//SaveUserScore save user's score for each frame
func (repo GameRepoMap) SaveUserScore(score UserScore) (err error) {

	if frameMap, present := repo.frames[score.FrameID]; present {
		frameMap[score.UserName] = score.Score
	} else {
		frameMap := make(map[string]int)
		frameMap[score.UserName] = score.Score
		repo.frames[score.FrameID] = frameMap
	}
	return err
}

//Display display user's scores for every frame
func (repo GameRepoMap) Display() (map[int]map[string]int, error) {
	return repo.frames, nil
}
