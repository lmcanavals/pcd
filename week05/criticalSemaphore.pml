#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte s = 1

active[2] proctype P() {
    do
    ::
        printf("%d NCS 1\n", _pid)
        printf("%d NCS 2\n", _pid)
        printf("%d NCS 3\n", _pid)

        wait(s)

        printf("%d CRITICAL 1\n", _pid)
        printf("%d CRITICAL 2\n", _pid)
        printf("%d CRITICAL 3\n", _pid)

        signal(s)
    od
}
