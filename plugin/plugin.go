package plugin // import "code.cloudfoundry.org/cpu-entitlement-plugin/plugin"

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"

	"code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/cf/trace"
	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cli/plugin/models"
	"code.cloudfoundry.org/cpu-entitlement-plugin/metricfetcher"
	"code.cloudfoundry.org/cpu-entitlement-plugin/token"
	"github.com/fatih/color"
)

type CPUEntitlementPlugin struct{}

func New() *CPUEntitlementPlugin {
	return &CPUEntitlementPlugin{}
}

func (p *CPUEntitlementPlugin) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "CLI-MESSAGE-UNINSTALL" {
		os.Exit(0)
	}

	traceLogger := trace.NewLogger(os.Stdout, true, os.Getenv("CF_TRACE"), "")
	ui := terminal.NewUI(os.Stdin, os.Stdout, terminal.NewTeePrinter(os.Stdout), traceLogger)

	if len(args) != 2 {
		ui.Failed("Usage: `cf cpu-entitlement APP_NAME`")
		os.Exit(1)
	}

	appName := args[1]

	info, err := getCFInfo(cli, appName)
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}

	dopplerURL, err := cli.DopplerEndpoint()
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}

	logCacheURL, err := buildLogCacheURL(dopplerURL)
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}

	tokenGetter := token.NewTokenGetter(cli.AccessToken)
	metricFetcher := metricfetcher.New(logCacheURL, tokenGetter)

	usageMetrics, err := metricFetcher.FetchLatest(info.app.Guid, info.app.InstanceCount)
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}

	ui.Say("Showing CPU usage against entitlement for app %s in org %s / space %s as %s ...\n", terminal.EntityNameColor(appName), terminal.EntityNameColor(info.org), terminal.EntityNameColor(info.space), terminal.EntityNameColor(info.username))

	table := ui.Table([]string{"", bold("usage")})
	for _, usageMetric := range usageMetrics {
		table.Add(fmt.Sprintf("#%d", usageMetric.InstanceId), fmt.Sprintf("%.2f%%", usageMetric.CPUUsage()*100))
	}
	table.Print()
}

type cfInfo struct {
	app      plugin_models.GetAppModel
	username string
	org      string
	space    string
}

func getCFInfo(cli plugin.CliConnection, appName string) (cfInfo, error) {
	app, err := cli.GetApp(appName)
	if err != nil {
		return cfInfo{}, err
	}

	user, err := cli.Username()
	if err != nil {
		return cfInfo{}, err
	}

	org, err := cli.GetCurrentOrg()
	if err != nil {
		return cfInfo{}, err
	}

	space, err := cli.GetCurrentSpace()
	if err != nil {
		return cfInfo{}, err
	}

	return cfInfo{
		app:      app,
		username: user,
		org:      org.Name,
		space:    space.Name,
	}, nil
}

func bold(message string) string {
	return terminal.Colorize(message, color.Bold)
}

func buildLogCacheURL(dopplerURL string) (string, error) {
	logStreamURL, err := url.Parse(dopplerURL)
	if err != nil {
		return "", err
	}

	regex, err := regexp.Compile("doppler(\\S+):443")
	match := regex.FindStringSubmatch(logStreamURL.Host)

	if len(match) != 2 {
		return "", errors.New("Unable to parse log-stream endpoint from doppler URL")
	}

	logStreamURL.Scheme = "http"
	logStreamURL.Host = "log-cache" + match[1]

	return logStreamURL.String(), nil
}
