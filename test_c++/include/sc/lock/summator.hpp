#ifndef TEST_CPP_SUMMATOR_HPP
#define TEST_CPP_SUMMATOR_HPP

namespace SC
{
    namespace Lock
    {
        template <typename A, typename B>
        class Summator
        {
        public:
            static B summator(std::vector<A> xs)
            {
                return _summator(xs);
            };

        protected:
            static B _summator(const std::vector<A> &xs)
            {
                B akk = 0;
                for (size_t i = 0; i < xs.size(); i++)
                {
                    akk += xs[i];
                }

                return akk;
            };
        };
    }
}

#endif // TEST_CPP_SUMMATOR_HPP