# Thread Safety & Concurrency

## 1. The Core Concept: Threads
Imagine a kitchen with **one chef**. The chef chops onions, then boils water, then plates the dish. This is **Single-Threaded**. It's safe, but slow if there's a lot of work.

Now imagine a kitchen with **10 chefs** (Threads). They can do multiple things at once. This is **Concurrency**. It's fast, but dangerous.

### The Problem: Race Conditions
If Chef A tries to chop an onion while Chef B is moving the cutting board, someone loses a finger.
In computer terms:
- Thread A tries to read a variable `x`.
- Thread B tries to write `x = 5` at the *exact same nanosecond*.
- Thread A might read garbage data, or the program might crash.

This is a **Race Condition**.

## 2. What is Thread Safety?
Code is **Thread Safe** if it functions correctly when accessed from multiple threads simultaneously.
We achieve this using **Synchronization Primitives** (Locks).

### The Tool: Mutex (Mutual Exclusion)
Think of a Mutex like a **Key to the Toilet**.
1. Chef A wants to use the specific station (Critical Section).
2. Chef A grabs the Key (Lock).
3. Chef A does their work. No one else can enter.
4. Chef A puts the key back (Unlock).
5. Chef B can now grab the Key.

## 3. Real World Application
In web servers, multiple users hit the database or cache at the same time. If your cache isn't thread-safe, your server will randomy crash or return the wrong user's data.

## 4. Common solution in different programming languages:

### Go
```go
mu.Lock()
// do sensitive work
mu.Unlock()
```
