#include "gtest/gtest.h"
#include "factorial/factorial.h"

TEST(factorialTest, test2) {
    //arrange
    //act
    //assert
    auto factorial0 = Factorial<int, int>::factorial (0);
    EXPECT_EQ (factorial0,  1);
    auto factorial5 = Factorial<int, int>::factorial (5);
    EXPECT_EQ (factorial5, 120);
    auto factorial10 = Factorial<int, int>::factorial (10);
    EXPECT_EQ (factorial10, 3628800);
}
