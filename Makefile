.PHONY: benchmark time test travis_install travis vet

test:
	go test -v ./...

vet: 
	go vet

benchmark:
	go test -bench .

time:
	cat data.csv | time -p tree , > /dev/null


travis_install:
	go get -d -v ./... && go build -v ./...

travis: travis_install test benchmark vet

