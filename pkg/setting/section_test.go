package setting

import (
	"fmt"
	"testing"

	dtt "sail/pkg/dingtalk"

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
			s, _ := NewSetting("../../configs")
			if err := s.ReadConfig(args); (err != nil) != tt.wantErr {
				t.Errorf("Setting.ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			d := &dtt.DingTalkConfig{}
			for k, v := range args.Senders {
				if k == "dingtalk" {
					// parse dingtalk config
					if err := mapstructure.Decode(v, d); err != nil {
						panic(err)
					}
					fmt.Printf("%+v\n", d)
				}
			}
			dc := &dtt.DingTalkMessageTemplate{}
			for _, j := range args.Template {
				// parse template config
				for _, jj := range j {
					if err := mapstructure.Decode(jj, dc); err != nil {
						panic(err)
					}
					fmt.Printf("%+v\n", dc)
				}
				// mapstructure.Decode(j, dc)
				// fmt.Printf("%+v\n", dc)
			}

		})
	}
}
