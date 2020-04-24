bool wantp = false
bool wantq = false

active proctype P() {
    do
    ::
        printf("P NCS 1\n")
        printf("P NCS 2\n")
        printf("P NCS 3\n")

        !wantq ->
        wantp = true

        printf("P CRITICAL 1\n")
        printf("P CRITICAL 2\n")
        printf("P CRITICAL 3\n")

        wantp = false
    od
}

active proctype Q() {
    do
    ::
        printf("Q NCS 1\n")
        printf("Q NCS 2\n")
        printf("Q NCS 3\n")

        !wantp ->
        wantq = true

        printf("Q CRITICAL 1\n")
        printf("Q CRITICAL 2\n")
        printf("Q CRITICAL 3\n")

        wantq = false
    od
}
