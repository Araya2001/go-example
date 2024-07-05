package service

import "fmt"

// SimpleStruct - Every Attribute that is in PascalCase gets a public access modifier,
// the camelCase is a private access modifier
type SimpleStruct struct {
	AttrOne            string
	AttrTwo            string
	AttrThree          string
	combinedAttributes string
}

// SimpleStructSampler - Every Method that is in PascalCase gets a public access modifier,
// the camelCase is a private access modifier
type SimpleStructSampler interface {
	GetOne() (string, error)
	GetTwo() (string, error)
	GetThree() (string, error)
	GetCombinedAttributes() (string, error)
	GetSimpleStruct() (*SimpleStruct, error)
	ModifySimpleStruct(attrOne string, attrTwo string, attrThree string) error
	callPrivateMethod()
}

func (s *SimpleStruct) GetOne() (string, error) {
	return s.AttrOne, nil
}

func (s *SimpleStruct) GetTwo() (string, error) {
	return s.AttrTwo, nil
}

func (s *SimpleStruct) GetThree() (string, error) {
	return s.AttrThree, nil
}

func (s *SimpleStruct) GetCombinedAttributes() (string, error) {
	return s.combinedAttributes, nil
}

func (s *SimpleStruct) GetSimpleStruct() (*SimpleStruct, error) {
	return s, nil
}

func (s *SimpleStruct) ModifySimpleStruct(attrOne string, attrTwo string, attrThree string) error {
	s.combinedAttributes = ""

	// This defer will make the s.callPrivateMethod call before the return, no matter the condition
	defer s.callPrivateMethod()

	if attrOne != "" {
		s.AttrOne = attrOne
		s.combinedAttributes += formatAttribute(attrOne)
	}

	if attrTwo != "" {
		s.AttrTwo = attrTwo
		s.combinedAttributes += formatAttribute(attrTwo)
	}

	if attrThree != "" {
		s.AttrThree = attrThree
		s.combinedAttributes += formatAttribute(attrThree)
	}

	return nil
}

func formatAttribute(attribute string) string {
	return "[" + attribute + "]"
}

func (s *SimpleStruct) callPrivateMethod() {
	fmt.Println("calling private method from within ModifySimpleStruct")
}
