package services

import (
	"github.com/Prots/bow-api/models"
	"net/http"
)

//RecordHandler - record player score for each frame
func RecordHandler(w http.ResponseWriter, r *http.Request) {
	var userScore models.UserScore

	if err := extractStruct(r, &userScore); err != nil {
		renderJSON(w, http.StatusBadRequest, models.Error{Description: err.Error()})
		return
	}

	if !models.PlayerPersistenceInstance.IsUserRegistered(userScore.UserName) {
		renderJSON(w, http.StatusUnauthorized, models.Error{Description: "User is not registered"})
		return
	}

	if err := models.GamePersistenceInstance.SaveUserScore(userScore); err != nil {
		renderJSON(w, http.StatusInternalServerError, models.Error{Description: err.Error()})
		return
	}

	renderJSON(w, http.StatusCreated, userScore)
}

//DisplayHandler display user's scores for each frame
func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	res, err := models.GamePersistenceInstance.Display()

	if err != nil {
		renderJSON(w, http.StatusInternalServerError, models.Error{Description: err.Error()})
		return
	}

	renderJSON(w, http.StatusOK, res)
}
