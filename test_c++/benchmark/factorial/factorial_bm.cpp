#include <benchmark/benchmark.h>
#include "factorial/factorial.h"
#include <boost/multiprecision/cpp_int.hpp>

using namespace boost::multiprecision;

static void BM_FactorialN(int n, benchmark::State &state)
{
  for (auto _ : state)
    Factorial<int, uint1024_t>::factorial(n);
}

static void BM_Factorial100(benchmark::State &state) { BM_FactorialN(100, state); }
// Register the function as a benchmark
BENCHMARK(BM_Factorial100);

static void BM_Factorial200(benchmark::State &state) { BM_FactorialN(200, state); }
// Register the function as a benchmark
BENCHMARK(BM_Factorial200);

static void BM_Factorial300(benchmark::State &state) { BM_FactorialN(300, state); }
// Register the function as a benchmark
BENCHMARK(BM_Factorial300);

static void BM_Factorial1000(benchmark::State &state) { BM_FactorialN(1000, state); }
// Register the function as a benchmark
BENCHMARK(BM_Factorial1000);