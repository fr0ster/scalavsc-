set(BINARY ${CMAKE_PROJECT_NAME})

find_package(PkgConfig REQUIRED)
pkg_check_modules(gtest REQUIRED gtest>=1.8.1)

file(GLOB_RECURSE TEST_SOURCES LIST_DIRECTORIES false *.h *.cpp)

set(SOURCES ${TEST_SOURCES})

add_executable(${BINARY}_fib_tst ${TEST_SOURCES})

add_test(NAME ${BINARY}_fib COMMAND ${BINARY}_fib_tst)

target_link_libraries(${BINARY}_fib_tst PUBLIC ${CMAKE_PROJECT_NAME}_lib gtest)