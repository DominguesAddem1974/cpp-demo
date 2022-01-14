default_target: all
.PHONY : default_target

prepare:
	@mkdir build &
	cd build && rm CMakeCache.txt &

cmake: prepare
	cd build && cmake ..

build: cmake
	cd build && make

test: build
	cd build && ctest --verbose

clean:
	@cd build && make clean

all: build
.PHONY : all