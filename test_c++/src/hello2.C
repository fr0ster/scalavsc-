#include <boost/multiprecision/cpp_int.hpp>
#include <iostream>

int imax = std::numeric_limits<int>::max() / 10;

int main(int argc, const char *argv[])
{
  using boost::multiprecision::cpp_int;
  cpp_int sum = 0;
  for (auto i = 1; i <= imax; i++) sum += i;
  std::cout << "Summa from 1 to "<< imax <<" == " << sum << std::endl;

  return 0;
}
