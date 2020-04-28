.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 go build -o bin/ListDocumentsFunction ./services/assets/get 
	GOOS=linux GOARCH=amd64 go build -o bin/CreateDocumentFunction ./services/assets/put 