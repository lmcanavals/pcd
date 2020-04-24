int n = 0

proctype P() {
    byte temp, i
    for (i : 1..10) {
        temp = n
        n = temp + 1
    }
}

init {
    atomic {
        run P()
        run P()
    }

    _nr_pr == 1 -> printf("n = %d\n", n)
    assert(n > 2)
}

/*
Para ejecutar este programa en modo verificador, ejecutar lo siguiente:

spin -a counter.pml
gcc pan.c
./a.out

*/
