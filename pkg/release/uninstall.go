package release

import (
	"errors"
	"github.com/helmwave/helmwave/pkg/feature"
	"github.com/helmwave/helmwave/pkg/helper"
	"github.com/helmwave/helmwave/pkg/parallel"
	log "github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	helm "helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func (rel *Config) Uninstall() (*release.UninstallReleaseResponse, error) {
	cfg, err := helper.ActionCfg(rel.Options.Namespace, helm.New())
	if err != nil {
		return nil, err
	}

	client := action.NewUninstall(cfg)

	return client.Run(rel.Name)

}

func Uninstall(releases []*Config) (err error) {
	if len(releases) == 0 {
		return emptyReleases
	}

	log.Info("🛥 Sync releases")
	var fails []*Config

	if feature.Parallel {
		wg := parallel.NewWaitGroup()
		log.Debug("🐞 Run in parallel mode")
		wg.Add(len(releases))
		for i := range releases {
			go func(wg *parallel.WaitGroup, release *Config, fails *[]*Config) {
				defer wg.Done()
				release.UninstallWithFails(fails)
			}(wg, releases[i], &fails)
		}

		err := wg.Wait()
		if err != nil {
			return err
		}

	} else {
		for _, r := range releases {
			r.UninstallWithFails(&fails)
		}
	}

	n := len(releases)
	k := len(fails)

	log.Infof("Success %d / %d", n-k, n)
	if k > 0 {
		for _, rel := range fails {
			log.Errorf("%q was not deleted from %q", rel.Name, rel.Options.Namespace)
		}

		return errors.New("uninstall failed")
	}
	return nil
}

func (rel *Config) UninstallWithFails(fails *[]*Config) {
	r, err := rel.Uninstall()
	if r != nil {
		log.Info(r.Info)
	}
	if err != nil {
		log.Error("❌ ", err)
		*fails = append(*fails, rel)
	} else {
		log.Infof("✅ %s", rel.UniqName())
	}
}
