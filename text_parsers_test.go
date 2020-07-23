package main

import (
	"reflect"
	"testing"
)

type TestCaseParseIssueKeysFromMsg struct {
	Input  string
	Output []string
}

func TestParseIssueKeysFromMsg(t *testing.T) {
	cases := []TestCaseParseIssueKeysFromMsg{
		{
			`CDI-4532 - Автоматическая выгрузка полей полнотекстового и расширенного поиска Демо-заказчика
			CDI-4532 - Автоматическая выгрузка полей полнотекстового и расширенного поиска Демо-заказчика
			CDI-4532 - Автоматическая выгрузка полей полнотекстового и расширенного поиска Демо-заказчика`,
			[]string{"CDI-4532"},
		},
		// spell-checker: disable
		{
			"CDI-4532 АвтонотFSUP-33 ДобавлеекстCDI-4652 Учесть ноового и",
			[]string{
				"CDI-4532",
				"FSUP-33",
				"CDI-4652",
			},
		},
		// spell-checker: enable
		{
			"",
			[]string{},
		},
	}

	for caseNum, item := range cases {
		res := parseIssueKeysFromMsg(item.Input)
		eq := reflect.DeepEqual(res, item.Output)
		if !eq {
			t.Errorf("Mistake in case number [%d]. Was %v but expect %v", caseNum+1, res, item.Output)
		}
	}

}
