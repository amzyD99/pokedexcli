package main

import "testing"

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{
		input: "  hello world  ",
		expected: []string{"hello", "world"}
		},
		{
		input: "Charmander Bulbasaur PIKACHU",
		expected: []string{"charmander", "bulbasaur", "pikachu"}
		},
		{
		input: "THIS IS a sentance with few WORDS ThAt are oF DiFFerent cases"
		expected: []string{"this", "is", "a", "sentance", "with", "few", "words", "that", "are", "of", "different", "cases"}
		},
}

	for _,c := cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
		}
	}
}