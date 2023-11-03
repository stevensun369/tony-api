package utils

import (
	"backend/env"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GenCode(s int) string {
  if (!env.Dev) {
    rand.Seed(time.Now().UnixNano())
  }
  
  var code string
  for i := 0; i < s; i++ {
    code += strconv.Itoa(rand.Intn(10));
  }

  if (env.Dev) {
    fmt.Println(code)
  }

  return code
}

func GetToday() string {
  t := time.Now()
  return fmt.Sprintf("%v.%v.%v", 
    addLeadingZero(t.Day()), 
    addLeadingZero(int(t.Month())), 
    addLeadingZero(t.Year()),
  )
}

func addLeadingZero(s int) string {
  if s < 10 {
    return fmt.Sprintf("0%v", s)
  } else {
    return fmt.Sprintf("%v", s)
  }
}

func GetNow() string {
  t := time.Now()
  return fmt.Sprintf("%v:%v:%v", 
    addLeadingZero(t.Hour()),
    addLeadingZero(t.Minute()),
    addLeadingZero(t.Second()),
  )
}

func GetLocals(c *fiber.Ctx, name string, result interface{}) {
  json.Unmarshal([]byte(fmt.Sprintf("%v", c.Locals(name))), &result)
} 

func SetLocals(c *fiber.Ctx, name string,  data interface{}) {
	bytes, _ := json.Marshal(data)
	json := string(bytes)
	c.Locals(name, json)
}

func Error(c *fiber.Ctx, err error) error  {
  return c.Status(500).SendString(fmt.Sprintf("%v", err))
} 

func MessageError(c *fiber.Ctx, message string) error {
  return c.Status(401).JSON(map[string]interface{} {
    "message": message,
  })
}