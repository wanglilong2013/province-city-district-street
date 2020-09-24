package parser

import (
	"encoding/json"
)

type RespData struct {
	Status    string
	Info      string
	Districts []District
}

type District struct {
	CityCode  interface{}
	AdCode    string
	Name      string
	Center    string
	Level     string
	Districts []District
}

func Parse(content []byte) (*RespData, error) {
	resp := &RespData{}
	err := json.Unmarshal(content, resp)
	if err != nil {
		return resp, nil
	}
	return resp, nil
}
