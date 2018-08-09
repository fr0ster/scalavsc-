#include <iostream>
#include <limits>
#include "InfInt.h"

int imax = std::numeric_limits<int>::max() / 10;

int main(int argc, const char *argv[])
{
  InfInt sum = 0;
  for (auto i = 1; i <= imax; i++) sum += i;
  std::cout << "Summa from 1 to "<< imax <<" == " << sum << std::endl;

  return 0;
}
