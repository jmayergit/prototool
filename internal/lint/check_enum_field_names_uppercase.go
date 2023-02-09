// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package lint

import (
	"github.com/emicklei/proto"
	"github.com/jmayergit/prototool/internal/strs"
	"github.com/jmayergit/prototool/internal/text"
)

var enumFieldNamesUppercaseLinter = NewLinter(
	"ENUM_FIELD_NAMES_UPPERCASE",
	"Verifies that all enum field names are UPPERCASE.",
	checkEnumFieldNamesUppercase,
)

func checkEnumFieldNamesUppercase(add func(*text.Failure), dirPath string, descriptors []*FileDescriptor) error {
	return runVisitor(enumFieldNamesUppercaseVisitor{baseAddVisitor: newBaseAddVisitor(add)}, descriptors)
}

type enumFieldNamesUppercaseVisitor struct {
	baseAddVisitor
}

func (v enumFieldNamesUppercaseVisitor) VisitMessage(message *proto.Message) {
	for _, element := range message.Elements {
		element.Accept(v)
	}
}

func (v enumFieldNamesUppercaseVisitor) VisitEnum(enum *proto.Enum) {
	for _, element := range enum.Elements {
		element.Accept(v)
	}
}

func (v enumFieldNamesUppercaseVisitor) VisitEnumField(field *proto.EnumField) {
	if !strs.IsUppercase(field.Name) {
		v.AddFailuref(field.Position, "Field name %q must be uppercase.", field.Name)
	}
}
