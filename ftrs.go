package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	//TextFile text file
	TextFile uint8 = 1 << iota
	//JSONFile JSON file
	JSONFile
)

//TransformObject object to be transformed
type TransformObject struct {
	filePath string
}

//TransformFile transform file
func TransformFile(file string, from string, to string) error {

	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(from), []byte(to), -1)

	if err = ioutil.WriteFile("test_new.txt", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}

func main() {
	err := TransformFile("testfile.txt", "${NEWVAL}", "whatever")
	if err != nil {
		panic(err)
	}
}
