#include <boost/multiprecision/cpp_int.hpp>
#include <iostream>


int main(int argc, const char *argv[])
{
using boost::multiprecision::cpp_int;
  cpp_int sum = 0;
  int imax = std::numeric_limits<int>::max() / 100;
  for (auto i = 1; i <= imax; i++) sum += i;
  std::cout << "Summa from 1 to "<< imax <<" == " << sum << std::endl;
  //std::cout << "Hello, World!!!" << std::endl;

  return 0;
}
