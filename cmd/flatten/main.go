package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/17twenty/flatter"
)

var (
	src = flag.String("s", "demo.json", "source JSON")
	out = flag.String("w", "", "Specifies an output file if present")
)

func main() {
	flag.Parse()

	filebytes, err := ioutil.ReadFile(*src)
	if err != nil {
		log.Fatalln("ioutil.ReadFile:", err)
		return
	}

	var jsonBlob map[string]interface{}
	json.Unmarshal(filebytes, &jsonBlob)

	flattenedObject := flatten.Flatten(jsonBlob)
	if *out != "" {
		file, err := os.Create(*out)
		defer file.Close()
		if err != nil {
			log.Fatalln("os.Create:", err)
			return
		}
		w := bufio.NewWriter(file)
		defer w.Flush()

		for k, v := range flattenedObject {
			fmt.Fprintf(w, "%s %s\n", k, v)
		}
	} else {
		for k, v := range flattenedObject {
			fmt.Println(k, v)
		}
	}
}
