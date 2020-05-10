package main

import (
	"strings"

	"github.com/gofiber/fiber"
	"github.com/gofiber/limiter"
	"github.com/gofiber/logger"
)

func main() {
	app := fiber.New()

	app.Use(limiter.New())
	app.Use(logger.New())

	app.Post("/", handleForm)
	app.Post("/action", handleForm)

	app.Use("/", func(c *fiber.Ctx) {
		afterError, ok := c.Locals("afterError").(string)
		if ok && afterError != "" {
			c.Redirect(afterError, 303)
		} else {
			c.Status(500)
			c.Write("Désolé, une erreur est survenue.\n")
			c.Write(c.Error())
		}
	})

	app.Listen(8080)
}

func handleForm(c *fiber.Ctx) {
	toMailAddress := c.FormValue("cruciforme-mail")
	mailSubject := c.FormValue("cruciforme-subject")
	afterSuccess := c.FormValue("cruciforme-success")
	afterError := c.FormValue("cruciforme-error")

	c.Locals("afterError", afterError)

	multipartForm, err := c.MultipartForm()
	if err != nil {
		c.Next(err)
		return
	}

	form := &Form{
		ToMailAddress: toMailAddress,
		MailSubject:   mailSubject,
		AfterSuccess:  afterSuccess,
		AfterError:    afterError,
	}

	for key, values := range multipartForm.Value {
		if strings.HasPrefix(key, "cruciforme-") {
			continue
		}

		form.Inputs = append(form.Inputs, &Input{
			Name:  key,
			Value: values[0],
		})
	}

	for key, files := range multipartForm.File {
		form.Files = append(form.Files, &File{
			Name:       key,
			FileHeader: files[0],
		})
	}

	if err = form.sendByMail(); err != nil {
		c.Next(err)
	}

	if form.AfterSuccess != "" {
		c.Redirect(form.AfterSuccess, 303)
	} else {
		c.Send("Formulaire envoyé !")
	}
}
