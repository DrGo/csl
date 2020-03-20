package csl

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		path           string
		wantErr        bool
		title          string //tests info.title
		citationPrefix string
	}{
		{"styles/apa-6th-edition.csl", false, "American Psychological Association 6th edition", "("},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			s, err := ParseFile(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if s.Info.Title != tt.title {
				t.Errorf("ParseFile() wanted = %v, got %v", tt.title, s.Info.Title)
			}
			if s.Citation.Layout.Prefix != tt.citationPrefix {
				t.Errorf("ParseFile() wanted = %v, got %v", tt.citationPrefix, s.Citation.Layout.Prefix)
			}
		})
	}
}

//ExampleParseFile tests parser against the APA-6th-Edition CSL
func ExampleParseFile() {
	s, err := ParseFile("styles/apa-6th-edition.csl")
	if err != nil {
		fmt.Printf("ParseFile() error = %v", err)
		return
	}
	fmt.Println(s.Info.Title)
	fmt.Println(s.Info.Contributor[1].Name)
	fmt.Println(s.Citation.Layout.Prefix)
	fmt.Println(s.Macro[0].Name)
	fmt.Println(s.Bibliography.Layout.Group[0].Group[0].Choose[0].If.Group[0].Text[0].Term)
	fmt.Println(s.Locale[0].Terms.Term[0].Single)
	// output:
	// American Psychological Association 6th edition
	// Curtis M. Humphrey
	// (
	// container-contributors-booklike
	// circa
	// ed. & trans.
}

// func ExampleParseGoAPA() {
// 	v := reflect.ValueOf(s)
// 	typeOfS := v.Type()

// 	for i := 0; i< v.NumField(); i++ {
// 			fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
// 	}
// }
