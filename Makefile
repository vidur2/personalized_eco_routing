ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# PHONY means that it doesn't correspond to a file; it always runs the build commands.

.PHONY: build
build:
	export CC=gcc
	cd lib/regression && cargo build --release
	cp lib/regression/target/release/libregression.a lib/
	go build line_integrals_fuel_efficiency

.PHONY: run
run:
	RUST_LOG=trace ./line_integrals_fuel_efficiency

.PHONY: test
test:
	make test-rust-lib && make test-go-lib

# This is just for running the Rust lib tests natively via cargo.
.PHONY: test-rust-lib
test-rust-lib:
	cd lib/regression && RUST_LOG=trace cargo test -- --nocapture

.PHONY: check
check:
	cd lib/regression && cargo check

# Running go test 
.PHONY: test-go-lib
test-go-lib:
	cd tests && go test

.PHONY: build-raspi
build-raspi:
	 CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o line_integrals_fuel_efficiency -ldflags="-extld=$CC"

.PHONY: clean
clean:
	rm -rf main_shared main_static lib/libregression.so lib/libregression.a lib/regression/target