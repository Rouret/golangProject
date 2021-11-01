package models

type Average struct {
	Date  string `json:date`
	Value string `json:Value`
}

type Averages []Average