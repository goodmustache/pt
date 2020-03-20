package command

import "github.com/goodmustache/pt/config"

//counterfeiter:generate . Config

type Config interface {
	GetUsers() ([]config.User, error)
}
