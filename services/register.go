package services

import (
	"github.com/Prots/bow-api/models"
	"net/http"
)

//RegisterHandler Register players for game. Unregistered players can't play a game.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	if err := extractStruct(r, &player); err != nil {
		renderJSON(w, http.StatusBadRequest, models.Error{Description: err.Error()})
		return
	}

	if err := validate(&player); err != nil {
		renderJSON(w, http.StatusBadRequest, models.Error{Description: err.Error()})
		return
	}

	if err := models.PlayerPersistenceInstance.Save(player); err != nil {
		renderJSON(w, http.StatusBadRequest, models.Error{Description: err.Error()})
		return
	}

	renderJSON(w, http.StatusCreated, player)
}


