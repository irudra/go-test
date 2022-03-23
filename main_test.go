package main

import "testing"

func Test_testValidity(t *testing.T) {
	var tests = []struct {
		name string
		args string
		want bool
	}{
		{
			name: "withValueFormat",
			args: "12-dafd-34-dfa",
			want: true,
		},
		{
			name: "withInvalidbutStringFirst",
			args: "dsafds-12-dfasdf-34",
			want: false,
		},
		{
			name: "withBlank",
			args: "   ",
			want: false,
		},
		{
			name: "withNoDash",
			args: "12dfad12sdfd123sdfdsf34dfasf",
			want: false,
		},
		{
			name: "withNoDash",
			args: "dfad_fdadf_",
			want: false,
		},
		{
			name: "withBlankInBetween",
			args: "12-dsfa-12-  sdfsa-12",
			want: false,
		},
		{
			name: "withNumericAsAss",
			args: "12-dsfa-12-sdfsa-12-1256564564554",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testValidity(tt.args); got != tt.want {
				t.Errorf("testValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}
