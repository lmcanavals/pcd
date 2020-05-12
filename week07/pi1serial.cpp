#include <iostream>
#include <iomanip>
#include <omp.h>

using namespace std;

double calcPi(int num_steps=1000000) {
    double step = 1. / (double)num_steps;
    double pi;
    double sum = 0.;
    double x;

    for (int i = 0; i < num_steps; ++i) {
        x = (i+0.5)*step;
        sum += 4. / (1 + x*x);
    }

    pi = sum * step;
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

