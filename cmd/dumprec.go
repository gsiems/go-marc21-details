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
	fmt.Println("   Extract the specified record from a MARC file and print the detailed results.")
	fmt.Printf("    Usage: %s <MARC file to sesrch> <control number for the record to parse>\n", os.Args[0])
	fmt.Println()
	os.Exit(0)
}
