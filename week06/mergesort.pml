#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte s1 = 0
byte s2 = 0

active proctype Merge() {
    wait(s1)
    wait(s2)
    printf("Merging both sides!\n")
    printf("Merging both sides!\n")
    printf("Merging both sides!\n")
}

active proctype Sort1Half() {
    byte i
    for (i : 1..10) {
        printf("Sorting first half\n")
    }
    signal(s1)
}

active proctype Sort2Half() {
    byte i
    for (i : 1..10) {
        printf("Sorting second half\n")
    }
    signal(s2)
}

