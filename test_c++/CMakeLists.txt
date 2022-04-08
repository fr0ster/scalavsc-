cmake_minimum_required(VERSION 3.10)
project(test_c++)

set(CMAKE_CXX_STANDARD 20)

include_directories(include include/factorial)

# add_subdirectory(benchmark)
add_subdirectory(src)
add_subdirectory(tst)
add_subdirectory(lib/googletest)
