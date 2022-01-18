package action

import (
	"os"

	"github.com/helmwave/helmwave/pkg/plan"
	"github.com/urfave/cli/v2"
)

// DiffLive is struct for running 'diff live' command.
type DiffLive struct {
	diff    *Diff
	plandir string
}

// Run is main function for 'diff live' command.
func (d *DiffLive) Run() error {
	p := plan.New(d.plandir)
	if err := p.Import(); err != nil {
		return err
	}
	if ok := p.IsManifestExist(); !ok {
		return os.ErrNotExist
	}

	p.DiffLive(d.diff.ShowSecret, d.diff.Wide)

	return nil
}

// Cmd returns 'diff live' *cli.Command.
func (d *DiffLive) Cmd() *cli.Command {
	return &cli.Command{
		Name:  "live",
		Usage: "plan 🆚 live",
		Flags: []cli.Flag{
			flagPlandir(&d.plandir),
		},
		Action: toCtx(d.Run),
	}
}
