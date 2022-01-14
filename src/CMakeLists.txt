cmake_minimum_required(VERSION 2.8.12)

set(LIBRARY_SOURCE function.cpp)

add_library(${PROJECT}_shared SHARED ${LIBRARY_SOURCE})
add_library(${PROJECT}_static STATIC ${LIBRARY_SOURCE})

add_executable(${PROJECT} main.cpp)
target_link_libraries(${PROJECT} ${CONAN_LIBS} ${CMAKE_ARCHIVE_OUTPUT_DIRECTORY}/lib${PROJECT}_static.a)

