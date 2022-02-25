package runcmd

import (
	"bytes"
	"context"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Configurable vars
var (
	DisableLogging = false
)

// CmdResult is the command execution result.
type CmdResult struct {
	ExitCode int
	Stdout   string
	Stderr   string
	Error    error
}

// RunCmd runs a command, logs it and returns the result.
func RunCmd(ctx context.Context, dir string, name string, arg ...string) (res CmdResult) {
	res = runCmd(ctx, dir, name, arg...)
	if DisableLogging {
		return
	}
	log.WithFields(log.Fields{
		"command":  strings.Join(append([]string{name}, arg...), " "),
		"dir":      dir,
		"exitCode": res.ExitCode,
		"error":    res.Error,
		"stderr":   res.Stderr,
		"stdout":   res.Stdout,
	}).Info("command execution result")
	return
}

func runCmd(ctx context.Context, dir string, name string, arg ...string) (res CmdResult) {
	cmd := exec.CommandContext(ctx, name, arg...)
	cmd.Dir = dir

	var (
		stdoutBuf bytes.Buffer
		stderrBuf bytes.Buffer
	)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	res.Error = cmd.Run()
	res.ExitCode = cmd.ProcessState.ExitCode()
	res.Stdout = stdoutBuf.String()
	res.Stderr = stderrBuf.String()
	return
}
