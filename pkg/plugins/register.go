package plugins

import (
	"github.com/HhhuYu/schedule-framework/pkg/plugins/sampleplugin"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

// Register custmor plugins register
func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(sampleplugin.Name, sampleplugin.New),
	)
}
