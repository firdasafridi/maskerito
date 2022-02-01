package masker

import "testing"

func TestName(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "Long word name data",
			args: args{
				input: "Firda Safridi",
			},
			wantOutput: "F**da S**ridi",
		},
		{
			name: "Short name word data",
			args: args{
				input: "Jam Hard",
			},
			wantOutput: "J**m H**d",
		},
		{
			name: "Short name word data",
			args: args{
				input: "Ben",
			},
			wantOutput: "B**n",
		},
		{
			name: "Very Short name word data",
			args: args{
				input: "A Firda",
			},
			wantOutput: "** F**da",
		},
		{
			name: "Very Short name word data",
			args: args{
				input: "A Firda",
			},
			wantOutput: "** F**da",
		},
		{
			name: "Validation",
			args: args{
				input: "",
			},
			wantOutput: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Name(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("Name() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
