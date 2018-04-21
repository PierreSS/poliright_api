package main

import "time"

//Environnement variables
type env struct {
	PortSocket     string `yaml:"portSocket"`
	PortWebRequest string `yaml:"portWebRequest"`
}

var (
	version = "1.0.0"
	build   = time.Now()
)
