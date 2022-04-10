#include <iostream>
#include "Formula.h"
#include "factorial/factorial.h"
#include "fib/fib.h"

int main() {
    std::cout << "Bla: " << Formula::bla(2) << std::endl;
    std::cout << "Factorial: " << Factorial::factorial(120) << std::endl;
    std::cout << "Fib: " << Fib::fib(5) << std::endl;
    std::cout << "Fib: " << Fib::fib(10) << std::endl;
    return 0;
}
