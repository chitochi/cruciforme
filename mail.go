package main

import (
	"mime"
	"net/smtp"
	"net/textproto"
	"path/filepath"

	"github.com/jordan-wright/email"
	"github.com/matcornic/hermes/v2"
)

var h = hermes.Hermes{
	Product: hermes.Product{
		Name:      "Cruciforme",
		Link:      "https://crucifor.me",
		Copyright: "Cruciforme, un projet libre et gratuit.",
	},
}

func (form *Form) generateHermesMail() *hermes.Email {
	email := &hermes.Email{
		Body: hermes.Body{
			Title: "Bonjour,",
			Intros: []string{
				"Un de vos formulaires vient d’être envoyé, en voici le contenu.",
			},
			Signature: "Cordialement",
		},
	}

	table := hermes.Table{
		Data: [][]hermes.Entry{},
	}

	for _, input := range form.Inputs {
		table.Data = append(table.Data, []hermes.Entry{
			{Key: "Champ", Value: input.Name},
			{Key: "Valeur", Value: input.Value},
		})
	}

	email.Body.Table = table

	if len(form.Files) == 1 {
		email.Body.Outros = []string{
			"Un fichier a été envoyé avec le formulaire, il est attaché à ce mail.",
		}
	} else if len(form.Files) > 1 {
		email.Body.Outros = []string{
			"Plusieurs fichiers ont été envoyés avec le formulaire, ils sont attachés à ce mail.",
		}
	}

	return email
}

func (form *Form) sendByMail() error {
	hermesMail := form.generateHermesMail()

	htmlBody, err := h.GenerateHTML(*hermesMail)
	if err != nil {
		return err
	}

	textBody, err := h.GeneratePlainText(*hermesMail)
	if err != nil {
		return err
	}

	mail := &email.Email{
		To:      []string{form.ToMailAddress},
		From:    "Cruciforme <no-reply@koicactus.net>",
		Subject: "Formulaire envoyé",
		Text:    []byte(textBody),
		HTML:    []byte(htmlBody),
		Headers: textproto.MIMEHeader{},
	}

	if err = form.attachFiles(mail); err != nil {
		return err
	}

	if err = mail.Send("mail.koicactus.net:587", smtp.PlainAuth("", "no-reply@koicactus.net", "paEez3rDTwyGNwCnjfJQEV3x", "mail.koicactus.net")); err != nil {
		return err
	}

	return nil
}

func (form *Form) attachFiles(mail *email.Email) error {
	for _, formFile := range form.Files {
		file, err := formFile.FileHeader.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		filename := formFile.Name + filepath.Ext(formFile.FileHeader.Filename)
		ct := mime.TypeByExtension(filepath.Ext(filename))

		if _, err = mail.Attach(file, filename, ct); err != nil {
			return err
		}
	}

	return nil
}
