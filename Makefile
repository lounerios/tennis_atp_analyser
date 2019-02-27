#!/bin/bash

CurDir=$(pwd)
export GOPATH=$(CurDir)/tennis_atp_analyser/:$(HOME)/go

edit:
	go build -o bin/import_stats import_stats.go

