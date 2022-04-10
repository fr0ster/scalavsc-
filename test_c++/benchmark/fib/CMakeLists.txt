set(BINARY ${CMAKE_PROJECT_NAME}_fib_bm)

file(GLOB_RECURSE BM_SOURCES LIST_DIRECTORIES false *.h *.cpp)

add_executable(${BINARY} ${BM_MAIN} ${BM_SOURCES})

target_link_libraries(${BINARY} PUBLIC ${CMAKE_PROJECT_NAME}_lib benchmark)