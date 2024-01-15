DESTINATION = /usr/local/bin
TARGET = aoc

all: build

build:
	go build -o $(TARGET)

clean:
	rm -f $(TARGET)

install: $(TARGET)
	cp $(TARGET) $(DESTINATION)

.PHONY: build clean install