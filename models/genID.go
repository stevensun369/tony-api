package models

import (
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