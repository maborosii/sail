package validate

import (
	"testing"
	"time"
)

func TestIsExpired(t *testing.T) {
	type args struct {
		sample   int
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test case for not in time range",
			args: args{
				sample:   1655959720,
				duration: 5 * time.Minute,
			},
			want:    false,
			wantErr: true,
		}, {
			name: "test case for error",
			args: args{
				sample:   1655959720111,
				duration: 5 * time.Minute,
			},
			want:    false,
			wantErr: true,
		}, {
			name: "test case for in time range",
			args: args{
				sample:   1658743435,
				duration: 5 * time.Minute,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsExpired(tt.args.sample, tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsExpired() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}
