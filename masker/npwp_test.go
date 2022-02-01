package masker

import "testing"

func TestNPWP(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Long word npwp data",
			args: args{
				input: "08.123.123.12-123.322",
			},
			wantOutput: "08.123*********23.322",
		},
		{
			name: "Short word npwp data",
			args: args{
				input: "08",
			},
			wantOutput: "0*********",
		},
		{
			name: "Validation word npwp data",
			args: args{
				input: "",
			},
			wantOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := NPWP(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("NPWP() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
