BIN = WatchIt

default: 
		go build -o $(GOPATH)/src/poliright_api/bin/$(BIN) src/*.go

deps: 
		go get -u github.com/golang/lint/golint
		go get -u github.com/nsf/gocode

clean:
		go clean


.PHONY: clean install

