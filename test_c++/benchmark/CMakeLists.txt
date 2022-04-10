set(BINARY ${CMAKE_PROJECT_NAME})

find_package(PkgConfig REQUIRED)
pkg_check_modules(benchmark REQUIRED benchmark>=1.5.2)

file(GLOB_RECURSE TEST_SOURCES LIST_DIRECTORIES false *.h *.cpp)

set(SOURCES ${TEST_SOURCES})

add_executable(${BINARY}_bm ${TEST_SOURCES})

target_link_libraries(${BINARY}_bm PUBLIC ${CMAKE_PROJECT_NAME}_lib benchmark)
