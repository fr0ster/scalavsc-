#include <iostream>
#include <limits>
#include "InfInt.h"

int main(int argc, const char *argv[])
{
  InfInt sum = 0;
  int imax = std::numeric_limits<int>::max();
  for (auto i = 1; i <= imax; i++) sum += i;
  std::cout << "Summa from 1 to "<< imax <<" == " << sum << std::endl;
  //std::cout << "Hello, World!!!" << std::endl;

  return 0;
}
