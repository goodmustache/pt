package ui

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/landoop/tableprinter"
)

type UI struct {
	STDERR io.Writer
	STDIN  io.Reader
	STDOUT io.Writer

	printer *tableprinter.Printer
	mux     sync.Mutex
}

func NewUI() *UI {
	return &UI{
		STDERR:  os.Stderr,
		STDIN:   os.Stdin,
		STDOUT:  os.Stdout,
		printer: tableprinter.New(os.Stdout),
	}
}

func NewTestUI(in io.Reader, out io.Writer, err io.Writer) *UI {
	return &UI{
		STDERR:  err,
		STDIN:   in,
		STDOUT:  out,
		printer: tableprinter.New(out),
	}
}

func (main *UI) PrintWarning(format string, args ...interface{}) {
	main.mux.Lock()
	defer main.mux.Unlock()

	fmt.Fprintf(main.STDERR, format+"\n", args...)
}

func (main *UI) PrintError(err error) {
	main.mux.Lock()
	defer main.mux.Unlock()

	fmt.Fprintln(main.STDERR, err.Error())
}

func (main *UI) PrintTable(in interface{}) {
	main.mux.Lock()
	defer main.mux.Unlock()

	main.printer.Print(in)
}
