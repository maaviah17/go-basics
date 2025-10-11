## Basics of golang 

# Why is time.Sleep(time.Second * 1) used (18go_routines) ?
it’s needed because: 

-Each go task(i) starts a new goroutine (a lightweight thread).
-But the main goroutine (the main function) doesn’t wait for them to finish.
-If main() ends, the whole program exits — even if goroutines are still running! 😬
-So time.Sleep(...) gives enough time for all the goroutines to complete their fmt.Println() work before the program ends.