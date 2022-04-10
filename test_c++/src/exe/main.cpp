#include <iostream>
#include "Formula.h"
#include "factorial/factorial.h"
#include "fib/fib.h"
#include "sc/freelock/sum.h"
#include "sc/freelock/summator.h"

int main()
{
    std::cout << "Bla: " << Formula::bla(2) << std::endl;
    std::cout << "Factorial: " << Factorial::factorial(120) << std::endl;
    std::cout << "Fib: " << Fib::fib(5) << std::endl;
    std::cout << "Fib: " << Fib::fib(10) << std::endl;
    std::vector<int> v(1000000,1);
    std::cout << "SC Sum: " << SC::Freelock::Sum<int, int>::sum(v) << std::endl;
    std::cout << "SC Sum: " << SC::Freelock::Summator<int, int>::summator(v) << std::endl;
    return 0;
}
