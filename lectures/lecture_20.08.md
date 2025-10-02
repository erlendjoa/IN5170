# Lecture 20.08.2025

Sequential program vs. Parallel/concurrent program (shared memory)
- how do communication between threads work with a shared memory synchronized?
- communications with messages between processes


#### IMPORTANT PROBLEMS/TOPICS INCLUDE:
• Atomic operations
• Interference
• Deadlock, livelock, liveness, fairness
• Locks, critical sections and (active) waiting
• Semaphores and passive waiting
• Monitors
• Java: threads and synchronization

<hr>
<br>

## Shared Memory

Parallel Operator
1. Cooperation to obtain a result
2. Competition for common resources


***Conditional synchronization*** must wait for a specific condition to be satisified before execution can continue.

***Atomic operations*** can not be subdivided into smaller operations. (= cannot be atomic?).

***Program order*** is in place where one sequence happens before the other.

***Interleaving*** of two sequences A,B is a sequence C where all elements of A,B are in C and the order of elements are respected.

***Execution-space Explotion*** factorial growth of atomic operations. Different interleaving can lead to different final states. (Factorial Explosion) (n x m)! / m!^n is number of executions.

***At-most-once property*** is atomic based on a set of conditions.


> state of a program consists of values of the variables at a given moment in the execution.
> 

***Properties*** is a predicate over a program (mathematical formula), resp. its execution and reachable states
- Safety: no undesireable state, 
- Liveness: program will reach a desirable state
- Termination: all execution are finate
- Partial correctness
- Total correctness

Invariant is a property of program states
- Strong invariant: holds for all reacable states
- Weak invariant: holds for all states where atomic block starts or ends
- Loop invariant: holds at the start and end of loop body
- Global invariant: reasons about state of many processes
- Local invariant: reasons about state of one process