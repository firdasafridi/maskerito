package masker

import "testing"

func TestAddress(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Long word address data",
			args: args{
				input: "Jln Jati IV Blok A Depok",
			},
			wantOutput: "Jln Ja********",
		},
		{
			name: "Short word address data",
			args: args{
				input: "Jln A",
			},
			wantOutput: "Jl********",
		},
		{
			name: "Very Short word address data",
			args: args{
				input: "A",
			},
			wantOutput: "A********",
		},
		{
			name: "Validation word address data",
			args: args{
				input: "",
			},
			wantOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Address(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("Address() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
