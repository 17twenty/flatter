/*
Package flatten is a simple library to flatten JSON into dot notation. Forked from https://github.com/17twenty/flatter
*/
package flatten

// package main
//
// import (
//	"bufio"
//	"bytes"
//	"encoding/json"
//	"flag"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//
//	flatten "github.com/cameronnewman/go-flatten"
// )
//
// var (
//	src = flag.String("s", "demo.json", "source JSON")
//	out = flag.String("w", "", "Specifies an output file if present")
// )
//
// func main() {
//	flag.Parse()
//
//	filebytes, err := ioutil.ReadFile(*src)
//	if err != nil {
//		log.Fatalln("ioutil.ReadFile:", err)
//		return
//	}
//
//	var jsonBlob map[string]interface{}
//
//	d := json.NewDecoder(bytes.NewReader(filebytes))
//	d.UseNumber()
//	if err := d.Decode(&jsonBlob); err != nil {
//		log.Fatal(err)
//	}
//
//	flattenedObject := flatten.Flatten(jsonBlob)
//	if *out != "" {
//		file, err := os.Create(*out)
//		defer file.Close()
//		if err != nil {
//			log.Fatalln("os.Create:", err)
//			return
//		}
//		w := bufio.NewWriter(file)
//		defer w.Flush()
//
//		for k, v := range flattenedObject {
//			fmt.Fprintf(w, "%s %s\n", k, v)
//		}
//	} else {
//		for k, v := range flattenedObject {
//			fmt.Println(k, v)
//		}
//	}
// }
//
