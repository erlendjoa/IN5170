

#### Monitors

- Implicit mutual exclusion - Two procedures in a monitor is never executed concurrently
- Condition synchronization is expressed by condition variables
- Monitors are implemented by using locks or semaphores

- Process = active -> Monitor = passive/re-active

Monitor based concurrency is about:
1. all shared variables are inside monitor
2. processes communicate by calling monitor procedures
3. processes do not need to know all implementation details

# Condition variables: cond
- each condition is associated with a wait condition
- the value of a condition variable is not accessible and is in a queue of delayed processes.
- cond dv; empty(cv); wait(cv); signal(cv); signal.all(cv);
<br>
- signal(cv): head of queue in cv will wake up (guarantees FIFO principle)

# Bounded Buffer Synchronization
- Buffer of size n
- Producer / Consumer

# Invariants and Signaling
- nr: number of readers / nw: number of writers
- RW: monitor invariant: (nr = 0 or nw = 0) and nw <= 1