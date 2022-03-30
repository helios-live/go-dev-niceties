package dev

import (
	"fmt"
	stdlog "log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/ideatocode/go-log"
	er "github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
)

var (
	// ErrInternalServer global error values can be useful when wrapping errors or inspecting error types
	ErrInternalServer = er.New("error internal server")
)

const timeFormat = "2006-01-02 15:04:05"

// Setup sets the logging end error defaults
func Setup() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zl.Logger = zl.Output(log.NewZero(os.Stdout).Pretty())

	stdlog.SetFlags(0)
	stdlog.SetOutput(zl.Logger)
}

// Trace prints a pretty trace of the error
func Trace(err error) {
	uErr := er.Unpack(err)

	out := ""

	for i := len(uErr.ErrChain) - 1; i >= 0; i-- {
		e := uErr.ErrChain[i]
		out += perr(e.Msg, er.Stack{e.Frame}, len(uErr.ErrChain)-i)
	}
	out += perr(uErr.ErrRoot.Msg, uErr.ErrRoot.Stack, 0)

	fmt.Println(out)
}

func perr(msg string, st er.Stack, level int) string {

	out := ""

	if level == 0 {
		out += fmt.Sprintln("   " + color.RedString(msg))
	} else if level == 1 {
		out += fmt.Sprintln(time.Now().Format(timeFormat), color.HiRedString("Error: ")+color.YellowString(msg))
	} else {
		out += fmt.Sprintln("   " + color.HiBlackString(msg))
	}
	for _, s := range st {
		out += fmt.Sprintf("   %s: %s:%s\n", color.CyanString(s.Name), s.File, color.HiBlueString("%d", s.Line))
	}
	out += fmt.Sprintln()
	return out
}
