Exercise 1:

Actor Calculator (state, memory):
    state = DUAL;
    memory = 0;

    while True:
        M <- (from, cmd, args);

        if state = DUAL:
            switch (M.cmd):
                case (ADD, n, m):
                    res = n+m;
                    M.from <- res;
                case (SGNL):
                    state = SGNL;
                else:
                    pass;
        if state = SGNL:
            switch (M.cmd):
                case (STORE, n):
                    memory = n;
                case (INC, n):
                    M.from <- n + memory;
                case (DUAL):
                    state = DUAL;
                else:
                    pass


Exercise 2:
    go program

Exercise 3:

    - To encode multi-read futures with channels we have to implement asynchronous functionality for recieving and sending multiple values from and to a channel for multiple routines. To ensure termination, implement so that all readers have read before termination. sync.Once ensures an async function is run only once.