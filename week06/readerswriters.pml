#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte roomEmpty = 1
byte readers = 0
byte mutex = 1

active proctype Writer() {
    do
    ::
        wait(roomEmpty)
        printf("Writer writing! (1)\n")
        printf("Writer writing! (2)\n")
        printf("Writer writing! (3)\n")
        signal(roomEmpty)
    od
}

active[3] proctype Reader() {
    do
    ::
        wait(mutex)
        readers++
        if
        :: readers == 1 -> wait(roomEmpty)
        :: else -> skip
        fi
        signal(mutex)

        printf("Student%d 1\n", _pid)
        printf("Student%d 2\n", _pid)
        printf("Student%d 3\n", _pid)

        wait(mutex)
        readers--
        if
        :: readers == 0 -> signal(roomEmpty)
        :: else -> skip
        fi
        signal(mutex)
    od
}
