cmake_minimum_required(VERSION 2.8.12)

add_executable(${PROJECT} main.cpp)
target_link_libraries(${PROJECT} ${CONAN_LIBS})