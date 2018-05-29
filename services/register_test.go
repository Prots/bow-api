package services

import (
	"testing"
	"net/http"
	"github.com/Prots/bow-api/models"
	"encoding/json"
	"net/http/httptest"
	"bytes"
)

func TestRegisterHandler(t *testing.T) {

	var tests = []struct{
		testName           string
		reqBody            interface{}
		expectedStatusCode int
		expectedRespBody   string
	} {
		{
			testName: "Positive registration",
			reqBody:models.Player{Name: "player 1"},
			expectedStatusCode: 201,
			expectedRespBody: `{"name":"player 1"}`,
		},
		{
			testName: "Already exists",
			reqBody:models.Player{Name: "player 1"},
			expectedStatusCode: 400,
			expectedRespBody: `{"description":"Player with name player 1 already exists"}`,
		},
		{
			testName: "Not valid request body",
			reqBody:struct{
				badField string
			}{
				badField:"user 1",
			},
			expectedStatusCode: 400,
			expectedRespBody: `{"description":"name: non zero value required"}`,
		},
	}

	for _, test := range tests {

		playerBin, err := json.Marshal(test.reqBody)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/register",bytes.NewReader(playerBin))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(RegisterHandler)
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != test.expectedStatusCode {
			t.Errorf("Test %s, handler returned wrong status code: got %v want %v",
				test.testName, status, test.expectedStatusCode)
		}

		// Check the response body is what we expect.
		if rr.Body.String() != test.expectedRespBody {
			t.Errorf("Test %s, handler returned unexpected body: got %v want %v",
				test.testName, rr.Body.String(), test.expectedRespBody)
		}
	}


}
