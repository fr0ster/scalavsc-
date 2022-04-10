#include "gtest/gtest.h"
#include "sc/freelock/summator.h"

TEST(scFreelockSummatorTest, test4)
{
    // arrange
    // act
    // assert
    auto v1 = std::vector<int>({1, 1, 1, 1, 1});
    auto v5 = std::vector<int>({5, 5, 5, 5, 5});
    auto v10 = std::vector<int>({10, 10, 10, 10, 10});
    auto res1 = SC::Freelock::Summator<int, int>::summator(v1);
    EXPECT_EQ(res1, 5);
    auto res5 = SC::Freelock::Summator<int, int>::summator(v5);
    EXPECT_EQ(res5, 25);
    auto res10 = SC::Freelock::Summator<int, int>::summator(v10);
    EXPECT_EQ(res10, 50);
}
