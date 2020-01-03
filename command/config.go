package command

import "github.com/goodmustache/pt/config"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Config

type Config interface {
	GetUsers() ([]config.User, error)
}
