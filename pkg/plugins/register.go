package plugins

import (
	"github.com/HhhuYu/schedule-framework/pkg/plugins/demoplugin"
	"github.com/HhhuYu/schedule-framework/pkg/plugins/nodelabel"
	"github.com/HhhuYu/schedule-framework/pkg/plugins/scoreplugin"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

// Register custom plugins register
func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(demoplugin.Name, demoplugin.New),
		app.WithPlugin(nodelabel.Name, nodelabel.New),
		app.WithPlugin(scoreplugin.Name, scoreplugin.New),
	)
}
