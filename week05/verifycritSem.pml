#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte s = 1

byte cont = 0

active[2] proctype P() {
    do
    ::
        wait(s)

        cont++
        assert(cont < 2)
        cont--

        signal(s)
    od
}
