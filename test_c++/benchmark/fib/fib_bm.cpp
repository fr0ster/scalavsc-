#include <benchmark/benchmark.h>
#include "fib/fib.h"

static void BM_fib1(benchmark::State& state) {
  for (auto _ : state)
    Fib::fib(1);
}
// Register the function as a benchmark
BENCHMARK(BM_fib1);

static void BM_fib2(benchmark::State& state) {
  for (auto _ : state)
    Fib::fib(2);
}
// Register the function as a benchmark
BENCHMARK(BM_fib2);

static void BM_fib3(benchmark::State& state) {
  for (auto _ : state)
    Fib::fib(3);
}
// Register the function as a benchmark
BENCHMARK(BM_fib3);

static void BM_fib10(benchmark::State& state) {
  for (auto _ : state)
    Fib::fib(10);
}
// Register the function as a benchmark
BENCHMARK(BM_fib10);

static void BM_fib20(benchmark::State& state) {
  for (auto _ : state)
    Fib::fib(20);
}
// Register the function as a benchmark
BENCHMARK(BM_fib20);

static void BM_fib30(benchmark::State& state) {
  for (auto _ : state)
    Fib::fib(30);
}
// Register the function as a benchmark
BENCHMARK(BM_fib30);

static void BM_fib40(benchmark::State& state) {
  for (auto _ : state)
    Fib::fib(40);
}
// Register the function as a benchmark
BENCHMARK(BM_fib40);