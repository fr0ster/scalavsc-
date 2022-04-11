#include "gtest/gtest.h"
#include "fibbonacci/fibbonacci.hpp"
#include <boost/multiprecision/cpp_int.hpp>

using namespace boost::multiprecision;

TEST(fibTest, test3) {
    //arrange
    //act
    //assert
    auto fb0 = Fibbonacci<int, uint1024_t>::fibbonacci (0);
    auto fb5 = Fibbonacci<int, uint1024_t>::fibbonacci (5);
    auto fb10 = Fibbonacci<int, uint1024_t>::fibbonacci (10);
    EXPECT_EQ (fb0,  0);
    EXPECT_EQ (fb5, 5);
    EXPECT_EQ (fb10, 55);
}
