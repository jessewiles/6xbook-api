all: clean build run

clean:
	@rm -f ./server 

build:
	@go build -o server *.go

run:
	./server
