package actor

import "github.com/goodmustache/pt/config"

//counterfeiter:generate . Config

type Config interface {
	AddUser(user config.User) error
}
