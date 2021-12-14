// Package flag implements flag-related functionality.
package flag

import (
	"github.com/spf13/cobra"
)

const (
	// AccessTokenName denotes the argument name of the access token flag.
	AccessTokenName = "access-token"

	// VerboseName denotes the name of the verbose flag.
	VerboseName = "verbose"

	// JSONOutputName denotes the name of the json output flag.
	JSONOutputName = "json"

	// LocalOnlyName denotes the name of the local-only flag.
	LocalOnlyName = "local-only"

	// OrgName denotes the name of the org flag.
	OrgName = "org"

	// AppName denotes the name of the application name flag.
	AppNameFlagName = "name"

	// AppName denotes the name of the application name flag.
	RegionName = "region"

	// ImageName denotes the name of the application name flag.
	ImageName = "image"

	// YesName denotes the name of the yes flag.
	YesName = "yes"

	// NowName denotes the name of the now flag.
	NowName = "now"

	// NoDeploy denotes the name of the no deploy flag.
	NoDeployName = "no-deploy"

	// GenerateName denotes the name of the generate name flag.
	GenerateNameFlagName = "generate-name"
)

// Flag wraps the set of flags.
type Flag interface {
	addTo(*cobra.Command)
}

// Add adds flag to cmd, binding them on v should v not be nil.
func Add(cmd *cobra.Command, flags ...Flag) {
	for _, flag := range flags {
		flag.addTo(cmd)
	}
}

// Bool wraps the set of boolean flags.
type Bool struct {
	Name        string
	Shorthand   string
	Description string
	Default     bool
	Hidden      bool
}

func (b Bool) addTo(cmd *cobra.Command) {
	flags := cmd.Flags()

	if b.Shorthand != "" {
		_ = flags.BoolP(b.Name, b.Shorthand, b.Default, b.Description)
	} else {
		_ = flags.Bool(b.Name, b.Default, b.Description)
	}

	f := flags.Lookup(b.Name)
	f.Hidden = b.Hidden
}

// String wraps the set of string flags.
type String struct {
	Name        string
	Shorthand   string
	Description string
	Default     string
	ConfName    string
	EnvName     string
	Hidden      bool
}

func (s String) addTo(cmd *cobra.Command) {
	flags := cmd.Flags()

	if s.Shorthand != "" {
		_ = flags.StringP(s.Name, s.Shorthand, s.Default, s.Description)
	} else {
		_ = flags.String(s.Name, s.Default, s.Description)
	}

	f := flags.Lookup(s.Name)
	f.Hidden = s.Hidden
}

// Int wraps the set of int flags.
type Int struct {
	Name        string
	Shorthand   string
	Description string
	Default     int
	Hidden      bool
}

func (i Int) addTo(cmd *cobra.Command) {
	flags := cmd.Flags()

	if i.Shorthand != "" {
		_ = flags.IntP(i.Name, i.Shorthand, i.Default, i.Description)
	} else {
		_ = flags.Int(i.Name, i.Default, i.Description)
	}

	f := flags.Lookup(i.Name)
	f.Hidden = i.Hidden
}

// StringSlice wraps the set of string slice flags.
type StringSlice struct {
	Name        string
	Shorthand   string
	Description string
	Default     []string
	ConfName    string
	EnvName     string
}

func (ss StringSlice) addTo(cmd *cobra.Command) {
	flags := cmd.Flags()

	if ss.Shorthand != "" {
		_ = flags.StringSliceP(ss.Name, ss.Shorthand, ss.Default, ss.Description)
	} else {
		_ = flags.StringSlice(ss.Name, ss.Default, ss.Description)
	}
}

// Org returns an org string flag.
func Org() String {
	return String{
		Name:        OrgName,
		Description: "The organization to operate on",
	}
}

// Yes returns a yes bool flag.
func Yes() Bool {
	return Bool{
		Name:        YesName,
		Shorthand:   "y",
		Description: "Accept all confirmations",
	}
}

// Name returns an app name string flag.
func AppName() String {
	return String{
		Name:        AppNameFlagName,
		Description: "The name of the application to create",
	}
}

// Region returns an region code string flag.
func Region() String {
	return String{
		Name:        RegionName,
		Description: "The target region for the operation",
	}
}

func Image() String {
	return String{
		Name:        ImageName,
		Description: "The image to deploy",
	}
}

func Now() Bool {
	return Bool{
		Name:        NowName,
		Description: "Deploy now without confirmation",
		Default:     false,
	}
}

func NoDeploy() Bool {
	return Bool{
		Name:        NoDeployName,
		Description: "Do not prompt for deployment",
		Default:     false,
	}
}

func GenerateName() Bool {
	return Bool{
		Name:        GenerateNameFlagName,
		Description: "Always generate a name for the app",
		Default:     false,
	}
}

func RemoteOnly() Bool {
	return Bool{
		Name:        "remote-only",
		Description: "Perform builds on a remote builder instance instead of using the local docker daemon",
		Default:     true,
	}
}
