package main

import (
	"math/rand"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// getUniqueSlice
// return form any slice only uniq values
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

// randSeq generate N length random string
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
