#include "factorial/factorial.h"

uint1024_t Factorial::_factorial(uint1024_t akk, uint1024_t n)
{
    if (n == 0)
    {
        return akk;
    }
    else
    {
        return _factorial(akk * n, n - 1);
    }
}

uint1024_t Factorial::factorial(uint1024_t n)
{
    return _factorial(1, n);
}