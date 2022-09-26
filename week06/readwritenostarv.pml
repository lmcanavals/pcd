#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte roomEmpty = 1
byte mutex = 1
byte wannaWrite = 1

byte numReaders = 0

active proctype Writer() {
	do
	::
		/* NCS */
		wait(wannaWrite)
		wait(roomEmpty)

		/* taking posession */
		printf("Writer(%d) being in the room 1\n", _pid)
		printf("Writer(%d) being in the room 2\n", _pid)
		printf("Writer(%d) being in the room 3\n", _pid)
		printf("Writer(%d) being in the room 4\n", _pid)

		signal(wannaWrite)
		signal(roomEmpty)
	od
}

active[10] proctype reader() {
	do
	::
		wait(wannaWrite)
		signal(wannaWrite)

		wait(mutex)
		numReaders++
		if
		:: numReaders == 1 -> wait(roomEmpty)
		:: else -> skip
		fi
		signal(mutex)

		/* being in the room */
		printf("Reader(%d) being in the room 1\n", _pid)
		printf("Reader(%d) being in the room 2\n", _pid)
		printf("Reader(%d) being in the room 3\n", _pid)
		printf("Reader(%d) being in the room 4\n", _pid)

		wait(mutex)
		numReaders--
		if
		:: numReaders == 0 -> signal(roomEmpty)
		:: else -> skip
		fi
		signal(mutex)
	od
}
