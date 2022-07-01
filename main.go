package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/elimity-com/abnf"
)

func main() {
	rawABNF, err := ioutil.ReadFile("./license.abnf")
	if err != nil {
		fmt.Printf("got error %v", err)
		return
	}
	generator := abnf.ParserGenerator{
		RawABNF: rawABNF,
	}
	functions := generator.GenerateABNFAsOperators()
	for k, v := range functions {
		fmt.Printf("function %s => %v\n", k, v)
	}

	fmt.Printf("found %d rules.\n", len(functions))
	licenseOperator := functions["simple-exp"]

	if licenseOperator == nil {
		fmt.Println("!!! WRONG ABNF maybe, could not load license-exp")
		return
	}

	testdata := []string{
		// "DocumentRef-spdx-tool-1.2:LicenseRef-MIT-Style-2",
		"MIT",
		"MIT-notvalid",
		"LicenseRef-23",
		"LicenseRef-MIT-Style-1",
		"invalid",
		"totally-invalid",
	}

	for _, v := range testdata {
		start := time.Now()
		best := licenseOperator([]byte(v)).Best()
		elapsed := time.Since(start)

		if !best.IsEmpty() && string(best.Value[:]) == v {
			fmt.Printf("%s => matches found %v. Duration = %d ms\n", v, best, elapsed.Milliseconds())
		} else {
			fmt.Printf("%s => no matches found. Duration = %d ms\n", v, elapsed.Milliseconds())
		}
	}
}
