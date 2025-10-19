# Basics of golang 

## Why is time.Sleep(time.Second * 1) used (chapter18) ?
it’s needed because: 

- Each go task(i) starts a new goroutine (a lightweight thread).
- But the main goroutine (the main function) doesn’t wait for them to finish.
- If main() ends, the whole program exits — even if goroutines are still running! 😬
- So time.Sleep(...) gives enough time for all the goroutines to complete their fmt.Println() work before the program ends.

## Go Routines : 
> these are lightweight concurrent function in Go, allows function to run independently in background w/out blocking the main program.
- uses way less memory and faster than usual threads+these are managed by go runtime not OS
- Concurrency build in go = goRoutines + channels.
- study about fork-join model later (basically the go routines are forked and then joins the main function after the execution of reach routine)

> The main function is also a Goroutine called the main goroutine, if it stops then automatically all the other routines stop.