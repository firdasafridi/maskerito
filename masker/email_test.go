package masker

import "testing"

func TestEmail(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Long email id data",
			args: args{
				input: "sample.coba@gmail.com",
			},
			wantOutput: "s****coba@gmail.com",
		},
		{
			name: "Long email id data",
			args: args{
				input: "sa@gmail.com",
			},
			wantOutput: "s****@gmail.com",
		},
		{
			name: "Invalid email id data",
			args: args{
				input: "sagmail.com",
			},
			wantOutput: "s****.com",
		},
		{
			name: "Validation email id data",
			args: args{
				input: "",
			},
			wantOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Email(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("Email() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
