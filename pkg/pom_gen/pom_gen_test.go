package pom_gen

import (
	"errors"
	"fmt"
	"github.com/perottobc/mvn-pom-mutator/pkg/xsd_model"
	"testing"
)

func traverseComplexType(schema *xsd_model.Schema, complexType *xsd_model.ComplexType, pad string) {
	fmt.Println(pad + "   complexType: " + complexType.Name)

	for _, childElement := range complexType.All.Element {
		traverse(schema, childElement, "   "+pad)
	}
}

func traverse(schema *xsd_model.Schema, element xsd_model.Element, pad string) error {
	elementType := element.Type
	if element.ComplexType.Sequence == nil {
		return errors.New("element.ComplexType.Sequence is nil")
	}

	if element.ComplexType.Sequence.Element != nil {
		elementType = element.ComplexType.Sequence.Element.Type
	}
	if element.ComplexType.Sequence.Any != nil {
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

	return nil
}

func TestShouldTraversePomModel(t *testing.T) {
	xsd, err := ReadXsd("../../resources/maven-4.0.0.xsd")
	if nil != err {
		t.Error(err)
	}

	err = traverse(xsd, xsd.Element, "")
	if err != nil {
		t.Fail()
	}
}

func TestShouldWritePomModelGoSrc(t *testing.T) {
	err := WritePomModelGoSource("../../resources/maven-4.0.0.xsd", "pom", "../../target/pom.go")

	if err != nil {
		t.Error(err)
	}
}
