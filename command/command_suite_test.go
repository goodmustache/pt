package command_test

import (
	"testing"

	"github.com/goodmustache/pt/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

func TestCommand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Command Suite")
}

func NewTestUI() (*ui.UI, *Buffer, *Buffer, *Buffer) {
	in := NewBuffer()
	out := NewBuffer()
	err := NewBuffer()
	return ui.NewTestUI(in, out, err), in, out, err
}
