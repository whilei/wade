package main

import (
	"bytes"
	"fmt"
)

const (
	CreateElementOpener    = "vdom.NewElement"
	CreateTextNodeOpener   = "vdom.NewTextNode"
	CreateComElementOpener = "vdom.NewComElement"
	AttributeMapOpener     = "vdom.Attributes"
	NodeTypeName           = "vdom.Node"
	NodeListOpener         = "[]vdom.Node"
	RenderFuncOpener       = "func %vRender(stateData interface{}) *vdom.Element "
	RenderEmbedString      = "(this %v) "
	ComponentDataOpener    = "wade.Com"
)

func domElType(elTag string) (string, string) {
	switch elTag {
	case "input":
		return "vdom.DOMInputEl", "wade.VdomDrv().ToInputEl"
	}

	return "vdom.DOMNode", ""
}

func valueToStringCode(vcode string) string {
	return fmt.Sprintf(`wade.Str(%v)`, vcode)
}

func componentSetStateCode(sField, sType string, isPointer bool) string {
	typ := sType
	if isPointer {
		typ = "*" + sType
	}

	return fmt.Sprintf("if stateData != nil { this.%v = stateData.(%v) }", sField, typ)
}

func componentRefsVarCode(comName string) (string, string) {
	return fmt.Sprintf("refs := %vRefs{}", comName),
		"this.Com.InternalRefsHolder = refs"
}

func componentSetRefCode(refName string, varName string, elTag string) string {
	_, elMk := domElType(elTag)
	return fmt.Sprintf("%v.OnRendered = func(dNode vdom.DOMNode) { refs.%v = %v(dNode) }",
		varName, refName, elMk)
}

func prelude(pkgName string, imports []importInfo) string {
	var importCode bytes.Buffer
	if imports != nil {
		for _, imp := range imports {
			importCode.WriteString(fmt.Sprintf(`%v "%v"`+"\n", imp.as, imp.path))
		}
	}
	return `package ` + pkgName + `

// THIS FILE IS AUTOGENERATED BY WADE.GO FUEL
// CHANGES WILL BE OVERWRITTEN
import (
	"fmt"

	"github.com/gowade/wade/vdom"
	"github.com/gowade/wade"
	` + importCode.String() + `
)

func init() {
	_, _, _ = fmt.Printf, vdom.NewElement, wade.Str
}

`
}
