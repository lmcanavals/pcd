#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

#define N 5

int buffer[N]
byte top
byte notFull = N
byte notEmpty = 0
byte mutex = 1
active proctype Producer() {
    int d
    do
    ::  d++
        wait(notFull)
        wait(mutex)
        buffer[top] = d
        top++
        signal(mutex)
        signal(notEmpty)
    od
}
active[3] proctype Consumer() {
    int d
    do
   ::   wait(notEmpty)
        wait(mutex)
        top--
        d = buffer[top]
        signal(mutex)
        signal(notFull)
        printf("%d consuming %d\n", _pid, d)
    od
}

