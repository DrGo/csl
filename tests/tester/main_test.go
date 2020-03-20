package tester

import "fmt"

func ExampleParseFile() {
	s, err := ParseFile("../test-suite/processor-tests/humans/magic_SecondFieldAlign.txt")
	if err != nil {
		fmt.Printf("ParseFile() error = %v", err)
		return
	}
	fmt.Println(s.Mode)
	fmt.Print(s.Result)
	// fmt.Print(s.Csl)
	// fmt.Print(s.Input)

	// output:
	// 0
	// Doe

}
