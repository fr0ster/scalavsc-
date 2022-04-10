#include "fib/fib.h"

uint1024_t Fib::fib(uint1024_t n)
{
	if (n < 2)
		return n;
	else
		return fib(n - 1) + fib(n - 2);
}