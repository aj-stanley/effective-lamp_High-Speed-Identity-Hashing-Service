.PHONY: all clean run

all:
	cd rustlib && cargo build --release
	cp rustlib/target/release/librustlib.a gobin/
	cd gobin && go build -o fastid main.go

run: all
	./gobin/fastid

clean:
	rm -f gobin/fastid gobin/librustlib.a
	cd rustlib && cargo clean
