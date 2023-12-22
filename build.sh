#!/bin/bash

# Change to the aoc-cli directory, suppressing errors
cd aoc-cli 2>/dev/null || true

# Build it
go build

# Move the binary to /usr/local/bin/aoc
sudo mv aoc-cli /usr/local/bin/aoc

echo "aoc-cli has been installed to /usr/local/bin/aoc"
