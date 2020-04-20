package actor

import (
	"github.com/goodmustache/pt/config"
	"github.com/goodmustache/pt/tracker"
)

type User tracker.Me

func (actor *Main) AddUser() (User, error) {
	me, err := actor.API.Me()
	if err != nil {
		return User{}, err
	}

	err = actor.Config.AddUser(config.User{
		APIToken: me.APIToken,
		Email:    me.Email,
		ID:       me.ID,
		Name:     me.Name,
		Username: me.Username,
	})
	return User(me), err
}
