package cmd

import (
	"strings"
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logLevel string
	psqlDBStr string
)

var ctx = log.WithFields(log.Fields{
	"env": viper.GetString("environment"),
})

var RootCmd = &cobra.Command{
	Use:   "proteus-notify",
	Short: "I tell probes what to do",
	Long: ``,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/proteus-notify.toml)")
	RootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "", "info", "Set the log level")
	RootCmd.PersistentFlags().StringP("db-url", "", "", "Set the url of the postgres database (ex. postgres://username:password@host/dbname?sslmode=verify-full)")
	viper.BindPFlag("database-url", RootCmd.PersistentFlags().Lookup("db-url"))
	viper.SetDefault("active-probes-table", "active_probes")
	viper.SetDefault("probe-updates-table", "probe_updates")
	viper.SetDefault("ios-max-retries", 5)
	viper.SetDefault("android-max-retries", 5)
	viper.SetDefault("environment", "production")
}

func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("proteus-notify")
	viper.AddConfigPath("/etc/proteus/")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	replacer := strings.NewReplacer("-", "_") // Allows us to defined keys with -, but set them in via env variables with _
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()
	if err != nil {
		ctx.WithError(err).Error("failed to read config file")
		panic(err)
	}
	ctx.Infof("using config file:", viper.ConfigFileUsed())

	log.SetHandler(cli.Default)
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		ctx.Error("Invalid log level. Must be one of debug, info, warn, error, fatal")
	}
	log.SetLevel(level)
}
