#include <iostream>
#include <iomanip>
#include <omp.h>

#define NUM_THREADS 2

using namespace std;

double calcPi(int num_steps=1000000) {
    double step = 1. / (double)num_steps;
    double pi;
    double sum[NUM_THREADS];
    int nthreads;

    omp_set_num_threads(NUM_THREADS);
#pragma omp parallel
    {

    int id = omp_get_thread_num();
    int nthrds = omp_get_num_threads();
    if (id == 0) nthreads = nthrds;

    double x;
    sum[id] = 0.;
    for (int i = id; i < num_steps; i += nthrds) {
        x = (i+0.5)*step;
        sum[id] += 4. / (1 + x*x);
    }

    }

    for (int i = 0; i < nthreads; ++i) {
        pi += sum[i];
    }

    pi = pi * step;
    return pi;
}

int main() {
    double pi;

    double start = omp_get_wtime();

    pi = calcPi();

    double et = omp_get_wtime() - start;

    cout << "Pi: " << setprecision(11) << pi << endl;
    cout << "Time: " << et << " seconds\n";

    return 0;
}

