run:
	@go build -o libfetch.so -buildmode=c-shared internal/fetch.go
	@gcc main.c -L. -lfetch -o main
	@./main