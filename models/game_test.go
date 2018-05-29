package models

import "testing"

func TestGameRepoMapSaveUserScore(t *testing.T) {
	var testCases = []struct {
		testName  string
		userScore UserScore
		err       string
	}{
		{
			testName: "Positive case",
			userScore: UserScore{
				UserName: "user 3",
				FrameID:  1,
				Score:    10,
			},
			err: "",
		},
		{
			testName: "User is not registered",
			userScore: UserScore{
				UserName: "user 4",
				FrameID:  1,
				Score:    5,
			},
			err: "Player with name user 4 is not registered",
		},
	}

	PlayerPersistenceInstance.Save(Player{Name: testCases[0].userScore.UserName})

	for _, test := range testCases {
		actualRes := GamePersistenceInstance.SaveUserScore(test.userScore)
		if actualRes != nil && actualRes.Error() != test.err {
			t.Fatalf("Test case %s failed, expected %s but got %s",
				test.testName, test.err, actualRes.Error())
		}
	}
}
