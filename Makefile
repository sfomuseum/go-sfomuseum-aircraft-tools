cli:
	go build -o bin/build-sfomuseum-data cmd/build-sfomuseum-data/main.go
	go build -o bin/build-icao-data cmd/build-icao-data/main.go

data:
	@make cli
	@make sfomuseum-data

sfomuseum-data:
	bin/build-sfomuseum-data > /usr/local/sfomuseum/go-sfomuseum-aircraft/sfomuseum/data.go
