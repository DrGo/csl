// Package tester processes test fixtures for the CSL test suite https://github.com/citation-style-language/test-suite
package tester

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	startBlockLeft   = ">>====="
	startBlockRight  = "=====>>"
	endBlock         = "<<====="
	tagLength        = 7 // len(">>=====")
	modeCitation     = 0
	modeBibliography = 1
)

type Fixture struct {
	Mode          int
	Result        string
	Csl           string
	Input         string
	Bibentries    string
	Bibsection    string
	CitationItems string
	Citations     string
}

// ParseFile parses
func ParseFile(path string) (*Fixture, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data Fixture
	block := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.TrimSpace(line)
		fmt.Println(block, "=>", s)
		switch {
		case s == "":
			continue
		case strings.HasPrefix(s, startBlockLeft): // block starts eg >>===== MODE =====>>
			block = strings.TrimSpace(strings.TrimSuffix(s[tagLength:], startBlockRight))
			fmt.Println("=>", block)
			continue
		case strings.HasPrefix(s, endBlock):
			block = ""
			continue
		case strings.HasPrefix(s, "#"):
			continue
		}
		// within a block
		switch block {
		case "MODE":
			if s == "bibliography" {
				data.Mode = modeBibliography // otherwise defaults to citation
			}
			fmt.Println("MODE:", s, data.Mode)
		case "RESULT":
			data.Result += line + "\n"
		case "CSL":
			data.Csl += line + "\n"
		case "INPUT":
			data.Input += line + "\n"
		// optional sections
		case "BIBENTRIES":
			data.Bibentries += line + "\n"
		case "BIBSECTION":
			data.Bibsection += line + "\n"
		case "CITATION-ITEMS":
			data.CitationItems += line + "\n"
		case "CITATIONS":
			data.Citations += line + "\n"
		default:
			//error
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &data, nil
}
