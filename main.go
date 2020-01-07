package main

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
	"io/ioutil"
)

func main() {

	// data scraped from: https://www.dmcc.ae/business-search
	writeFile("0-6000.json", "0-6000-pretty.json")
	writeFile("6000-13000.json", "6000-13000-pretty.json")

}

func writeFile(inFilename, outFilename string) {
	b, _ := ioutil.ReadFile(inFilename)

	// mask bytes to string
	str := string(b)

	// get the wanted data from JSON array
	result := gjson.Get(str, "0.result.resultData.v").String()

	// make JSON format more human readable
	result = string(pretty.Pretty([]byte(result)))

	// write file
	ioutil.WriteFile(outFilename, []byte(result), 0)
}
