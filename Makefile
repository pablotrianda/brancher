build:
	docker run --rm -v $(CURDIR):/app/brancher -w /app/brancher golang:1.16 && go install && go build -v

