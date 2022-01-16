default_target: all
.PHONY : default_target

prepare:
	@mkdir build &
	cd build && conan install ..
	cd build && rm CMakeCache.txt &

cmake: prepare
	cd build && cmake .. -DCMAKE_BUILD_TYPE=Release

build: cmake
	cd build && make

test: build
	cd build && ctest --verbose

clean:
	@cd build && make clean

go-binding: release
	mkdir -p ./bindings/go/release
	cp -r build/{bin,lib} ./bindings/go/release/
	mkdir -p ./bindings/go/include
	cp -r include ./bindings/go/
	@# CGO_CFLAGS="-I${PWD}/include" CGO_LDFLAGS="-I${PWD}/release/lib" go build -o release/go-bindings/go-demo ./bindings/go/demo.go
	go build -o release/go-bindings/go-demo ./bindings/go/demo.go

go-test: go-binding
	cd ./bindings/go/ && go test -v

go-benchmark: go-binding
	cd ./bindings/go/ && go test -benchmem -bench .

release: build
	mkdir -p release
	cp -r build/{bin,lib} release/

bindings: go-binding
.PHONY : bindings

all: build bindings
.PHONY : all