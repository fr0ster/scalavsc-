#ifndef TEST_CPP_SUM_H
#define TEST_CPP_SUM_H

namespace SC
{
    namespace Lock
    {
        template <typename A, typename B>
        class Sum
        {
        public:
            static B sum(std::vector<A> xs)
            {
                return _sum(new B(0), xs);
            };

        protected:
            static B _sum(B* akk, std::vector<A> xs)
            {
                if (xs.size() != 0)
                {
                    *akk += xs.back();
                    xs.pop_back();
                    _sum(akk, xs);
                }

                return *akk;
            };
        };
    }
}

#endif // TEST_CPP_SUM_H