package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	flatten "github.com/cameronnewman/go-flatten"
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

	d := json.NewDecoder(bytes.NewReader(filebytes))
	d.UseNumber()
	if err := d.Decode(&jsonBlob); err != nil {
		log.Fatal(err)
	}

	flattenedObject := flatten.Flatten(jsonBlob)
	if *out != "" {
		file, err := os.Create(*out)
		defer check(file.Close)
		if err != nil {
			log.Fatalln("os.Create:", err)
			return
		}
		w := bufio.NewWriter(file)
		defer check(w.Flush)

		for k, v := range flattenedObject {
			fmt.Fprintf(w, "%s %s\n", k, v)
		}
	} else {
		for k, v := range flattenedObject {
			fmt.Println(k, v)
		}
	}
}

func check(f func() error) {
	if err := f(); err != nil {
		log.Println("Received error:", err)
	}
}
