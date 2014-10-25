package = scramble

.PHONY: all clean doc

all:
	go build ./...

clean:
	rm -f $(package)

doc:
	pandoc -s -t man $(package).1.md -o $(package).1
