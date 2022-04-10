#include "gtest/gtest.h"
#include "fib/fib.h"

TEST(fibTest, test3) {
    //arrange
    //act
    //assert
    EXPECT_EQ (Fib::fib (0),  0);
    EXPECT_EQ (Fib::fib (5), 5);
    EXPECT_EQ (Fib::fib (10), 55);
}
