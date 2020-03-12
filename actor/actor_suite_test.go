package actor_test

import (
	"testing"

	. "github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/actor/actorfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestActor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Actor Suite")
}

func NewTestActor() (*Main, *actorfakes.FakeTrackerClient, *actorfakes.FakeConfig) {
	api := new(actorfakes.FakeTrackerClient)
	config := new(actorfakes.FakeConfig)
	return NewActor(api, config), api, config
}
