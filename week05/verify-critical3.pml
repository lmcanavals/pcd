bool wantp = false
bool wantq = false

byte count = 0

active proctype P() {
    do
    ::  wantp = true
        !wantq ->
        count++
        assert(count < 2)
        count--
        wantp = false
    od
}

active proctype Q() {
    do
    ::  wantq = true
        !wantp ->
        count++
        assert(count < 2)
        count--
        wantq = false
    od
}
