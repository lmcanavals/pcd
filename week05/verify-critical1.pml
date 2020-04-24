byte turn = 1

byte count = 0

active proctype P() {
    do
    ::  turn == 1 ->
        count++
        assert(count < 2)
        count--
        turn = 2
    od
}

active proctype Q() {
    do
    ::  turn == 2 ->
        count++
        assert(count < 2)
        count--
        turn = 1
    od
}
