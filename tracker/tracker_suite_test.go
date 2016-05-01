package tracker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tracker Suite")
}
