# shared variables
communicate betweem process A nd B
# messages
also communicates between A and B


# Shared vs. Distributed memory
- each process has private memory, communication over connections in a "network"

### Message passing

# Synchronous message passing: 
- sender <-> reciever (block after sending)
- no memory buffer
- less concurrency
- more prone to deadlocks

# Asynchronous message passing:
- like sending an email (no blocks)
- memory buffer
- more concurrency
- less prone to deadlocks

- FIFO (queue)
- Preserves message order
- Atomic access
- No packet loss
- Typed

# Filters
- the output is a function of the input (and the initial state)

# Classical monitor
- controlled access to shared resource
- global var safeguard the resource state
- access via procedures
- mutual exclusion