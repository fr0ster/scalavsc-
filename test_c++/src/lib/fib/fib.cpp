#include "fib/fib.hpp"

// A tail recursive function to
// calculate n th fibonacci number
uint1024_t Fib::fib(int n, uint1024_t a, uint1024_t b)
{
    if (n == 0)
        return a;
    if (n == 1)
        return b;
    return fib(n - 1, b, a + b);
}