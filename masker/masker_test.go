package masker

import (
	"testing"
)

func TestChangeDefaultMask(t *testing.T) {
	type args struct {
		mask string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Set default mask",
			args: args{
				"*",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChangeDefaultMask(tt.args.mask)
		})
	}
}

func TestReplacer(t *testing.T) {
	type args struct {
		str     string
		overlay string
		start   int
		end     int
	}
	tests := []struct {
		name          string
		args          args
		wantOverlayed string
	}{
		{
			name: "Count equal zero",
			args: args{
				str: "",
			},
			wantOverlayed: "",
		},
		{
			name: "Count equal zero",
			args: args{
				str:   "testing",
				start: -1,
				end:   -2,
			},
			wantOverlayed: "testing",
		},
		{
			name: "Count validatin start more than end",
			args: args{
				str:   "testing",
				start: 1,
				end:   -2,
			},
			wantOverlayed: "esting",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOverlayed := Replacer(tt.args.str, tt.args.overlay, tt.args.start, tt.args.end); gotOverlayed != tt.wantOverlayed {
				t.Errorf("Replacer() = %v, want %v", gotOverlayed, tt.wantOverlayed)
			}
		})
	}
}
