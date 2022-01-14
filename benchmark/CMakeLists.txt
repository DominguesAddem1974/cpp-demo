cmake_minimum_required(VERSION 2.8.12)


set(TEST gtest_benchmark)
add_executable(${TEST} ${TEST}.cpp)
target_link_libraries(${TEST} ${CONAN_LIBS})

add_test(NAME ${TEST}
    COMMAND ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/${TEST})

