BIN = poliright_api

default: 
		go build -o $(GOPATH)/src/poliright_api/bin/$(BIN) *.go

deps: 
		go get -u github.com/golang/lint/golint
		go get -u github.com/nsf/gocode

clean:
		go clean


.PHONY: clean install

