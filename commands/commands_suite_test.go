package commands_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/ghttp"

	"testing"
)

func TestCommands(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands Suite")
}

var (
	commandPath string
	homeDir     string
	server      *Server
)

var _ = SynchronizedBeforeSuite(func() []byte {
	path, err := Build("github.com/goodmustache/pt")
	Expect(err).NotTo(HaveOccurred())
	return []byte(path)
}, func(data []byte) {
	commandPath = string(data)
})

var _ = SynchronizedAfterSuite(func() {}, func() {
	CleanupBuildArtifacts()
})

var _ = BeforeEach(func() {
	var err error
	homeDir, err = ioutil.TempDir("", "pt-test")
	Expect(err).NotTo(HaveOccurred())

	if runtime.GOOS == "windows" {
		os.Setenv("USERPROFILE", homeDir)
	} else {
		os.Setenv("HOME", homeDir)
	}

	server = NewServer()

	SetDefaultEventuallyTimeout(3 * time.Second)
})

var _ = AfterEach(func() {
	server.Close()
	os.RemoveAll(homeDir)
})

func runCommand(args ...string) *Session {
	cmdArgs := append([]string{"--override-tracker-url", server.URL()}, args...)
	cmd := exec.Command(commandPath, cmdArgs...)

	session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	<-session.Exited

	return session
}

func runCommandWithInput(args ...string) (*Session, io.WriteCloser) {
	cmdArgs := append([]string{"--override-tracker-url", server.URL()}, args...)
	cmd := exec.Command(commandPath, cmdArgs...)

	stdin, err := cmd.StdinPipe()
	Expect(err).NotTo(HaveOccurred())

	session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session, stdin
}

func inputValue(input string, stdin io.WriteCloser) {
	_, err := fmt.Fprintf(stdin, input+"\n")
	Expect(err).ToNot(HaveOccurred())
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}

	return os.Getenv("HOME")
}
