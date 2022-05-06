package pom_gen

import (
	"encoding/xml"
	"fmt"
	"github.com/devdimensionlab/mvn-pom-mutator/pkg/xsd_model"
	"io/ioutil"
	"os"
	"sort"
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
		newStructs := createStructFromComplexType(complexType, structs)
		structs = append(structs, newStructs...)
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
				Name:       "Attrs",
				Type:       "[]xml.Attr",
				XmlMapping: "`xml:\",any,attr\"`",
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

func createStructFromComplexType(complexType xsd_model.ComplexType, alreadyDefinedStructs []Struct) []Struct {
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
		field, downstreamStructs := createStructFieldFromElement(complexType, element, alreadyDefinedStructs)

		fields = append(fields, field)
		structs = append(structs, downstreamStructs...)
	}

	return append(structs, Struct{
		Name:   complexType.Name,
		Fields: fields,
	})
}

func createStructFieldFromElement(parent xsd_model.ComplexType, element xsd_model.Element, alreadyDefinedStructs []Struct) (Field, []Struct) {
	var structs []Struct
	t := element.Type
	sequence := element.ComplexType.Sequence
	if nil != sequence {
		if nil != sequence.Element {
			var typeName = strings.Title(element.Name)
			if !(t == "string" || t == "bool") {
				newStruct := createStructFromInlineElement(typeName, sequence.Element)
				if hasDuplicate(newStruct, alreadyDefinedStructs) {
					println("FOUND DUPLICATE!" + typeName)
					typeName = fmt.Sprintf("%s%s", strings.Title(parent.Name), strings.Title(element.Name))
					newStruct.Name = fmt.Sprintf("%s%s", strings.Title(parent.Name), strings.Title(element.Name))
				}
				structs = append(structs, newStruct)
			}
			t = typeName
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
	}, structs
}

func hasDuplicate(newStruct Struct, structs []Struct) bool {
	for _, a := range structs {
		if a.Name == newStruct.Name && !hasSameFields(newStruct.Fields, a.Fields) {
			return true
		}
	}
	return false
}

func hasSameFields(aFields []Field, bFields []Field) bool {

	if len(aFields) != len(bFields) {
		return false
	}

	sort.Sort(FieldSort(aFields))
	sort.Sort(FieldSort(bFields))

	for i := range aFields {
		if aFields[i].Name != bFields[i].Name ||
			aFields[i].Type != bFields[i].Type ||
			aFields[i].XmlMapping != bFields[i].XmlMapping {
			return false
		}
	}

	return true
}

func xsdTypeToGo(xsdType string) string {
	t := strings.Replace(xsdType, "xs:", "", 1)
	if "boolean" == xsdType {
		t = "bool"
	}

	return t
}

func createStructFromInlineElement(parentElementName string, element *xsd_model.InlineElement) Struct {

	typeToGo := xsdTypeToGo(element.Type)

	if element.MaxOccurs == "unbounded" {
		typeToGo = "[]" + typeToGo
	}

	comment := Field{
		Name:       "Comment",
		Type:       "string",
		XmlMapping: "`xml:\",comment\"`",
	}

	field := Field{
		Name:       strings.Title(element.Name),
		Type:       typeToGo,
		XmlMapping: "`xml:\"" + element.Name + ",omitempty\"`",
	}

	return Struct{
		Name:   strings.Title(parentElementName),
		Fields: []Field{comment, field},
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

type FieldSort []Field

func (a FieldSort) Len() int      { return len(a) }
func (a FieldSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a FieldSort) Less(i, j int) bool {
	return a[i].Name > a[j].Name
}

type Struct struct {
	Name   string
	Fields []Field
}
