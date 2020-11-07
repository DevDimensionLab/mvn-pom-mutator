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

func TestDependencies_FindDuplicates(t *testing.T) {
	model, err := GetModelFrom("test/pom.xml")
	if err != nil {
		t.Fail()
	}

	if model != nil {
		duplicates := model.Dependencies.FindDuplicates()
		if len(duplicates) != 1 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
