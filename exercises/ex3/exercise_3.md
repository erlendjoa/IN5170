## 1

## 2

Set pot;
Semaphore collect = 1; # counting semaphore
Semaphore eat = 0; # binary semaphore

Process Bee [i=1 to n]:
    while True:
        P(collect);
        add honey to pot

        if pot is full then
            V(eat)
        else then
            V(collect)

Process Bear:
    while True:
        P(eat)
        eat all honey
        V(collect)


## 4

monitor RW Controller { # Invariant: (nr = 0 ∨ nw = 0) ∧ nw ≤ 1
    int nr := 0, nw := 0 # number of readers, number of writers
    cond oktoread; # signalled when nw = 0
    cond oktowrite; # signalled when nr = 0 and nw = 0
    int write_wish := 0

    procedure request read() {
        while (nw > 0) and (write_wish > 0) wait(oktoread);
        nr := nr + 1;
    }

    procedure release read() {
        nr := nr − 1;
        if nr = 0 signal (oktowrite); # wake up one writer
    }

    procedure request write() {
        while (nr > 0 ∨ nw > 0) {
            write_wish := write_wish + 1
            wait(oktowrite);
            write_wish := write_wish - 1
        }
        nw := nw + 1;
    }


    procedure release write() {
        nw := nw − 1;
        signal(oktowrite); # wake up one writer
        while (nr > 0) signal(okroread)
    }
}


monitor RW Controller { # Invariant: (nr = 0 ∨ nw = 0) ∧ nw ≤ 1
    int nr := 0, nw := 0 # number of readers, number of writers
    cond oktoread; # signalled when nw = 0
    cond oktowrite; # signalled when nr = 0 and nw = 0
    bool turn;

    procedure request read() {
        while (nw > 0) and (turn = True) wait(oktoread);
        nr := nr + 1;
        turn = True
    }

    procedure release read() {
        nr := nr − 1;
        if nr = 0 signal (oktowrite); # wake up one writer
    }

    procedure request write() {
        while (nr > 0 ∨ nw > 0) and (turn = False) wait(oktowrite);
        nw := nw + 1;
        turn = False
    }


    procedure release write() {
        nw := nw − 1;
        signal(oktowrite); # wake up one writer
        while (nr > 0) signal(okroread)
    }
}