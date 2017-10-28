package util

import (
	"io/ioutil"
)

func LoadInputFile(filename string) (content string, err error) {
	fileContentBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	content = string(fileContentBytes[:])
	return
}
