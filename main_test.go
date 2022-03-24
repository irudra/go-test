package main

import (
	"reflect"
	"testing"
)

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

func Test_wholeStory(t *testing.T) {
	var tests = []struct {
		name string
		args string
		want string
	}{
		{
			name: "withValueFormat",
			args: "12-hello-34-world",
			want: "hello world",
		},
		{
			name: "withInvalidbutStringFirst",
			args: "dsafds-12-dfasdf-34",
			want: "",
		},
		{
			name: "withBlank",
			args: "   ",
			want: "",
		},
		{
			name: "withNoDash",
			args: "12dfad12sdfd123sdfdsf34dfasf",
			want: "",
		},
		{
			name: "withNoDash",
			args: "dfad_fdadf_",
			want: "",
		},
		{
			name: "withBlankInBetween",
			args: "12-dsfa-12-  sdfsa-12",
			want: "",
		},
		{
			name: "withNumericAsAss",
			args: "12-hello-12-world-12-number-23-3456",
			want: "hello world number 3456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := wholeStory(tt.args); got != tt.want {
				t.Errorf("wholeStory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageNumber(t *testing.T) {
	var tests = []struct {
		name string
		args string
		want uint64
	}{
		{
			name: "withValueFormat",
			args: "12-hello-34-world",
			want: 23,
		},
		{
			name: "withInvalidbutStringFirst",
			args: "dsafds-12-dfasdf-34",
			want: 0,
		},
		{
			name: "withBlank",
			args: "   ",
			want: 0,
		},
		{
			name: "withNoDash",
			args: "12dfad12sdfd123sdfdsf34dfasf",
			want: 0,
		},
		{
			name: "withNoDash",
			args: "dfad_fdadf_",
			want: 0,
		},
		{
			name: "withBlankInBetween",
			args: "12-dsfa-12- sdfsa-12",
			want: 0,
		},
		{
			name: "withNumericAsAss",
			args: "12-hello-12-world-12-number-23-3456",
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := averageNumber(tt.args); got != tt.want {
				t.Errorf("averageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storyStats(t *testing.T) {

	tests := []struct {
		name  string
		args  string
		want  string
		want1 string
		want2 int
		want3 []string
	}{
		{
			name:  "testShortNumber",
			args:  "12-hello-12-world-12-number-23-3456",
			want:  "3456",
			want1: "number",
			want2: 5,
			want3: []string{"hello", "world"},
		},
		{
			name:  "testNoAverageNumber",
			args:  "12-hello-12-world-12-number-23-3",
			want:  "3",
			want1: "number",
			want2: 4,
			want3: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := storyStats(tt.args)
			if got != tt.want {
				t.Errorf("storyStats() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("storyStats() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("storyStats() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("storyStats() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
