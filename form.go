package main

import (
	"mime/multipart"
)

type Form struct {
	ToMailAddress string
	MailSubject   string
	AfterSuccess  string
	AfterError    string
	Inputs        []*Input
	Files         []*File
}

type Input struct {
	Name  string
	Value string
}

type File struct {
	Name       string
	FileHeader *multipart.FileHeader
}
