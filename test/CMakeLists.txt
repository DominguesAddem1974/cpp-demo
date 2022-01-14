cmake_minimum_required(VERSION 2.8.12)

set(TEST gtest_test)
add_executable(${TEST} ${TEST}.cpp)
target_link_libraries(${TEST} ${CONAN_LIBS} ${CONAN_LIBS} ${CMAKE_ARCHIVE_OUTPUT_DIRECTORY}/lib${PROJECT}_static.a)

add_test(NAME ${TEST}
    COMMAND ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/${TEST})
