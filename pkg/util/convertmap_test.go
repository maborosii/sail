package util

import (
	"fmt"
	"testing"
)

func TestStructToMap(t *testing.T) {
	m := struct {
		C *struct {
			AA string `json:"aa"`
			BB []struct {
				DD string `json:"dd"`
			} `json:"bb"`
		}
		A string `json:""`
		B string `json:""`
	}{
		A: "1",
		B: "2",
		C: &struct {
			AA string `json:"aa"`
			BB []struct {
				DD string `json:"dd"`
			} `json:"bb"`
		}{
			AA: "3",
			BB: []struct {
				DD string `json:"dd"`
			}{{DD: "5"}, {DD: "6"}},
		}}
	tests := []struct {
		name string
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got := StructToMap(m)
			got, _ := SpreadToMap(m, "json")
			fmt.Printf("%+v\n", got)
		})
	}
}
