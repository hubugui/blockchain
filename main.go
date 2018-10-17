package main

import (
	"fmt"	
)

func main() {
	act := NewAction("hu", "bugui", 100)
	fmt.Printf("action:hash=%s, from=%s, to=%s, amount=%d, timestamp=%s\n",
				act.Hash, act.From, act.To, act.Amount, act.Timestamp)
}