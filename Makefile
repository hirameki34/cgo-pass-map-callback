.PHONY: so

so:
	go build -o libdemo.so -buildmode=c-shared *.go

run: so
	g++ wrapper.cpp main.cpp -o main.out -I. -L. -ldemo -std=c++17
	env LD_LIBRARY_PATH=. ./main.out
