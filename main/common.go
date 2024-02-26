package main

import (
	"encoding/json"
	"strings"
)

func AsJSON(item any) string {
	js, err := json.Marshal(item)
	if err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

func MapF[A any, B any](items []A, f func(A) (B, error)) ([]B, error) {
	var out []B
	for _, item := range items {
		result, err := f(item)
		if err != nil {
			return nil, err
		}
		out = append(out, result)
	}
	return out, nil
}

func ToByte(value bool) byte {
	if value {
		return 1
	} else {
		return 0
	}
}

func Trim(s string) string {
	return strings.Trim(s, " \t\n\r\b")
}
