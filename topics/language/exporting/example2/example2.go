// Sample program to show how the program can't access an unexported identifier from another package.
package main

func main() {
	// NOTE: cannot refer to unexported name counters.alertCounter
	// counter := counters.alertCounter(10)
}
