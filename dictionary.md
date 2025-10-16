# IN5170 Dictionary

## Shared Memory

**Sequential program**: one thread of control, full control over the whole memory. <nr>

**Concurrent/Parallel program**: several threads of control, which need to exchange information and coordinate execution. <nr>

**Shared Memory**: a shared memory that multiple threads read from or write to. <nr>

<hr>

**Atomic Operations**: An operation that executes as a single step, appearing to the rest of the system as if it completed instantaneously without interruption.

**Atomic Expression**: An expression whose evaluation and result assignment are performed as a single, uninterrupted unit.

**Mutual Exclusion**: A property that ensures only one thread can access a shared resource or execute a specific segment of code (critical section) at any given time.

**Synchronization**: The process of coordinating the order and timing of execution of concurrent threads to ensure consistent and correct access to shared data.

**Critical Reference**: A portion of the code that accesses shared resources.

**Race Condition**: unpredictable order or timing in which multiple threads access and modify shared data.

**Deadlock**: A condition in which two or more threads are permanently blocked, each waiting for the other to release a resource.


<hr>

**Parallelism**: Parallelism is the technique of running multiple tasks or sub-tasks simultaneously using separate hardware resources (like multiple CPU cores or processors) to significantly reduce the overall computation time.

**Interleaving** The technique used by an operating system to achieve concurrency by switching the CPU's attention between multiple threads, running small segments of each sequentially.

**Interference**: The negative effect that occurs when one thread's actions (such as a partial write) become visible to, and disrupt the correctness of, another thread's operation on shared data.

<hr>

**Factorial Explosion**: (n*m)!/m!^n. 

**Amo-property**: x := e satisfies amo if -> e contains no critical reference or e contains at most one critical reference AND x is not referenced by other processes.

**Disjoint Processes**: V(S1) M V(S2) = Ø

<hr>

**Safety property**: program will never reach an undesirable state.

**Liveness property**: program will eventually reach a desirable state.

**Invariant**: constant, unchanging

<br>

## Lock and Barriers


**Absence of deadlock**: if all processes are trying to enter the CS, *at least one* will succeed.

**Absence of unnecessary delay**: if some processes are trying to enter the CS while other processes are in their non-CS, at least one will succeed.

**Eventual entry**: a process attempting to enter the CS will eventually succeed.

<br>

🚦 *Traffic light analogy for fairness*


**Fairness**: awaiting processes will eventually be run.<br>

**Weakly fair schedule**: atomic condition expression that is continiously often enabled (GREEN-GREEN-GREEN-GREEN-GREEN-...)

**Strong fair schedule**: atomic condition expression that is infinitely often enabled (RED-YELLOW-GREEN-YELLOW-RED-YELLOW-...)



<br>

## Semaphores

**Semaphore**: a special kind of shared program variable, non-negative integer and can only be manipulated by two atomic operations (P and V)

**Binary / General Semaphore**: 0 and 1 or 0 and <=1

P operation:
⟨await (s > 0)s := s − 1⟩

V operation:
⟨s := s + 1⟩

## Actors, Active Objects and Async Communication

**Actor**: concept for distributed concurrency which combined a number of topics like active monitors, objects, race-free.