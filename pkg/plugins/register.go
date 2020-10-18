package plugins

import (
	"github.com/HhhuYu/schedule-framework/pkg/plugins/sampleplugins"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

// Register custmor plugins register
func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(sampleplugins.Name, sampleplugins.New),
	)
}
