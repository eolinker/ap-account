package service

import "strings"

type Gender uint8

func (g *Gender) String() string {
	if *g < 0 || *g >= maxGender {
		return "unknown"
	}

	return genderForStr[*g]
}
func (g *Gender) MarshalJSON() ([]byte, error) {
	return []byte(g.String()), nil
}
func (g *Gender) UnmarshalJSON(bytes []byte) error {
	v := string(bytes)
	gv, h := genderOfValue[strings.ToLower(v)]
	if h {
		*g = gv
	} else {
		*g = Unknown
	}
	return nil
}

const (
	Unknown Gender = iota
	Male
	Female
	maxGender
)

var (
	genderOfValue = make(map[string]Gender)
	genderForStr  = make([]string, maxGender)
)

func init() {
	genderForStr[Unknown] = "Unknown"
	genderForStr[Male] = "Male"
	genderForStr[Female] = "Female"

	for i := Unknown; i < maxGender; i++ {
		genderOfValue[strings.ToLower(genderForStr[i])] = i
	}

}
