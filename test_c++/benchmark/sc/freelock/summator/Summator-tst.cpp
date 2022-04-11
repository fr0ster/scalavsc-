#include <benchmark/benchmark.h>
#include "sc/freelock/summator.hpp"
#include <vector>

std::vector<uint8_t> summatorvec1(1, 1);
std::vector<uint8_t> summatorvec10(10, 1);
std::vector<uint8_t> summatorvec100(100, 1);
std::vector<uint8_t> summatorvec1000(1000, 1);
std::vector<uint8_t> summatorvec10000(10000, 1);
std::vector<uint8_t> summatorvec100000(100000, 1);
std::vector<uint8_t> summatorvec1000000(1000000, 1);

static void BM_Summator1(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Summator<uint8_t, uint64_t>::summator(summatorvec1);
}
// Register the function as a benchmark
BENCHMARK(BM_Summator1);

static void BM_Summator10(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Summator<uint8_t, uint64_t>::summator(summatorvec10);
}
// Register the function as a benchmark
BENCHMARK(BM_Summator10);

static void BM_Summator100(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Summator<uint8_t, uint64_t>::summator(summatorvec100);
}
// Register the function as a benchmark
BENCHMARK(BM_Summator100);

static void BM_Summator1000(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Summator<uint8_t, uint64_t>::summator(summatorvec1000);
}
// Register the function as a benchmark
BENCHMARK(BM_Summator1000);

static void BM_Summator10000(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Summator<uint8_t, uint64_t>::summator(summatorvec10000);
}
// Register the function as a benchmark
BENCHMARK(BM_Summator10000);

static void BM_Summator100000(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Summator<uint8_t, uint64_t>::summator(summatorvec100000);
}
// Register the function as a benchmark
BENCHMARK(BM_Summator100000);

static void BM_Summator1000000(benchmark::State &state)
{
    for (auto _ : state)
        SC::Freelock::Summator<uint8_t, uint64_t>::summator(summatorvec1000000);
}
// Register the function as a benchmark
BENCHMARK(BM_Summator1000000);