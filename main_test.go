package main

import "testing"

func TestParser(t *testing.T) {
	tests := []struct {
		input   string
		output  float64
		wantErr bool
	}{
		{
			input:   "4",
			output:  4.0,
			wantErr: false,
		},
		{
			input:   "4 ",
			output:  4.0,
			wantErr: false,
		},

		{
			input:   "4.0",
			output:  4.0,
			wantErr: false,
		},
		{
			input:   "9i0",
			output:  0.0,
			wantErr: true,
		},
		{
			input:   "4 + 3",
			output:  7.0,
			wantErr: false,
		},
		{
			input:   "4 - 3",
			output:  1.0,
			wantErr: false,
		},
		{
			input:   "4 - 3 + 1",
			output:  2.0,
			wantErr: false,
		},
		{
			input:   "4 * 3",
			output:  12.0,
			wantErr: false,
		},
		{
			input:   "4 * 2 - 2",
			output:  6.0,
			wantErr: false,
		},
		{
			input:   "4 - 2 * 2",
			output:  0.0,
			wantErr: false,
		},
		{
			input:   "4 / 2",
			output:  2.0,
			wantErr: false,
		},
		{
			input:   "4 / 2 - 2",
			output:  0.0,
			wantErr: false,
		},
		{
			input:   "4 - 2 / 2",
			output:  3.0,
			wantErr: false,
		},
		{
			input:   "4 / 2 * 2",
			output:  4.0,
			wantErr: false,
		},
		{
			input:   "4 * 2 / 2",
			output:  4.0,
			wantErr: false,
		},
		{
			input:   "3 + 4/4 + 5*2/5 - 1",
			output:  5.0,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			parser := NewParser()
			got, err := parser.Evaluate(test.input)
			if test.wantErr && err == nil {
				t.Fatalf("%s: got no error, wanted an error", test.input)
			}
			if !test.wantErr && err != nil {
				t.Fatalf("%s: got error %v, expected no error", test.input, err)
			}
			if got != test.output {
				t.Fatalf("%s: got %v, expected %v", test.input, got, test.output)
			}
		})
	}
}
