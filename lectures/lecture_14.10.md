### Actors

Actors work as a buffer/channel for asynchronous message passing.
Fundamental idea is to decouple communication and control.

An actor can react to incoming messages to:
- changes its state
- send infinite numbers of messages to other actors
- create infinite number of new actors

Active Objects: messages between objects with cooperative scheduling