#include <benchmark/benchmark.h>
#include "factorial/factorial.h"

static void BM_Factorial1(benchmark::State& state) {
  for (auto _ : state)
    Factorial::factorial(1);
}
// Register the function as a benchmark
BENCHMARK(BM_Factorial1);

static void BM_Factorial2(benchmark::State& state) {
  for (auto _ : state)
    Factorial::factorial(2);
}
// Register the function as a benchmark
BENCHMARK(BM_Factorial2);

static void BM_Factorial3(benchmark::State& state) {
  for (auto _ : state)
    Factorial::factorial(3);
}
// Register the function as a benchmark
BENCHMARK(BM_Factorial3);

static void BM_Factorial10(benchmark::State& state) {
  for (auto _ : state)
    Factorial::factorial(10);
}
// Register the function as a benchmark
BENCHMARK(BM_Factorial10);

static void BM_Factorial20(benchmark::State& state) {
  for (auto _ : state)
    Factorial::factorial(20);
}
// Register the function as a benchmark
BENCHMARK(BM_Factorial20);

static void BM_Factorial30(benchmark::State& state) {
  for (auto _ : state)
    Factorial::factorial(30);
}
// Register the function as a benchmark
BENCHMARK(BM_Factorial30);

static void BM_Factorial100(benchmark::State& state) {
  for (auto _ : state)
    Factorial::factorial(100);
}
// Register the function as a benchmark
BENCHMARK(BM_Factorial100);