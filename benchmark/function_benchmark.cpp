#include <benchmark/benchmark.h>

#include "function.h"

static void BM_DemoAdd(benchmark::State& state) {
  for (auto _ : state)
    Add(1111, 1111);
}
// Register the function as a benchmark
BENCHMARK(BM_DemoAdd);

BENCHMARK_MAIN();