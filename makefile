ctest:
	gcc -o test ./src/test.c && ./test

gtest:
	GODEBUG=cgocheck=0 go run ./src/main.go 

clean:
	rm -rf test