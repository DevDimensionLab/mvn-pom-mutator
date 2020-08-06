package xsd_model

import "encoding/xml"

type Element struct {
	Text       string `xml:",chardata"`
	MinOccurs  string `xml:"minOccurs,attr"`
	Name       string `xml:"name,attr"`
	Type       string `xml:"type,attr"`
	Default    string `xml:"default,attr"`
	Annotation struct {
		Text          string `xml:",chardata"`
		Documentation []struct {
			Text   string `xml:",chardata"`
			Source string `xml:"source,attr"`
		} `xml:"documentation"`
	} `xml:"annotation"`
	ComplexType InlineComplexType `xml:"complexType"`
}

type InlineElement struct {
	Text      string `xml:",chardata"`
	Name      string `xml:"name,attr"`
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
	Type      string `xml:"type,attr"`
}

type InlineComplexType struct {
	Text     string `xml:",chardata"`
	Sequence *struct {
		Text    string         `xml:",chardata"`
		Element *InlineElement `xml:"element"`
		Any     *struct {
			Text            string `xml:",chardata"`
			MinOccurs       string `xml:"minOccurs,attr"`
			MaxOccurs       string `xml:"maxOccurs,attr"`
			ProcessContents string `xml:"processContents,attr"`
		} `xml:"any"`
	} `xml:"sequence"`
}

type ComplexType struct {
	Text       string `xml:",chardata"`
	Name       string `xml:"name,attr"`
	Annotation struct {
		Text          string `xml:",chardata"`
		Documentation []struct {
			Text   string `xml:",chardata"`
			Source string `xml:"source,attr"`
		} `xml:"documentation"`
	} `xml:"annotation"`
	All struct {
		Text    string    `xml:",chardata"`
		Element []Element `xml:"element"`
	} `xml:"all"`
	Attribute []struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name,attr"`
		Type       string `xml:"type,attr"`
		Use        string `xml:"use,attr"`
		Annotation struct {
			Text          string `xml:",chardata"`
			Documentation []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"documentation"`
		} `xml:"annotation"`
	} `xml:"attribute"`
}

type Schema struct {
	XMLName            xml.Name      `xml:"schema"`
	Text               string        `xml:",chardata"`
	Xs                 string        `xml:"xs,attr"`
	ElementFormDefault string        `xml:"elementFormDefault,attr"`
	Xmlns              string        `xml:"xmlns,attr"`
	TargetNamespace    string        `xml:"targetNamespace,attr"`
	Element            Element       `xml:"element"`
	ComplexType        []ComplexType `xml:"complexType"`
}

func (s Schema) GetType(name string) *ComplexType {
	for _, comlexType := range s.ComplexType {
		if comlexType.Name == name {
			return &comlexType
		}
	}
	return nil
}
