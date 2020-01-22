package main

import (
	"reflect"
	"testing"
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
