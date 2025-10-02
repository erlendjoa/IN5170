#### Semaphores

ensures:
- mutex
- condition sync
Semaphores are:
1. a special kind of shared program variable
2. value is non negative
3. can only be operated by two atomic operations
<hr>

P and V
- P: wait for signal/want to pass
- V: signal an event/release, increment the value
<hr>
<br>
binary vs. counting: true and false vs. using more than one resource

binray: 0 or 1 <br>
counting: s[n] = [1, 1, ..., n] / ([4]1) <br>

P OPERATION: await: <await (s > 0) {s:=s-1}>
<br>
V OPERATION: signal: <s:=s+1>

binary all values are togheter 1 or 0 (!= < 1)


Readers and Writers
(only 1 writer at the time, reader must wait for no writer, writer must wait for no writer or reader)