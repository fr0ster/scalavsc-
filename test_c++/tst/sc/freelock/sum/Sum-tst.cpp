#include "gtest/gtest.h"
#include "sc/freelock/sum.hpp"

TEST(scFreelockSumTest, test4)
{
    // arrange
    // act
    // assert
    auto v1 = std::vector<int>({1, 1, 1, 1, 1});
    auto v5 = std::vector<int>({5, 5, 5, 5, 5});
    auto v10 = std::vector<int>({10, 10, 10, 10, 10});
    auto res1 = SC::Freelock::Sum<int, int>::sum(v1);
    EXPECT_EQ(res1, 5);
    auto res5 = SC::Freelock::Sum<int, int>::sum(v5);
    EXPECT_EQ(res5, 25);
    auto res10 = SC::Freelock::Sum<int, int>::sum(v10);
    EXPECT_EQ(res10, 50);
}
