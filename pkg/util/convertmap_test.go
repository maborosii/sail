package util

import (
	"fmt"
	"testing"
)

func TestStructToMap(t *testing.T) {
	m := struct {
		C *struct {
			AA string `json:"aa"`
			BB string `json:"bb"`
		}
		A string `json:""`
		B string `json:""`
	}{A: "1", B: "2", C: &struct {
		AA string `json:"aa"`
		BB string `json:"bb"`
	}{AA: "3", BB: "4"}}
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
