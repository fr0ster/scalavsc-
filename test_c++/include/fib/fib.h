#ifndef TEST_CPP_FIB_H
#define TEST_CPP_FIB_H

#include <boost/multiprecision/cpp_int.hpp>
using namespace boost::multiprecision;

class Fib {
public:
    // static uint1024_t fib(uint1024_t n);
    static uint1024_t fib(int n, uint1024_t a = 0, uint1024_t b = 1);
};

#endif //TEST_CPP_FIB_H