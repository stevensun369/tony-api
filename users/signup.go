package users

import (
	"backend/db"
	"backend/env"
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func signup(r fiber.Router) {
  g := r.Group("/signup")

  g.Post("/phone", func (c *fiber.Ctx) error {
    // getting body
    var body map[string]string
    json.Unmarshal(c.Body(), &body)

    phone := body["phone"]

    // if user with phone doesn't exist
    if (!models.UserCheck(bson.M{"phone": phone})) {
      // gen a code
      code := utils.GenCode(4)

      // get it across
      if (!env.Dev) { 
        utils.SendSMS("+4" + phone, code) 
      }

      // hash it
      hashedCode, err := bcrypt.GenerateFromPassword([]byte(code), 10)

      // and set it on the db
      db.Set("code:" + phone, string(hashedCode))

      if err != nil {
        return utils.MessageError(c, "Eroare internă :((")
      }

      // returning
      if (env.Dev) { 
        return c.JSON(bson.M{"code": code, "phone": phone})
      } else { 
        utils.SendSMS("+4" + phone, code) 
      }

      return c.Status(200).Send([]byte(""))
    } else { // if user witn phone exists, err
      return utils.MessageError(c, 
        "Un utilizator cu acest număr de telefon există deja.")
    }
  })

  g.Post("/code", func (c *fiber.Ctx) error {
    // getting body
    var body map[string]string
    json.Unmarshal(c.Body(), &body)

    // getting hashed code
    hashedCode, err := db.Get(fmt.Sprintf("code:%v", body["phone"]))

    if err != nil {
      return utils.MessageError(c, "Eroare interna. (CODE REDIS)")
    }

    // comparing hashed with provided code
    compareErr := bcrypt.CompareHashAndPassword(
      []byte(hashedCode),
      []byte(body["code"]),
    )

    // if the code is wrong, err
    if compareErr != nil {
      return utils.MessageError(c, "Codul este greșit")
    }
    
    // deleting the code from caches
    db.Del(fmt.Sprintf("code:%v", body["phone"]))

    // build & save a user
    user := models.User {}
    user.Create(body["phone"])

    // get a token
    token, err := user.GenToken()
  
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    return c.JSON(bson.M{
      "token": token,
    })
  })

  g.Post("/username", AuthMiddleware, func (c *fiber.Ctx) error {
    ID := fmt.Sprintf("%v", c.Locals("ID"))
    
    // getting body
    var body map[string]string
    json.Unmarshal(c.Body(), &body)
    username := body["username"]

    user := models.User{}

    // adding username
    err := user.AddUsername(ID, username)
    if err != nil {
      return utils.MessageError(c, err.Error())
    }
    
    // generating token
    token, err := user.GenToken()
    if err != nil {
      return utils.MessageError(c, err.Error())
    }

    // returning
    return c.JSON(bson.M{"token": token, "user": user})
  })
}