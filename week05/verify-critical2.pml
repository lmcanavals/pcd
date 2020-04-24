bool wantp = false
bool wantq = false

byte count = 0

active proctype P() {
    do
    ::  !wantq ->
        wantp = true
        count++
        assert(count < 2)
        count--
        wantp = false
    od
}

active proctype Q() {
    do
    ::  !wantp ->
        wantq = true
        count++
        assert(count < 2)
        count--
        wantq = false
    od
}
