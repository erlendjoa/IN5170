

## 1


## 3

int amount;
type opKind = enum(DEPOSIT, WITHDRAWL);
chan request(int clientID, amount, kind);
chan reply[n](bool reply);

process SavingAccount {
    queue w_queue
    int clientID; amount;
    int balance = 0;
    opKind kind;

    while (true) {
        recieve request(clientID, amount, kind);
        if (kind == DEPOSIT) {
            balance := balance + amount;
            

        } else { # WITHDRAWL
            if (balance <= amount) {
                insert(w_queue, set(clientID, amount))
            } else {
                balance := balance - amount;
                send reply[clientID](true);
            }
        }
    }

}

process client[i=1 to n] {
    int clientID := i;
    int amount;
    opKind kind;

    while (true) {
        send request(clientID, amount, kind);
        recieve reply[i](reply);
    }
}