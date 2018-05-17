package main

import (
	"net"
	"time"
)

var (
	version = "1.0.0"
	build   = time.Now()
	con     net.Conn
)

//Environnement variables
type env struct {
	PortSocket     string `yaml:"portSocket"`
	PortWebRequest string `yaml:"portWebRequest"`
}

type ia struct {
	RelationBetween [][]interface{} `json:"relation_between"`
	ImportantWords  []string        `json:"important_words"`
	Politician      []string        `json:"politician"`
	Sentiment       int             `json:"sentiment"`
	Topic           string          `json:"topic"`
	FinalDeduction  []interface{}   `json:"final_deduction"`
	Error           string
}

type er struct {
	er string
}
