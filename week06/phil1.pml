#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte forks[5] = { 1, 1, 1, 1, 1 }

active[5] proctype Philosopher() {
    byte i = _pid
    do
    ::
        printf("%d thinking 1\n", i)
        printf("%d thinking 2\n", i)
        printf("%d thinking 3\n", i)

        wait(forks[i])
        wait(forks[(i+1) % 5])

        printf("%d eat 1\n", i)
        printf("%d eat 2\n", i)
        printf("%d eat 3\n", i)

        signal(forks[i])
        signal(forks[(i+1) % 5])
    od
}

