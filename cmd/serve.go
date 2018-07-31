package cmd

import (
	"github.com/rms1000watt/ioc-mock-server/serve"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the server",
	Example: `./ioc-mock-server serve`,
	Run:     serveFunc,
}

var serveCfg serve.Config

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&serveCfg.Host, "host", "", "Host to listen on")
	serveCmd.Flags().IntVar(&serveCfg.Port, "port", 7100, "Port to listen on")

	serveCmd.Flags().StringVar(&serveCfg.DBType, "db-type", "redis", "Type of DB to use (redis, mock)")
	serveCmd.Flags().StringVar(&serveCfg.RedisAddr, "redis-addr", "redis:6379", "Redis addr to connect to")

	setFlagsFromEnv(serveCmd)
}

func serveFunc(cmd *cobra.Command, args []string) {
	configureLogging()

	serve.Serve(serveCfg)
}
