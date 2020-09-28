package pom

import "testing"

func TestGetFirstTwoPartsOfGroupId(t *testing.T) {
	firstTwoParts, err := GetFirstTwoPartsOfGroupId("com.example.application")
	if err != nil {
		t.Errorf("%v", err)
	}

	if firstTwoParts != "com.example" {
		t.Errorf("The first two parts of com.example.application is not com.example")
	}

	_, err = GetFirstTwoPartsOfGroupId("com")
	if err == nil {
		t.Errorf("com got accepted as a at-least-two part group id")
	}
}

