package streamdeck

import (
	"flag"
	"fmt"
)

type RegistrationParams struct {
	Port          int
	PluginUUID    string
	RegisterEvent string
	Info          string
}

func ParseRegistrationParams(args []string) (RegistrationParams, error) {
	f := flag.NewFlagSet("registration_params", flag.ContinueOnError)

	port := f.Int("port", -1, "")
	pluginUUID := f.String("pluginUUID", "", "")
	registerEvent := f.String("registerEvent", "", "")
	info := f.String("info", "", "")

	if err := f.Parse(args[1:]); err != nil {
		return RegistrationParams{}, err
	}

	if *port == -1 {
		return RegistrationParams{}, fmt.Errorf("missing -port flag")
	}
	if *pluginUUID == "" {
		return RegistrationParams{}, fmt.Errorf("missing -pluginUUID flag")
	}
	if *registerEvent == "" {
		return RegistrationParams{}, fmt.Errorf("missing -registerEvent flag")
	}
	if *info == "" {
		return RegistrationParams{}, fmt.Errorf("missing -info flag")
	}

	return RegistrationParams{
		Port:          *port,
		PluginUUID:    *pluginUUID,
		RegisterEvent: *registerEvent,
		Info:          *info,
	}, nil
}
