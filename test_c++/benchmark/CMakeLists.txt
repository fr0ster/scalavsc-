set(BINARY ${CMAKE_PROJECT_NAME})

find_package(PkgConfig REQUIRED)
pkg_check_modules(benchmark REQUIRED benchmark>=1.5.2)

file(GLOB_RECURSE BM_SOURCES LIST_DIRECTORIES false *.h *.cpp)
file(GLOB_RECURSE BM_MAIN LIST_DIRECTORIES false *.h *.cpp)

add_executable(${BINARY}_bm ${BM_SOURCES})

target_link_libraries(${BINARY}_bm PUBLIC ${CMAKE_PROJECT_NAME}_lib benchmark)

add_subdirectory(factorial)