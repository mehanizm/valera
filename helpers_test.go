package main

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

type TestCaseGetUniqueMap struct {
	Input  []string
	Output []string
}

func TestGetUniqueMap(t *testing.T) {
	cases := []TestCaseGetUniqueMap{
		TestCaseGetUniqueMap{
			[]string{"123", "123", "223"},
			[]string{"123", "223"},
		},
		TestCaseGetUniqueMap{
			[]string{"123", "123"},
			[]string{"123"},
		},
		TestCaseGetUniqueMap{
			[]string{"hflf", "dfiae", "fdaef"},
			[]string{"hflf", "dfiae", "fdaef"},
		},
	}

	for caseNum, item := range cases {
		res := getUniqueSlice(item.Input)
		eq := reflect.DeepEqual(res, item.Output)
		if !eq {
			t.Errorf("Mistake in case number [%d]", caseNum+1)
		}
	}
}

func Test_randSeq(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		wantLength int
	}{
		{
			name:       "Case 1. Check lenth",
			n:          10,
			wantLength: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randSeq(tt.n); utf8.RuneCountInString(got) != tt.wantLength {
				t.Errorf("randSeq() = %v, wantLength %v", tt.n, tt.wantLength)
			}
		})
	}
}
