package helmwave

import (
	log "github.com/sirupsen/logrus"
)

const PLANFILE = "planfile"
const MANIFEST = ".manifest/"

func (c *Config) InitPlan() {
	if c.PlanPath[len(c.PlanPath)-1:] != "/" {
		c.PlanPath += "/"
	}
	log.Info("🛠 Your planfile is ", c.PlanPath+PLANFILE)
}
