#ifndef TEST_CPP_FIBBONACCI_HPP
#define TEST_CPP_FIBBONACCI_HPP

template <typename A, typename B>
class Fibbonacci
{
public:
    // A tail recursive function to
    // calculate n th fibonacci number
    static B fibbonacci(A n, B a = 0, B b = 1)
    {
        if (n == 0)
            return a;
        if (n == 1)
            return b;
        return fibbonacci(n - 1, b, a + b);
    };
};

#endif // TEST_CPP_FIBBONACCI_HPP