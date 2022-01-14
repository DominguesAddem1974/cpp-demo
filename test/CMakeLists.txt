cmake_minimum_required(VERSION 2.8.12)


set(TEST_1 gtest_test)
add_executable(${TEST_1} ${TEST_1}.cpp)
target_link_libraries(${TEST_1} ${CONAN_LIBS})

add_test(NAME ${TEST_1}
    COMMAND ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/${TEST_1}
    )