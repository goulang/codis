package models

import "encoding/json"

type Group struct {
	Id        int      `json:"id"`
	Servers   []string `json:"servers"`
	Promoting bool     `json:"promoting,omitempty"`
}

func (g *Group) ToJson() string {
	b, err := json.MarshalIndent(g, "", "    ")
	if err != nil {
		return "{}"
	}
	return string(b)
}
