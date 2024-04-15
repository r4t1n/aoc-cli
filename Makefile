target = aoc
target_path = /usr/local/bin/aoc

build:
	go build -o $(target)

check-root:
	@if [ `id -u` -ne 0 ]; then \
		echo "Root privileges needed"; \
		exit 1; \
	fi

clean:
	rm -f $(target)

install: check-root
	cp $(target) $(target_path)

uninstall: check-root
	rm -f $(target_path)