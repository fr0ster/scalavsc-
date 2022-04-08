#ifndef TEST_CPP_FACTORIAL_H
#define TEST_CPP_FACTORIAL_H

class Factorial {
public:
    static long long int factorial(long long int n);
protected:
    static long long int _factorial(long long int akk, long long int n);
};

#endif //TEST_CPP_FACTORIAL_H