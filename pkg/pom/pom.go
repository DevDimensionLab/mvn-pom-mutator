package pom

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func GetModelFrom(path string) (*Model, error) {
	xmlFile, err := os.Open(path)
	if nil != err {
		return nil, err
	}

	defer xmlFile.Close()
	data, _ := ioutil.ReadAll(xmlFile)

	model := Model{}
	err = xml.Unmarshal(data, &model)
	if nil != err {
		return nil, err
	}
	return &model, nil
}

func Marshall(project *Model) ([]byte, error) {
	raw, err := xml.MarshalIndent(project, "", "   ")
	if nil != err {
		return nil, err
	}

	rawLines := bytes.Split(raw, []byte("\n"))
	firstLine := searchAndReplace(rawLines[0],
		replace{
			s1: "xsi=",
			s2: "xmlns:xsi=",
		},
		replace{
			s1: "schemaLocation=",
			s2: "xsi:schemaLocation=",
		})

	whitespaceFix := cleanUnwanted(bytes.Join(rawLines[1:], []byte("\n")), "&#xA;", "&#x9;", "&#x20;", "&#xD;")
	namespaceFix := searchAndReplace(whitespaceFix,
		replace{
			s1: " xmlns=\"http://maven.apache.org/POM/4.0.0\"",
			s2: "",
		},
		replace{
			s1: "xsi=",
			s2: "xmlns:xsi=",
		},
		replace{
			s1: "schemaLocation=",
			s2: "xsi:schemaLocation=",
		})

	cleaned := removeTrailingWhitespace(namespaceFix)
	return append(append(firstLine, []byte("\n")...), cleaned...), nil
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

func removeTrailingWhitespace(input []byte) []byte {
	var output []byte
	br := bufio.NewReader(bytes.NewReader(input))
	for {
		line, _, err := br.ReadLine()
		cleanLine := strings.TrimRight(string(line), " ")
		output = append(output, []byte(cleanLine+"\n")...)
		if err == io.EOF {
			break
		}
	}

	return output
}

type replace struct {
	s1 string
	s2 string
}
