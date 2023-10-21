package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GenCode(s int) string {
  var code string
  for i := 0; i < s; i++ {
    code += strconv.Itoa(rand.Intn(10));
  }
  return code
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