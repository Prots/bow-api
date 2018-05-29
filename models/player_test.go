package models

import "testing"

func TestPlayerRepoMapSave(t *testing.T) {
	var testCases = []struct {
		testName string
		userName string
		err      string
	}{
		{
			testName: "Positive case",
			userName: "user 1",
			err:      "",
		},
		{
			testName: "Duplicate user name",
			userName: "user 1",
			err:      "Player with name user 1 already exists",
		},
	}

	for _, test := range testCases {
		actualRes := PlayerPersistenceInstance.Save(Player{Name: test.userName})
		if actualRes != nil && actualRes.Error() != test.err {
			t.Fatalf("Test case %s failed, expected %s but got %s",
				test.testName, test.err, actualRes.Error())
		}
	}
}

func TestPlayerRepoMapIsUserRegistered(t *testing.T) {
	var testCases = []struct {
		testName string
		userName string
		result   bool
	}{
		{
			testName: "Registered user",
			userName: "user 1",
			result:   true,
		},
		{
			testName: "User is not registered",
			userName: "user 2",
			result:   false,
		},
	}

	PlayerPersistenceInstance.Save(Player{Name: testCases[0].userName})

	for _, test := range testCases {
		actualRes := PlayerPersistenceInstance.IsUserRegistered(test.userName)
		if actualRes != test.result {
			t.Fatalf("Test case %s failed, expected %v but got %v",
				test.testName, test.result, actualRes)
		}
	}
}
