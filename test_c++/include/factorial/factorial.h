#ifndef TEST_CPP_FACTORIAL_H
#define TEST_CPP_FACTORIAL_H

#include <boost/multiprecision/cpp_int.hpp>
using namespace boost::multiprecision;

template <typename A, typename B>
class Factorial
{
public:
    static B factorial(A n)
    {
        return _factorial(1, n);
    };

protected:
    static B _factorial(const B akk, const A n)
    {
        if (n == 0)
            return akk;
        return _factorial(akk * n, n - 1);
    };
};

#endif // TEST_CPP_FACTORIAL_H