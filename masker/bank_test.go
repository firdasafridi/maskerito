package masker

import "testing"

func TestBankAccountNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Long word bank account Number data",
			args: args{
				input: "7123123123",
			},
			wantOutput: "7123****23",
		},
		{
			name: "Short word bank account Number data",
			args: args{
				input: "72",
			},
			wantOutput: "7****",
		},
		{
			name: "Validation word bank account Number data",
			args: args{
				input: "",
			},
			wantOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := BankAccountNumber(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("BankAccountNumber() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
