#include <benchmark/benchmark.h>
#include "sc/freelock/sum.hpp"
#include <vector>

std::vector<uint8_t> vec1(1, 1);
std::vector<uint8_t> vec10(10, 1);
std::vector<uint8_t> vec100(100, 1);
std::vector<uint8_t> vec1000(1000, 1);
std::vector<uint8_t> vec10000(10000, 1);
std::vector<uint8_t> vec100000(100000, 1);
std::vector<uint8_t> vec1000000(1000000, 1);

static void BM_Sum1(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Sum<uint8_t, uint64_t>::sum(vec1);
}
// Register the function as a benchmark
BENCHMARK(BM_Sum1);

static void BM_Sum10(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Sum<uint8_t, uint64_t>::sum(vec10);
}
// Register the function as a benchmark
BENCHMARK(BM_Sum10);

static void BM_Sum100(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Sum<uint8_t, uint64_t>::sum(vec100);
}
// Register the function as a benchmark
BENCHMARK(BM_Sum100);

static void BM_Sum1000(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Sum<uint8_t, uint64_t>::sum(vec1000);
}
// Register the function as a benchmark
BENCHMARK(BM_Sum1000);

static void BM_Sum10000(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Sum<uint8_t, uint64_t>::sum(vec10000);
}
// Register the function as a benchmark
BENCHMARK(BM_Sum10000);

static void BM_Sum100000(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Sum<uint8_t, uint64_t>::sum(vec100000);
}
// Register the function as a benchmark
BENCHMARK(BM_Sum100000);

static void BM_Sum1000000(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Sum<uint8_t, uint64_t>::sum(vec1000000);
}
// Register the function as a benchmark
BENCHMARK(BM_Sum1000000);