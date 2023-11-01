package models

import (
	"fmt"
	"math/rand"
	"time"
)

var Encoding string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"

func GenID(s int) string {
  rand.Seed(time.Now().UnixNano())
  var ID string
  for i := 0; i < s; i++ {
    ID += string(Encoding[rand.Intn(64)])
  }
  return ID
}

func GenWalletID(s int) string {
  rand.Seed(time.Now().UnixNano())
  var ID string
  for i := 0; i < s; i++ {
    ID += fmt.Sprintf("%v", (rand.Intn(10)))
  }
  return ID
}