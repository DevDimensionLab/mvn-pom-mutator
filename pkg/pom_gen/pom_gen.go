package pom_gen

import (
	"encoding/xml"
	"github.com/perottobc/mvn-pom-mutator/pkg/xsd_model"
	"io/ioutil"
	"os"
	"strings"
)

func ReadXsd(path string) (*xsd_model.Schema, error) {
	xsdFile, err := os.Open(path)
	if nil != err {
		return nil, err
	}
	defer xsdFile.Close()

	data, err := ioutil.ReadAll(xsdFile)
	if err != nil {
		return nil, err
	}

	model := xsd_model.Schema{}
	err = xml.Unmarshal(data, &model)
	if nil != err {
		return nil, err
	}

	return &model, nil
}

func WritePomModelGoSource(xsdPath string, packageName string, goPath string) error {
	xsd, err := ReadXsd(xsdPath)
	if nil != err {
		return err
	}

	var structs []Struct
	for _, complexType := range xsd.ComplexType {
		structs = append(structs, createStructFromComplexType(complexType)...)
	}

	structs = append(structs, addStructForAnyElement())
	structs = filterUnique(structs)
	goSource := structsToSource(packageName, structs)

	return ioutil.WriteFile(goPath, goSource, 0644)
}

func filterUnique(structs []Struct) []Struct {
	var unique []Struct
	for _, s := range structs {
		duplicate := false
		for _, u := range unique {
			if u.Name == s.Name {
				duplicate = true
			}
		}
		if !duplicate {
			unique = append(unique, s)
		}
	}
	return unique
}

func addStructForAnyElement() Struct {
	return Struct{
		Name: "Any",
		Fields: []Field{
			{
				Name:       "XMLName",
				Type:       "xml.Name",
				XmlMapping: "",
			},
			{
				Name:       "Value",
				Type:       "string",
				XmlMapping: "`xml:\",chardata\"`",
			},
			{
				Name:       "AnyElements",
				Type:       "[]Any",
				XmlMapping: "`xml:\",any\"`",
			},
		},
	}
}

func structsToSource(packageName string, structs []Struct) []byte {
	var lines []string
	lines = append(lines, "package "+packageName+"\n")
	lines = append(lines, "import \"encoding/xml\"\n")

	for _, structSrc := range structs {
		lines = append(lines, "type "+structSrc.Name+" struct {")

		for _, field := range structSrc.Fields {
			lines = append(lines, "	"+field.Name+" "+field.Type+" "+field.XmlMapping)
		}

		lines = append(lines, "}\n")
	}
	return []byte(strings.Join(lines, "\n"))
}

func createStructFromComplexType(complexType xsd_model.ComplexType) []Struct {
	var structs []Struct
	var fields []Field

	fields = append(fields, Field{
		Name:       "Comment",
		Type:       "string",
		XmlMapping: "`xml:\",comment\"`",
	})

	if "Model" == complexType.Name {
		fields = createProjectModelStandardFields()
	}

	for _, element := range complexType.All.Element {
		field := createFieldFromElement(element)
		fields = append(fields, field)
	}

	return append(structs, Struct{
		Name:   complexType.Name,
		Fields: fields,
	})
}

func createFieldFromElement(element xsd_model.Element) Field {
	t := element.Type
	sequence := element.ComplexType.Sequence
	if nil != sequence {
		if nil != sequence.Element {
			name := strings.Title(element.Name)
			if !(t == "string" || t == "bool") {
				return createFieldFromInlineElement(sequence.Element)
			}
			t = name
		} else if nil != sequence.Any {
			t = "Any"
			//if "unbounded" == sequence.Any.MaxOccurs {
			//	t = "[]" + t
			//}
		}

	}
	t = strings.Replace(t, "xs:", "", 1)
	if "boolean" == t {
		t = "bool"
	}

	if !(t == "string" || t == "bool") {
		t = "*" + t
	}

	elementName := element.Name
	if "type" == elementName {
		elementName = elementName + "_"
	}
	return Field{
		Name:       strings.Title(elementName),
		Type:       t,
		XmlMapping: "`xml:\"" + element.Name + ",omitempty\"`",
	}
}

func xsdTypeToGo(xsdType string) string {
	t := strings.Replace(xsdType, "xs:", "", 1)
	if "boolean" == xsdType {
		t = "bool"
	}

	return t
}

func createFieldFromInlineElement(element *xsd_model.InlineElement) Field {

	typeToGo := xsdTypeToGo(element.Type)

	if element.MaxOccurs == "unbounded" {
		typeToGo = "[]" + typeToGo
	}

	return Field{
		Name:       strings.Title(element.Name),
		Type:       typeToGo,
		XmlMapping: "`xml:\"" + element.Name + ",omitempty\"`",
	}
}

func createProjectModelStandardFields() []Field {
	return []Field{
		{
			Name:       "Comment",
			Type:       "string",
			XmlMapping: "`xml:\",comment\"`",
		},
		{
			Name:       "XMLName",
			Type:       "xml.Name",
			XmlMapping: "`xml:\"project\"`",
		},
		{
			Name:       "Xmlns",
			Type:       "string",
			XmlMapping: "`xml:\"xmlns,attr\"`",
		},
		{
			Name:       "SchemaLocation",
			Type:       "string",
			XmlMapping: "`xml:\"xsi,attr\"`",
		},
		{
			Name:       "Xsi",
			Type:       "string",
			XmlMapping: "`xml:\"schemaLocation,attr\"`",
		},
	}
}

type Field struct {
	Name       string
	Type       string
	XmlMapping string
}

type Struct struct {
	Name   string
	Fields []Field
}
