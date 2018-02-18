// Use the results from parsing the LoC field list pages to create
// the code maps and functions for parsing/translating MARC controlfields

// 'tis ugly code that follows

package main

import (
	"fmt"
	"strings"
	//
	codegen "github.com/gsiems/go-marc21/codegen/pkg"
)

var htmlfiles = map[string]string{
	"Authority":      "input/ecadlist.html",
	"Bibliography":   "input/ecbdlist.html",
	"Classification": "input/eccdlist.html",
	"Community":      "input/eccilist.html",
	"Holdings":       "input/echdlist.html",
}

func main() {

	fmt.Println("package details")
	fmt.Println()
	fmt.Println("// Auto-generated code. Do not edit.")

	// Ensure that the order does not change from run to run
	fl := []string{
		"Authority",
		"Bibliography",
		"Classification",
		"Community",
		"Holdings",
	}

	for _, format := range fl {
		file := htmlfiles[format]
		formatBanner(format)
		ldr := codegen.ExtractLdrStruct(file)

		ve := validElements(ldr.Elements)

		for _, ldrelement := range ve {
			if len(ldrelement.LookupValues) > 0 && ldrelement.FnType != "read" {
				varname := strings.ToLower(format) + "Ldr" + ldrelement.CamelName
				makeLookupList(ldrelement, varname)
			}
		}
		makeFuncs(format, ldr)
	}
}

func validElements(cfe []*codegen.LdrElement) (f []*codegen.LdrElement) {
	for _, c := range cfe {
		if !strings.Contains(c.Name, "OBSOLETE") {
			f = append(f, c)
		}
	}
	return f
}

func formatBanner(format string) {
	fmt.Println()
	bannerLine()
	fmt.Printf("// %s\n", format)
}

func bannerLine() {
	fmt.Println("////////////////////////////////////////////////////////////////////////")
}

func makeLookupList(ldre *codegen.LdrElement, varname string) {
	fmt.Printf("var %s = map[string]string{\n", varname)

	for _, lv := range ldre.LookupValues {
		if strings.Contains(lv.Label, "OBSOLETE") {
			continue
		}

		if lv.Code == "#" {
			fmt.Printf("\t%q: %q,\n", " ", lv.Label)
		} else if lv.Code == "##" {
			fmt.Printf("\t%q: %q,\n", "  ", lv.Label)
		} else if lv.Code == "###" {
			fmt.Printf("\t%q: %q,\n", "   ", lv.Label)
		} else {
			fmt.Printf("\t%q: %q,\n", lv.Code, lv.Label)
		}
	}
	fmt.Println("}")
}

func makeFuncs(format string, ldr codegen.Ldr) {

	funcName := strings.Join([]string{"parse", format, "Ldr"}, "")

	type fn func(e *codegen.LdrElement, fieldName, varname string)

	m := map[string]fn{
		"lookup": makeLdrLookupFunc,
		"read":   makeLdrReadFunc,
	}

	fmt.Println()
	fmt.Printf("// %s parses leader data for %s records data\n", funcName, format)
	fmt.Printf("func %s(s string) (ldr LdrDesc) {\n\n", funcName)
	fmt.Println("\tldr = make(LdrDesc)")
	fmt.Println()
	fmt.Println("\tvar c string")
	fmt.Println("\tvar l string")
	ve := validElements(ldr.Elements)
	for _, e := range ve {
		varname := strings.ToLower(format) + "Ldr" + e.CamelName
		fieldName := fmt.Sprintf("(%02d/%02d) %s", e.Offset, e.Width, e.Name)

		if e.CamelName == "EntryMap" {
			fmt.Printf("\t// %s\n", fieldName)
			continue
		}

		fcn, ok := m[e.FnType]
		if ok {
			fcn(e, fieldName, varname)
		}
	}
	fmt.Println()
	fmt.Println("\treturn ldr")
	fmt.Println("}")
	fmt.Println()
}

func makeLdrLookupFunc(e *codegen.LdrElement, fieldName, varname string) {
	if len(e.LookupValues) > 0 {
		fmt.Printf("\tc, l = codeLookup(%s, s, %d, %d)\n", varname, e.Offset, e.Width)
		fmt.Printf("\tldr[%q] = CodeValue{Code: c, Label: l, Offset: %d, Width: %d}\n",
			fieldName, e.Offset, e.Width)
	}
}

func makeLdrReadFunc(e *codegen.LdrElement, fieldName, varname string) {
	fmt.Printf("\tldr[%q] = CodeValue{Code: pluckBytes(s, %d, %d), Label: \"\", Offset: %d, Width: %d}\n",
		fieldName, e.Offset, e.Width, e.Offset, e.Width)
}
