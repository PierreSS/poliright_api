BIN = poliright_api

default: 
		go build -o bin/$(BIN) src/*.go

clean:
		go clean


.PHONY: clean install

