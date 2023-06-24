build:
	@go build -o wx

run: build
	./wx

clean:
	@rm ./wx ||: