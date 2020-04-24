byte turn = 1

active proctype P() {
    do
    ::  
        printf("P NCS 1\n")
        printf("P NCS 2\n")
        printf("P NCS 3\n")

        turn == 1 ->

        printf("P CRITICAL 1\n")
        printf("P CRITICAL 2\n")
        printf("P CRITICAL 3\n")

        turn = 2
    od
}

active proctype Q() {
    do
    ::  
        printf("Q NCS 1\n")
        printf("Q NCS 2\n")
        printf("Q NCS 3\n")

        turn == 2 ->

        printf("Q CRITICAL 1\n")
        printf("Q CRITICAL 2\n")
        printf("Q CRITICAL 3\n")

        turn = 1
    od
}
