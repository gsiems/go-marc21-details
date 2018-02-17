package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	//
	"github.com/gsiems/go-marc21-details/pkg/details"
	"github.com/gsiems/go-marc21/pkg/marc21"
)

func main() {

	var marcfile, cn string
	if len(os.Args) > 2 {
		marcfile = os.Args[1]
		cn = os.Args[2]
	}

	if marcfile == "" || cn == "" {
		showHelp()
	}

	fi, err := os.Open(marcfile)
	if err != nil {
		log.Fatal(fmt.Printf("File open failed: %q", err))
	}
	defer func() {
		if cerr := fi.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for {
		rec, err := marc21.ParseNextRecord(fi)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if rec.GetControlfield("001") == cn {

			//fmt.Println(rec)
			ldr := details.ParseLeader(*rec)
			dumpLeader(ldr)

			cfs := rec.GetControlfields("001,003,004,005")
			for _, v := range cfs {
				fmt.Printf("%s:    %s\n", v.Tag, v.Text)
			}

			p6 := details.Parse006(*rec)
			dumpCf006(p6)

			p7 := details.Parse007(*rec)
			dumpCf007(p7)

			p8 := details.Parse008(*rec)
			dumpCf008(p8)

			break
		}

	}
}

func dumpLeader(ldr details.LdrDesc) {

	fmt.Println("LDR:")

	// Sort things before printing...
	var keys []string
	for k := range ldr {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := ldr[k]
		dumpCV(v, k, 0)
	}
}

func dumpCf006(p6 details.Cf008Desc) {

	if len(p6) > 0 {
		fmt.Println("006:")
		// Sort things before printing...
		var keys6 []string
		for k := range p6 {
			keys6 = append(keys6, k)
		}
		sort.Strings(keys6)

		for _, k := range keys6 {
			l := p6[k]

			for i, v := range l {
				dumpCV(v, k, i)
			}
		}
	}
}

func dumpCf008(p8 details.Cf008Desc) {

	fmt.Println("008:")

	// Sort things before printing...
	var keys []string
	for k := range p8 {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		l := p8[k]

		for i, v := range l {
			dumpCV(v, k, i)
		}
	}
}

func dumpCf007(p7 []details.Cf007Desc) {

	for _, cf := range p7 {
		fmt.Println("007:")

		// Sort things before printing...
		var keys []string
		for k := range cf {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := cf[k]
			dumpCV(v, k, 0)
		}
	}
}

func dumpCV(v details.CodeValue, k string, i int) {

	v.Code = strings.Replace(v.Code, " ", "#", -1)
	if v.Code == "" {
		v.Code = " "
	}
	n := strings.SplitN(k, " ", 2)

	if i == 0 {
		if v.Width == 1 {
			fmt.Printf("  %02d -     %s: ( %s = %q )\n", v.Offset, v.Code, n[1], v.Label)
		} else {
			end := v.Offset + v.Width - 1
			fmt.Printf("  %02d-%02d -  %s: ( %s = %q )\n", v.Offset, end, v.Code, n[1], v.Label)
		}
	} else if v.Code != " " && v.Label != "" {
		fmt.Printf("           %s: ( %s = %q )\n", v.Code, n[1], v.Label)
	}
}

func showHelp() {
	fmt.Println(os.Args[0])
	fmt.Println("   Extract the specifed record from a MARC file and print the detailed results.")
	fmt.Printf("    Usage: %s <MARC file to sesrch> <control number for the record to parse>\n", os.Args[0])
	fmt.Println()
	os.Exit(0)
}

/*
// Implement the Stringer interface for "Pretty-printing"
func (ldr Leader) String() string {

	recFormat := ldr.RecordFormat()

	ret := fmt.Sprintf("MARC 21 %s\n", marcFormatName[recFormat])
	ret += fmt.Sprintf("LDR %s\n", ldr.Text)
	ret += fmt.Sprintf("  00-04 -  %s: ( Record length )\n", pluckBytes(ldr.Text, 0, 5))

	switch recFormat {
	case Bibliography:

		code, label := ldrlk(bibliographyFieldData["RecordStatus"], ldr.Text, 5)
		ret += fmt.Sprintf("  05 -     %s: ( Record status = %q )\n", code, label)

		code, label = ldrlk(bibliographyFieldData["TypeOfRecord"], ldr.Text, 6)
		ret += fmt.Sprintf("  06 -     %s: ( Type of Record = %q )\n", code, label)

		code, label = ldrlk(bibliographyFieldData["BibliographicLevel"], ldr.Text, 7)
		ret += fmt.Sprintf("  07 -     %s: ( Bibliographic level = %q )\n", code, label)

		code, label = ldrlk(bibliographyFieldData["TypeOfControl"], ldr.Text, 8)
		ret += fmt.Sprintf("  08 -     %s: ( Type of control = %q )\n", code, label)

		code, label = ldrlk(bibliographyFieldData["CharacterCodingScheme"], ldr.Text, 9)
		ret += fmt.Sprintf("  09 -     %s: ( Character coding scheme = %q )\n", code, label)

		ret += fmt.Sprintf("  10 -     %s: ( Indicator count )\n", pluckByte(ldr.Text, 10))
		ret += fmt.Sprintf("  11 -     %s: ( Subfield code count )\n", pluckByte(ldr.Text, 11))
		ret += fmt.Sprintf("  12-16 -  %s: ( Base address of data )\n", pluckBytes(ldr.Text, 12, 5))

		code, label = ldrlk(bibliographyFieldData["EncodingLevel"], ldr.Text, 17)
		ret += fmt.Sprintf("  17 -     %s: ( Encoding level = %q )\n", code, label)

		code, label = ldrlk(bibliographyFieldData["DescriptiveCatalogingForm"], ldr.Text, 18)
		ret += fmt.Sprintf("  18 -     %s: ( Descriptive cataloging form = %q )\n", code, label)

		code, label = ldrlk(bibliographyFieldData["MultipartResourceRecordLevel"], ldr.Text, 19)
		ret += fmt.Sprintf("  19 -     %s: ( Multipart resource record level = %q )\n", code, label)

	case Holdings:

		code, label := ldrlk(holdingsFieldData["RecordStatus"], ldr.Text, 5)
		ret += fmt.Sprintf("  05 -     %s: ( Record status = %q )\n", code, label)

		code, label = ldrlk(holdingsFieldData["TypeOfRecord"], ldr.Text, 6)
		ret += fmt.Sprintf("  06 -     %s: ( Type of Record = %q )\n", code, label)

		ret += fmt.Sprintf("  07-08 -  %s: ( Undefined character positions )\n", pluckBytes(ldr.Text, 7, 2))

		code, label = ldrlk(holdingsFieldData["CharacterCodingScheme"], ldr.Text, 9)
		ret += fmt.Sprintf("  09 -     %s: ( Character coding scheme = %q )\n", code, label)

		ret += fmt.Sprintf("  10 -     %s: ( Indicator count )\n", pluckByte(ldr.Text, 10))
		ret += fmt.Sprintf("  11 -     %s: ( Subfield code length )\n", pluckByte(ldr.Text, 11))
		ret += fmt.Sprintf("  12-16 -  %s: ( Base address of data )\n", pluckBytes(ldr.Text, 12, 5))

		code, label = ldrlk(holdingsFieldData["EncodingLevel"], ldr.Text, 17)
		ret += fmt.Sprintf("  17 -     %s: ( Encoding level = %q )\n", code, label)

		code, label = ldrlk(holdingsFieldData["ItemInformationInRecord"], ldr.Text, 18)
		ret += fmt.Sprintf("  18 -     %s: ( Item information in record = %q )\n", code, label)

		ret += fmt.Sprintf("  19 -     %s: ( Undefined character position )\n", pluckByte(ldr.Text, 19))

	case Authority:

		code, label := ldrlk(authorityFieldData["RecordStatus"], ldr.Text, 5)
		ret += fmt.Sprintf("  05 -     %s: ( Record status = %q )\n", code, label)

		code, label = ldrlk(authorityFieldData["TypeOfRecord"], ldr.Text, 6)
		ret += fmt.Sprintf("  06 -     %s: ( Type of Record = %q )\n", code, label)

		ret += fmt.Sprintf("  07-08 -  %s: ( Undefined character positions )\n", pluckBytes(ldr.Text, 7, 2))

		code, label = ldrlk(authorityFieldData["CharacterCodingScheme"], ldr.Text, 9)
		ret += fmt.Sprintf("  09 -     %s: ( Character coding scheme = %q )\n", code, label)

		ret += fmt.Sprintf("  10 -     %s: ( Indicator count )\n", pluckByte(ldr.Text, 10))
		ret += fmt.Sprintf("  11 -     %s: ( Subfield code length )\n", pluckByte(ldr.Text, 11))
		ret += fmt.Sprintf("  12-16 -  %s: ( Base address of data )\n", pluckBytes(ldr.Text, 12, 5))

		code, label = ldrlk(authorityFieldData["EncodingLevel"], ldr.Text, 17)
		ret += fmt.Sprintf("  17 -     %s: ( Encoding level = %q )\n", code, label)

		code, label = ldrlk(authorityFieldData["PunctuationPolicy"], ldr.Text, 18)
		ret += fmt.Sprintf("  18 -     %s: ( Punctuation policy = %q )\n", code, label)

		ret += fmt.Sprintf("  19 -     %s: ( Undefined character position )\n", pluckByte(ldr.Text, 19))

	case Classification:

		code, label := ldrlk(classificationFieldData["RecordStatus"], ldr.Text, 5)
		ret += fmt.Sprintf("  05 -     %s: ( Record status = %q )\n", code, label)

		code, label = ldrlk(classificationFieldData["TypeOfRecord"], ldr.Text, 6)
		ret += fmt.Sprintf("  06 -     %s: ( Type of Record = %q )\n", code, label)

		ret += fmt.Sprintf("  07-08 -  %s: ( Undefined character positions )\n", pluckBytes(ldr.Text, 7, 2))

		code, label = ldrlk(classificationFieldData["CharacterCodingScheme"], ldr.Text, 9)
		ret += fmt.Sprintf("  09 -     %s: ( Character coding scheme = %q )\n", code, label)

		ret += fmt.Sprintf("  10 -     %s: ( Indicator count )\n", pluckByte(ldr.Text, 10))
		ret += fmt.Sprintf("  11 -     %s: ( Subfield code length )\n", pluckByte(ldr.Text, 11))
		ret += fmt.Sprintf("  12-16 -  %s: ( Base address of data )\n", pluckBytes(ldr.Text, 12, 5))

		code, label = ldrlk(classificationFieldData["EncodingLevel"], ldr.Text, 17)
		ret += fmt.Sprintf("  17 -     %s: ( Encoding level = %q )\n", code, label)

		ret += fmt.Sprintf("  18-19 -  %s: ( Undefined character positions )\n", pluckBytes(ldr.Text, 18, 2))

	case Community:
		code, label := ldrlk(communityFieldData["RecordStatus"], ldr.Text, 5)
		ret += fmt.Sprintf("  05 -     %s: ( Record status = %q )\n", code, label)

		code, label = ldrlk(communityFieldData["TypeOfRecord"], ldr.Text, 6)
		ret += fmt.Sprintf("  06 -     %s: ( Type of Record = %q )\n", code, label)

		code, label = ldrlk(communityFieldData["KindOfData"], ldr.Text, 7)
		ret += fmt.Sprintf("  07 -     %s: ( Kind of data = %q )\n", code, label)

		ret += fmt.Sprintf("  08 -     %s: ( Undefined character position )\n", pluckByte(ldr.Text, 8))

		code, label = ldrlk(communityFieldData["CharacterCodingScheme"], ldr.Text, 9)
		ret += fmt.Sprintf("  09 -     %s: ( Character coding scheme = %q )\n", code, label)

		ret += fmt.Sprintf("  10 -     %s: ( Indicator count )\n", pluckByte(ldr.Text, 10))
		ret += fmt.Sprintf("  11 -     %s: ( Subfield code length )\n", pluckByte(ldr.Text, 11))
		ret += fmt.Sprintf("  12-16 -  %s: ( Base address of data )\n", pluckBytes(ldr.Text, 12, 5))

		ret += fmt.Sprintf("  17-19 -  %s: ( Undefined character positions )\n", pluckBytes(ldr.Text, 17, 3))

	}

	ret += fmt.Sprintf("  20 -     %s: ( Length of the length-of-field portion )\n", pluckByte(ldr.Text, 20))
	ret += fmt.Sprintf("  21 -     %s: ( Length of the starting-character-position portion )\n", pluckByte(ldr.Text, 21))
	ret += fmt.Sprintf("  22 -     %s: ( Length of the implementation-defined portion )\n", pluckByte(ldr.Text, 22))
	ret += fmt.Sprintf("  23 -     %s: ( Undefined )\n", pluckByte(ldr.Text, 23))

	return ret
}
*/

/*
// ldrlk is a helper function for the leader stringer interface
func ldrlk(codeList map[string]string, s string, i int) (code, label string) {

	code, label = shortCodeLookup(codeList, s, i)
	if code == " " {
		code = "#"
	}

	return code, label
}
*/
