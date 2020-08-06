package pom_gen

import (
	"fmt"
	"mvn-cli/pkg/xsd_model"
	"testing"
)

func traverseComplexType(schema *xsd_model.Schema, complexType *xsd_model.ComplexType, pad string) {
	fmt.Println(pad + "   complexType: " + complexType.Name)

	for _, childElement := range complexType.All.Element {
		traverse(schema, childElement, "   "+pad)
	}
}

func traverse(schema *xsd_model.Schema, element xsd_model.Element, pad string) {
	elementType := element.Type
	if nil != element.ComplexType.Sequence.Element {
		elementType = element.ComplexType.Sequence.Element.Type
	}
	if nil != element.ComplexType.Sequence.Any {
		elementType = "any"
	}

	fmt.Println(pad + element.Name + " : " + elementType)
	complexType := schema.GetType(elementType)
	if nil != complexType {
		traverseComplexType(schema, complexType, "   "+pad)
	} else {
		if "" == elementType {
			fmt.Println("\"\" == elementType")
		}
	}
}

func TestShouldTraversePomModel(t *testing.T) {
	xsd, err := ReadXsd("../../resources/maven-4.0.0.xsd")
	if nil != err {
		t.Error(err)
	}

	traverse(xsd, xsd.Element, "")
}

func TestShouldWritePomModelGoSrc(t *testing.T) {
	WritePomModelGoSource("../../resources/maven-4.0.0.xsd", "pom", "../../target/pom.go")
}
