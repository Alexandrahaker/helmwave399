package plan

import (
	"github.com/databus23/helm-diff/diff"
	"github.com/databus23/helm-diff/manifest"
	"github.com/helmwave/helmwave/pkg/helper"
	log "github.com/sirupsen/logrus"
	"os"
)

func (p *Plan) Diff(b *Plan, diffWide int) {
	var visited []string
	noChanges := 0

	for _, rel := range p.body.Releases {
		m := rel.Uniq() + ".yml"
		visited = append(visited, string(m))

		oldSpecs := manifest.Parse(b.manifests[m], rel.Namespace)
		newSpecs := manifest.Parse(p.manifests[m], rel.Namespace)

		change := diff.Manifests(oldSpecs, newSpecs, []string{}, true, diffWide, os.Stdout)
		if !change {
			noChanges++
			log.Info("❎  ", rel.Uniq(), " no changes")
		}
	}

	previous := false
	for _, rel := range b.body.Releases {
		m := string(rel.Uniq() + ".yml")
		if !helper.Contains(m, visited) {
			previous = true
			log.Warn(rel.Uniq(), " was found in previous planfile but not affected in new")
		}
	}

	if noChanges == len(p.body.Releases) && !previous {
		log.Info("Plan has no changes")
	}
}

// TODO: Diff with live
