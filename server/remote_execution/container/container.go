package container

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"io"
)

type CommandResult struct {
	Error error
	CommandDebugString string
	Stdout []byte
	Stderr []byte
	ExitCode int
}

type CommandContainer interface {
	Create(ctx context.Context, workingDir string) error
	Exec(ctx context.Context, command *repb.Command, stdin io.Reader, stdout io.Writer) *CommandResult
}
