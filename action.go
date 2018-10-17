package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"encoding/hex"
)

type Action struct {
	From string
	To string
	Timestamp time.Time
	Amount uint
	Hash string
}

var _action_sha256 = sha256.New()

func init() {
}

func NewAction(from, to string, amount uint) *Action {
	action := &Action{
		From:from, 
		To:to, 
		Timestamp:time.Now(), 
		Amount:amount,
	}
	content := fmt.Sprintf("%s,%s,%d,%d", action.From, action.To, action.Timestamp, action.Amount)
	_action_sha256.Write([]byte(content))
	action.Hash = hex.EncodeToString(_action_sha256.Sum(nil))
	return action
}