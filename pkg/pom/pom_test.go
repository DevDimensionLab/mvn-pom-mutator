package pom

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestShouldReadAndWritePomXml(t *testing.T) {

	model, err := GetModelFrom("../../resources/github.com/manouti/java-http-client-api/pom.xml")
	if nil != err {
		t.Error(err)
	}

	fmt.Println(model.ArtifactId)

	xml, _ := Marshall(model)

	_ = ioutil.WriteFile("../../target/pom.xml", xml, 0644)
}
