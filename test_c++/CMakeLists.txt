cmake_minimum_required(VERSION 3.10)
project(test_c++)
find_package(benchmark REQUIRED)

set(CMAKE_CXX_STANDARD 20)
# SET(CMAKE_CXX_FLAGS  "-pedantic -Wall -Wextra -lbenchmark -lpthread -fopenmp")
SET(CMAKE_CXX_FLAGS -pthread)

include_directories(include include/factorial)

add_subdirectory(benchmark)
add_subdirectory(src)
add_subdirectory(tst)