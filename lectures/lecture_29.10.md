# Linearity

- Data types: an abstraction over the contents of memory
- Behavioral types: an abstraction over allowed operations.

- Covariant
- Contravariant
- Invariant

## Linear Types

c -> chan?1!1   # T is Linear
c -> chan?0!0   # T cannot be used anymore
c -> chan?1!0   # T recieving possible but not sending
c -> chan?0!1   # T sending possible but not recieving

Limit capabilities for current thread, catch violation.

## Typing Environment
- Γ can be looked at as goroutines.

## Split the gamma into two goroutines:
- for each x with channel type we have Γ(x) = Γ1(x) + Γ2(x)
- chan?1!1 T = chan?0!1 T + chan?1!0 T

- {n -> Int, c -> chan?0,!1 Int} = {n -> Int, c -> chan?0,!0 Int} + {n -> Int, c -> chan?0,!1 Int}


Γ is unrestricted un(Γ) if all contained chan have n=0 and m=0 

## Usage Types

?.!.0 = recieve, send, no other usage
?.0&!.0 = recieve or send, no other usage
?.0+!.0 = use for synchronization once (one goroutine recieves and one sends)
?.!.0+!.?.0 = synchronize twice (would be a deadlock if both started with same !,! or ?,?)
