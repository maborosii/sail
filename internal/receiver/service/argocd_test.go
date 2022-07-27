package service

import (
	"fmt"
	"testing"
)

func TestArgocdNotifyRequest_Spread(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{keys: []string{"source", "app_name"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArgocdNotifyRequest{
				&CommonRequest{
					"1", 2, "3",
				},
				ArgocdEventData{
					"dg", "test_app", "test", "test",
				},
			}

			got, _ := a.Spread(tt.args.keys...)
			fmt.Printf("%+v\n", got)

		})
	}
}
