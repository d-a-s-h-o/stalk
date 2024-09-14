package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Port        int               `yaml:"port"`
	AltPort     int               `yaml:"alt_port"`
	ProfilePort int               `yaml:"profile_port"`
	Scrollback  int               `yaml:"scrollback"`
	DataDir     string            `yaml:"data_dir"`
	KeyFile     string            `yaml:"key_file"`
	Admins      map[string]string `yaml:"admins"`
	Censor      bool              `yaml:"censor,omitempty"`
	Private     bool              `yaml:"private,omitempty"`
	Allowlist   map[string]string `yaml:"allowlist,omitempty"`

	IntegrationConfig string `yaml:"integration_config"`
}

// IntegrationsType stores information needed by integrations.
// Code that uses this should check if fields are nil.
type IntegrationsType struct {
	RPC *RPCInfo `yaml:"rpc"`
}

type RPCInfo struct {
	Port int    `yaml:"port"`
	Key  string `yaml:"key"`
}

var (
	Config = ConfigType{ // first stores default config
		Port:        2222,
		AltPort:     3333,
		ProfilePort: 5555,
		Scrollback:  25,
		DataDir:     "talk-data",
		KeyFile:     "talk-sshkey",

		IntegrationConfig: "",
	}

	Integrations = IntegrationsType{} // all nil

	Log *log.Logger
)

func init() {
	cfgFile := os.Getenv("TALK_CONFIG")
	if cfgFile == "" {
		cfgFile = "talk.yml"
	}

	errCheck := func(err error) {
		if err != nil {
			fmt.Println("err: " + err.Error())
			os.Exit(0) // match `return` behavior
		}
	}

	var d []byte
	if _, err := os.Stat(cfgFile); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Config file not found, so using the default one and writing it to " + cfgFile)

			d, err = yaml.Marshal(Config)
			errCheck(err)
			err = os.WriteFile(cfgFile, d, 0644)
		}
		errCheck(err)
	} else {
		d, err = os.ReadFile(cfgFile)
		errCheck(err)
		err = yaml.UnmarshalStrict(d, &Config)
		errCheck(err)
		fmt.Println("Config loaded from " + cfgFile)
	}

	err := os.MkdirAll(Config.DataDir, 0755)
	errCheck(err)

	logfile, err := os.OpenFile(Config.DataDir+string(os.PathSeparator)+"log.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	errCheck(err)
	Log = log.New(io.MultiWriter(logfile, os.Stdout), "", log.Ldate|log.Ltime|log.Lshortfile)

	if os.Getenv("PORT") != "" {
		Config.Port, err = strconv.Atoi(os.Getenv("PORT"))
		errCheck(err)
	}

	Backlog = make([]backlogMessage, Config.Scrollback)

	if Config.IntegrationConfig != "" {
		d, err = os.ReadFile(Config.IntegrationConfig)
		errCheck(err)
		err = yaml.UnmarshalStrict(d, &Integrations)
		errCheck(err)

		fmt.Println("Integration config loaded from " + Config.IntegrationConfig)

		// Check for individual offline integrations
		if os.Getenv("TALK_OFFLINE_RPC") != "" {
			fmt.Println("Disabling RPC")
			Integrations.RPC = nil
		}
		// Check for global offline for backwards compatibility
		if os.Getenv("TALK_OFFLINE") != "" {
			fmt.Println("Offline mode")
			Integrations.RPC = nil
		}
	}
	rpcInit()
}
