file(GLOB_RECURSE FACTORIAL_BM_SOURCES LIST_DIRECTORIES false *.h *.cpp)

set(FACTORIAL_SOURCES ${BM_MAIN} ${FACTORIAL_BM_SOURCES})

add_executable(${BINARY}_factorial_bm ${FACTORIAL_SOURCES})

target_link_libraries(${BINARY}_factorial_bm PUBLIC ${CMAKE_PROJECT_NAME}_lib benchmark)