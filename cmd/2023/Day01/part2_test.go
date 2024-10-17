package main

import "testing"

func Test_replaceNamedNumbersWithNumerics(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_1",
			args: args{
				s: "one9nine",
			},
			want: 19,
		},
		{
			name: "test_2",
			args: args{
				s: "oneight7",
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := performExtractionWithNumericNames(tt.args.s); got != tt.want {
					t.Errorf("performExtractionWithNumericNames() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
