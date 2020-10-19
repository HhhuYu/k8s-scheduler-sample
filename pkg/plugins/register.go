package plugins

import (
	"github.com/HhhuYu/schedule-framework/pkg/plugins/demoplugin"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

// Register custmor plugins register
func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(demoplugin.Name, demoplugin.New),
	)
}
