#include "gtest/gtest.h"
#include "factorial/factorial.h"

TEST(factorialTest, test2) {
    //arrange
    //act
    //assert
    EXPECT_EQ (Factorial::factorial (0),  1);
    EXPECT_EQ (Factorial::factorial (5), 120);
    EXPECT_EQ (Factorial::factorial (10), 3628800);
}

