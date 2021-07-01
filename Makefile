# Make will use bash instead of sh
SHELL := /usr/bin/env bash
# use clang as required by bazel
CC=clang
# GNU make man page
# http://www.gnu.org/software/make/manual/make.html

.PHONY: help
help:
	@echo ''
	@echo "Development targets: "
	@echo '    make build   		Builds the code base incrementally (fast).'
	@echo '    make rebuild   		Regenerates all dependencies & builds code base (slow).'
	@echo '    make run   			Runs the  project binary.'

.PHONY: build
build:
	@source scripts/build.sh

.PHONY: rebuild
rebuild:
	@source scripts/rebuild.sh

.PHONY: run
run:
	@source scripts/run.sh
