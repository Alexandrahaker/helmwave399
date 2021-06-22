package action

import (
	"github.com/helmwave/helmwave/pkg/plan"
	"github.com/helmwave/helmwave/tests"
	"os"
	"testing"
)

func TestRenderEnv(t *testing.T) {
	defer clean()

	s := &Yml{
		tests.Root + "01_helmwave.yml.tpl",
		tests.Root + "01_helmwave.yml",
	}

	value := "Test01"
	_ = os.Setenv("PROJECT_NAME", value)
	_ = os.Setenv("NAMESPACE", value)

	err := s.Run()

	if err != nil {
		t.Error(err)
	}

	b, err := plan.NewBody(s.to)
	if err != nil {
		t.Error(err)
	}

	if (value != b.Project) || (value != b.Releases[0].Options.Namespace) {
		t.Error("Failed Test01")
	}

	// Clean
	_ = os.Remove(s.to)
}
