package action

import (
	"github.com/helmwave/helmwave/pkg/plan"
	"github.com/urfave/cli/v2"
)

type Rollback struct {
	plandir string
}

func (i *Rollback) Run() error {
	p := plan.New(i.plandir)
	if err := p.Import(); err != nil {
		return err
	}

	return p.Rollback()
}

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
