#include <benchmark/benchmark.h>
#include "fibbonacci/fibbonacci.hpp"
#include <boost/multiprecision/cpp_int.hpp>

using namespace boost::multiprecision;

static void BM_fibN(int n, benchmark::State& state) {
  for (auto _ : state)
    Fibbonacci<int, uint1024_t>::fibbonacci(n);
}

static void BM_fib30(benchmark::State& state) { BM_fibN(30, state); }
// Register the function as a benchmark
BENCHMARK(BM_fib30);

static void BM_fib300(benchmark::State& state) { BM_fibN(300, state); }
// Register the function as a benchmark
BENCHMARK(BM_fib300);

static void BM_fib3000(benchmark::State& state) { BM_fibN(3000, state); }
// Register the function as a benchmark
BENCHMARK(BM_fib3000);

static void BM_fib10000(benchmark::State& state) { BM_fibN(10000, state); }
// Register the function as a benchmark
BENCHMARK(BM_fib10000);

static void BM_fib20000(benchmark::State& state) { BM_fibN(20000, state); }
// Register the function as a benchmark
BENCHMARK(BM_fib20000);

static void BM_fib25000(benchmark::State& state) { BM_fibN(25000, state); }
// Register the function as a benchmark
BENCHMARK(BM_fib25000);

static void BM_fib26000(benchmark::State& state) { BM_fibN(200, state); }
// Register the function as a benchmark
BENCHMARK(BM_fib26000);
