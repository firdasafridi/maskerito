package masker

import "testing"

func TestSIUP(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Long word siup data",
			args: args{
				input: "510/20/ABC/1888",
			},
			wantOutput: "08.123*********23.322",
		},
		{
			name: "Short word siup data",
			args: args{
				input: "08",
			},
			wantOutput: "0****",
		},
		{
			name: "Validation word siup data",
			args: args{
				input: "",
			},
			wantOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := SIUP(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("SIUP() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
