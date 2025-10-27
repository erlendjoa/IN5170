### Types

## Foundations

# Why Types?: 
Insight into compiler construction. Type systems are insight into how developers analyze.
- static types: compiler detects errors before execution.
- dynamic types: clearer error messages at runtime.
- enforcing certain programming patterns

- modularity by providing interfaces
- hides memory/implementation detail

# What is a Type Expression?
- e has a type T if e always evaluates to a value of type T. {0,1,2,...,n} are types int. 2+2 evaluates to 4, which has type int.
- Strong vs. Weak typing, aim to cover as many possible error sources vs more freedom.
- Type soundness as reachability vs. reduction

# Completeness of Type Systems
- Soundness and completeness share properties with logic.
- Static type systems are typically incomplete while dynamic detect the error too late.

# A Simple Type System
• A type syntax
• A subtyping relation
• A typing environment
• A type judgment
• A set of type rules (the type system itself)
• A notion of type soundness

Expression e with integer (n ∈ Z) and Boolean literals: <br>
e ::= n | true | false | e + e | e ∧ e | e ≤ e

 ((1 + 2) ≤ 3) ∧ true

# Well-typed Syntax

⊢ e : T

This means: e is well-typed with type T

• Judgment is true: ⊢ 1 + 1 : Int <br>
• Judgment is false: ⊢ 1 + 1 : Bool

