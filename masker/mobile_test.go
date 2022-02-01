package masker

import "testing"

func TestMobile(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Default mask data",
			args: args{
				input: "081311115258",
			},
			wantOutput: "0813****5258",
		},
		{
			name: "Validation empty data",
			args: args{
				input: "",
			},
			wantOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Mobile(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("Mobile() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
