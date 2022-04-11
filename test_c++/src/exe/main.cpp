#include <iostream>
#include "Formula.hpp"
#include "factorial/factorial.hpp"
#include "fib/fib.hpp"
#include "sc/freelock/sum.hpp"
#include "sc/freelock/summator.hpp"
#include <boost/multiprecision/cpp_int.hpp>

using namespace boost::multiprecision;

int main()
{
    std::cout << "Bla: " << Formula::bla(2) << std::endl;
    std::cout << "Factorial: " << Factorial<int, uint1024_t>::factorial(120) << std::endl;
    std::cout << "Factorial: " << Factorial<int, uint64_t>::factorial(40) << std::endl;
    std::cout << "Fib: " << Fib::fib(5) << std::endl;
    std::cout << "Fib: " << Fib::fib(10) << std::endl;
    std::vector<int> v(1000000,1);
    std::cout << "SC Sum: " << SC::Freelock::Sum<int, int>::sum(v) << std::endl;
    std::cout << "SC Sum: " << SC::Freelock::Summator<int, int>::summator(v) << std::endl;
    return 0;
}
