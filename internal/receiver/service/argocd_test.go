package service

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArgocdNotifyRequest_Spread(t *testing.T) {
	type fields struct {
		CommonRequest *CommonRequest
		EventData     ArgocdEventData
	}
	type args struct {
		tagType string
		keys    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				&CommonRequest{
					"1", 2, "3",
				},
				ArgocdEventData{
					"dg", "test_app", "test", "test",
				},
			},
			args: args{
				tagType: "json",
				keys:    []string{"source", "app_name"},
			},
			want:    map[string]interface{}{"source": "dg", "app_name": "test_app"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArgocdNotifyRequest{
				CommonRequest: tt.fields.CommonRequest,
				EventData:     tt.fields.EventData,
			}
			got, err := a.Spread(tt.args.tagType, tt.args.keys...)
			fmt.Printf("%+v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArgocdNotifyRequest.Spread() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgocdNotifyRequest.Spread() = %v, want %v", got, tt.want)
			}
		})
	}
}
