package utils

import "backend/env"

func ApplyCashbackRate(value int) int {
  return (env.CashbackRate * value) / 100
}