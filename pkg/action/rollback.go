package action

import (
	"github.com/helmwave/helmwave/pkg/plan"
	"github.com/urfave/cli/v2"
)

// Rollback is struct for running 'rollback' command.
type Rollback struct {
	plandir string
}

// Run is main function for 'rollback' command.
func (i *Rollback) Run() error {
	p := plan.New(i.plandir)
	if err := p.Import(); err != nil {
		return err
	}

	return p.Rollback()
}

// Cmd returns 'rollback' *cli.Command.
func (i *Rollback) Cmd() *cli.Command {
	return &cli.Command{
		Name:  "rollback",
		Usage: "⏮  Rollback your plan",
		Flags: []cli.Flag{
			flagPlandir(&i.plandir),
		},
		Action: toCtx(i.Run),
	}
}
