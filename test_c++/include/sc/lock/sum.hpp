#ifndef TEST_CPP_SUM_HPP
#define TEST_CPP_SUM_HPP

#include <future>
#include <vector>

namespace SC
{
    namespace Lock
    {
        template <typename A, typename B>
        class Sum
        {
        public:
            static B sum(const std::vector<A> xs)
            {
                return _sum(0, xs, 0);
            };

        protected:
            static B _sum(const B akk, const std::vector<A> &xs, const size_t i)
            {
                if (i == xs.size())
                    return akk;

                return _sum(akk + xs[i], xs, i + 1);
            }
        };
    }
}

#endif // TEST_CPP_SUM_HPP