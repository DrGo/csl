package csl

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

func ParseFile(path string) (*Style, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var data Style
	err = xml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
