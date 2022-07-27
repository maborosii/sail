package setting

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func TestSetting_ReadConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	args := &Config{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewSetting("../../config")
			if err := s.ReadConfig(args); (err != nil) != tt.wantErr {
				t.Errorf("Setting.ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			d := &DingTalkConfig{}
			for k, v := range args.Senders {
				if k == "dingtalk" {
					// parse dingtalk config
					mapstructure.Decode(v, d)
					fmt.Printf("%+v\n", d)
				}
			}
			dc := &DingTalkMessageTemplate{}
			for _, j := range args.Template {
				// parse template config
				mapstructure.Decode(j, dc)
				fmt.Printf("%+v\n", dc)
			}

		})
	}
}
