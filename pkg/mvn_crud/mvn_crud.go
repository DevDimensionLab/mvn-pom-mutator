package mvn_crud

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"mvn-cli/pkg/pom"
	"os"
)

func GetPomModel(path string) (*pom.Model, error) {

	xmlFile, err := os.Open(path)
	if nil != err {
		return nil, err
	}

	defer xmlFile.Close()
	data, _ := ioutil.ReadAll(xmlFile)

	model := pom.Model{}
	err = xml.Unmarshal(data, &model)
	if nil != err {
		return nil, err
	}
	return &model, nil
}

func Marshall(project *pom.Model) ([]byte, error) {
	raw, err := xml.MarshalIndent(project, "", "   ")
	if nil != err {
		return nil, err
	}

	whitespaceFix := cleanUnwanted(raw, "&#xA;", "&#x9;")
	namespaceFix := searchAndReplace(whitespaceFix,
		replace{
			s1: " xmlns=\"http://maven.apache.org/POM/4.0.0\">",
			s2: ">",
		}, replace{
			s1: "xsi=",
			s2: "xmlns:xsi=",
		}, replace{
			s1: "schemaLocation=",
			s2: "xsi:schemaLocation=",
		})

	return namespaceFix, nil
}

func searchAndReplace(content []byte, commands ...replace) []byte {
	newContent := bytesReplace(content, commands[0].s1, commands[0].s2)

	if 1 < len(commands) {
		return searchAndReplace(newContent, commands[1:]...)
	}
	return newContent
}

func cleanUnwanted(content []byte, str ...string) []byte {
	newContent := bytesReplace(content, str[0], "")
	if 1 < len(str) {
		return cleanUnwanted(newContent, str[1:]...)
	}
	return newContent
}

func bytesReplace(content []byte, from string, to string) []byte {
	return bytes.ReplaceAll(content, []byte(from), []byte(to))
}

type replace struct {
	s1 string
	s2 string
}
