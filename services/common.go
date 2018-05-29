package services

import (
	"encoding/json"
	"fmt"
	"github.com/Prots/bow-api/models"
	"io/ioutil"
	"net/http"
	"github.com/asaskevich/govalidator"
)

// extractStruct Unmarshal Request body to inputStructPtr struct
func extractStruct(r *http.Request, inputStructPtr interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("can't read input data: %v, err: %v", string(b), err)
	}

	// decode input data
	if err = json.Unmarshal(b, inputStructPtr); err != nil {
		return fmt.Errorf("can't decode input data: %v, err: %v", string(b), err)
	}

	return nil
}

func validate(input interface{}) error {
	_, err := govalidator.ValidateStruct(input)
	return err
}

// renderJSON is used for rendering JSON response body with appropriate headers
func renderJSON(w http.ResponseWriter, code int, response interface{}) {
	if fmt.Sprint(response) == "[]" {
		response = make([]interface{}, 0)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errData, _ := json.Marshal(models.Error{Description: err.Error()})
		w.Write(errData)
	}

	w.WriteHeader(code)
	w.Write(data)
}
