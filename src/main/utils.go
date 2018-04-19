package main

import(
	"math/rand"
	"io/ioutil"
	"strconv"
)

// Map of the printable ASCII range
const charmap = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

/**
  * @Type: Function
  * @Name: isNumeric
  * @Purpose: Checks if the string can be parsed numerically
  **/
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

/**
  * @Type: Function
  * @Name: readFileContents
  * @Purpose: Returns file contents as a string
  **/
func readFileContents(path string) (string, error) {
	fContents, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(fContents), nil
}

/**
  * @Type: Function
  * @Name: writeFileContents
  * @Purpose: Writes contents to a file
  **/
func writeFileContents(path string, contents string) error {
	err := ioutil.WriteFile(path, []byte(contents), 0644)
	return err
}

/**
  * @Type: Function
  * @Name: genRandomString
  * @Purpose: Generates a string of n length with random chars from the charmap array
  **/
func genRandomString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = charmap[rand.Intn(len(charmap))]
	}

	return string(b)
}