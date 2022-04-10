#ifndef TEST_CPP_FACTORIAL_H
#define TEST_CPP_FACTORIAL_H

#include <boost/multiprecision/cpp_int.hpp>
using namespace boost::multiprecision;

class Factorial {
public:
    static uint1024_t factorial(uint1024_t n);
protected:
    static uint1024_t _factorial(uint1024_t akk, uint1024_t n);
};

#endif //TEST_CPP_FACTORIAL_H