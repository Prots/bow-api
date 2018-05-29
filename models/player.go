package models

import "fmt"

//Player represents game player and could have any other additional fields
type Player struct {
	Name string `json:"name" valid:"required"`
}

//PlayerPersistence interface which describes Player persistence methods
type PlayerPersistence interface {
	Save(Player) error
	IsUserRegistered(string) bool
}

//PlayerRepoMap implementation of PlayerPersistence based on in memory map
type PlayerRepoMap struct {
	players map[string]struct{}
}

var (
	//PlayerPersistenceInstance instance of the PlayerRepoMap
	PlayerPersistenceInstance PlayerPersistence = PlayerRepoMap{
		players: make(map[string]struct{}),
	}
)

//Save save user to repo during registration
func (repo PlayerRepoMap) Save(player Player) error {

	if _, ok := repo.players[player.Name]; ok {
		return fmt.Errorf("Player with name %s already exists", player.Name)
	}

	repo.players[player.Name] = struct{}{}
	return nil
}

//IsUserRegistered checks whether user is registered.
func (repo PlayerRepoMap) IsUserRegistered(name string) bool {
	if _, ok := repo.players[name]; !ok {
		return false
	}
	return true
}
