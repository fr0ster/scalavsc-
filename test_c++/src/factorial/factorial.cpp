#include "factorial/factorial.h"

long long int Factorial::_factorial(long long int akk, long long int n)
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

long long int Factorial::factorial(long long int n)
{
    return _factorial(1, n);
}