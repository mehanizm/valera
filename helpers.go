package main

import "strings"

func getUniqueSlice(sl []string) []string {
	tempMap := make(map[string]struct{}, len(sl))
	res := make([]string, 0)
	for _, s := range sl {
		s := strings.Trim(s, " \n$")
		if _, ok := tempMap[s]; !ok {
			tempMap[s] = struct{}{}
			res = append(res, s)
		}
	}
	return res
}
