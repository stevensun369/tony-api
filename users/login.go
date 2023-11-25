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

func login(r fiber.Router) {
	g := r.Group("/login")

	g.Post("/phone", func(c *fiber.Ctx) error {
		// getting body
		var body map[string]string
		json.Unmarshal(c.Body(), &body)

		phone := body["phone"]

		if phone == env.DemoPhone {
			fmt.Println("demo login")

			// creating the token
			user := models.User{}
			err := user.Get(bson.M{"phone": env.DemoPhone})
			if err != nil {
				return utils.MessageError(c, err.Error())
			}
			token, err := user.GenToken()
			if err != nil {
				return utils.MessageError(c, err.Error())
			}

			return c.JSON(bson.M{"token": token, "user": user})
		}

		// if user with phone does exist
		if (models.UserCheck(bson.M{"phone": phone})) {
			// gen a code
			code := utils.GenCode(4)

			// get it across
			if !env.Dev {
				utils.SendSMS("+4"+phone, code)
			}

			// hash it
			hashedCode, err := bcrypt.GenerateFromPassword([]byte(code), 10)

			// and set it on the db
			db.Set("code:"+phone, string(hashedCode))

			if err != nil {
				return utils.MessageError(c, "Eroare internă :((")
			}

			// returning
			if env.Dev {
				return c.JSON(bson.M{"code": code, "phone": phone})
			} else {
				utils.SendSMS("+4"+phone, code)
			}

			return c.Status(200).Send([]byte(""))
		} else { // if user witn phone exists, err
			return utils.MessageError(c,
				"Un utilizator cu acest număr de telefon nu există.")
		}
	})

	g.Post("/code", func(c *fiber.Ctx) error {
		// getting body
		var body map[string]string
		json.Unmarshal(c.Body(), &body)

		// getting hashed code
		hashedCode, err := db.Get(fmt.Sprintf("code:%v", body["phone"]))

		if err != nil {
			return utils.MessageError(c, "Întoarceți-vă la pagina de conectare.")
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

		// deleting the code from cache
		db.Del(fmt.Sprintf("code:%v", body["phone"]))

		// creating the token
		user := models.User{}
		err = user.Get(bson.M{"phone": body["phone"]})
		if err != nil {
			return utils.MessageError(c, err.Error())
		}
		token, err := user.GenToken()
		if err != nil {
			return utils.MessageError(c, err.Error())
		}

		return c.JSON(bson.M{"token": token, "user": user})
	})
}
