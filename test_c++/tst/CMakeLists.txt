set(BINARY ${CMAKE_PROJECT_NAME})

find_package(PkgConfig REQUIRED)
pkg_check_modules(gtest REQUIRED gtest>=1.8.1)

file(GLOB_RECURSE TEST_SOURCES LIST_DIRECTORIES false *.h *.cpp)
file(GLOB TEST_MAIN LIST_DIRECTORIES false *.h *.cpp)

add_executable(${BINARY}_tst ${TEST_SOURCES})

add_test(NAME ${BINARY} COMMAND ${BINARY}_tst)

target_link_libraries(${BINARY}_tst PUBLIC ${CMAKE_PROJECT_NAME}_lib gtest)

add_subdirectory(factorial)
add_subdirectory(fib)
add_subdirectory(formula)