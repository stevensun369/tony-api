package utils

import "backend/env"

func ApplyCashbackRate(value float32) float32 {
  return (env.CashbackRate * value) / 100
}