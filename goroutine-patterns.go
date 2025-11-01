// 🎨 GOROUTINE PATTERN GENERATOR
package main

import "fmt"

func printPattern(symbol string, count int, done chan bool) {
	for i := 0; i < count; i++ {
		fmt.Print(symbol)
	}
	done <- true
}

func main() {
	done := make(chan bool)
	
	fmt.Println("🎨 GOROUTINE PATTERN GENERATOR\n")
	
	// Launch 4 different goroutines
	count := 30
	go printPattern("🔴", count, done)
	go printPattern("🔵", count, done)
	go printPattern("🟢", count, done)
	go printPattern("🟡", count, done)
	
	// Wait for all 4 to finish
	<-done
	<-done
	<-done
	<-done
	
	fmt.Println("\n\n✅ All goroutines complete!")
	fmt.Println("Run this program multiple times - notice how the pattern changes!")
}
