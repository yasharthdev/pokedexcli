package main

import (
	"testing"
	"reflect"
)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		name string
		input string
		want []string
	}{
		{
			name: "extra whitespaces",
			input: "hello    world    ", 
			want: []string{"hello", "world"},
		},
		{
			name: "uppercase test",
			input: "Charmander Bulbasaur PIKACHU",
			want: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			name: "uppercase extra whitespaces",
			input: " CHARMANDER  PiKAchU        ",
			want: []string{"charmander", "pikachu"},
		},
	}
	
	for _, tc := range cases {
		got := cleanInput(tc.input)
		if len(got) != len(tc.want) {
			t.Errorf("%s - expected: %v, got: %v", tc.name, tc.want, got)
			continue
		}
		
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s - expectec: %v, got: %v", tc.name, tc.want, got)
		}
	}
}