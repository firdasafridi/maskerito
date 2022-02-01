package masker

import "testing"

func TestID(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Long word id data",
			args: args{
				input: "CAZ8123123",
			},
			wantOutput: "CAZ812****",
		},
		{
			name: "Medium word id data",
			args: args{
				input: "B12345CAZ",
			},
			wantOutput: "B12345****",
		},
		{
			name: "Short word id data",
			args: args{
				input: "B123",
			},
			wantOutput: "B1****",
		},
		{
			name: "Validation id data",
			args: args{
				input: "",
			},
			wantOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := ID(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("ID() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
