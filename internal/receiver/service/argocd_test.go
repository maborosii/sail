package service

import (
	"fmt"
	"reflect"
	"sail/internal/model"
	"testing"
)

func TestArgocdNotifyRequest_Spread(t *testing.T) {
	type fields struct {
		CommonRequest *model.CommonRequest
		EventData     model.ArgocdEventData
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
				CommonRequest: &model.CommonRequest{
					Type: "1", OccurAt: 2, Operator: "3",
				},
				EventData: model.ArgocdEventData{
					Source: "dg", AppName: "test_app", SyncStatus: "test", HealthStatus: "test",
				},
			},
			args: args{
				tagType: "mapstructure",
				keys:    []string{"city", "app_name"},
			},
			want:    map[string]interface{}{"city": "dg", "app_name": "test_app"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &model.ArgocdNotifyRequest{
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
