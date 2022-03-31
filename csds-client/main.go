package main

import (
	"envoy-tools/csds-client/client"
	client_v2 "envoy-tools/csds-client/client/v2"
	client_v3 "envoy-tools/csds-client/client/v3"
	"flag"
	"fmt"
	"log"
	"time"
)

// flag vars
var uri string
var authority string
var platform string
var authnMode string
var apiVersion string
var requestFile string
var requestYaml string
var jwt string
var configFile string
var monitorInterval time.Duration
var visualization bool
var filterMode string
var filterPattern string
var tlsCertFilepath string
var tlsCACertsFilepath string
var tlsPrivateKeyFilepath string

// const default values for flag vars
const (
	uriDefault             string        = "trafficdirector.googleapis.com:443"
	authorityDefault       string        = ""
	platformDefault        string        = "gcp"
	authnModeDefault       string        = "auto"
	apiVersionDefault      string        = "v2"
	requestFileDefault     string        = ""
	requestYamlDefault     string        = ""
	jwtDefault             string        = ""
	configFileDefault      string        = ""
	monitorIntervalDefault time.Duration = 0
	visualizationDefault   bool          = false
	filterModeDefault      string        = ""
	filterPatternDefault   string        = ""
	TLSCertFilepath        string        = ""
	TLSPrivateKeyFilepath  string        = ""
	TLSCACertsFilepath     string        = ""
)

// init binds flags with variables
func init() {
	flag.StringVar(&uri, "service_uri", uriDefault, "the uri of the service to connect to")
	flag.StringVar(&authority, "authority", authorityDefault, "the :authority header to use when connecting to uri")
	flag.StringVar(&platform, "platform", platformDefault, fmt.Sprintf("the target platform, one of %v", client.SupportedPlatforms))
	flag.StringVar(&authnMode, "authn_mode", authnModeDefault, "the method to use for authentication (e.g. auto, jwt, ...)")
	flag.StringVar(&apiVersion, "api_version", apiVersionDefault, "which xds api major version to use (e.g. v2, v3, ...)")
	flag.StringVar(&requestFile, "request_file", requestFileDefault, "yaml file that defines the csds request")
	flag.StringVar(&requestYaml, "request_yaml", requestYamlDefault, "yaml string that defines the csds request")
	flag.StringVar(&jwt, "jwt_file", jwtDefault, "path of the -jwt_file")
	flag.StringVar(&configFile, "output_file", configFileDefault, "file name to save configs returned by csds response")
	flag.DurationVar(&monitorInterval, "monitor_interval", monitorIntervalDefault, "the interval of sending request in monitor mode (e.g. 500ms, 2s, 1m ...)")
	flag.BoolVar(&visualization, "visualization", visualizationDefault, "option to visualize the relationship between xDS")
	flag.StringVar(&filterMode, "filter_mode", filterModeDefault, "the filter mode for the filter on xDS nodes to be returned (e.g. prefix, suffix, regex, ...)")
	flag.StringVar(&filterPattern, "filter_pattern", filterPatternDefault, "the filter pattern for the filter on xDS nodes to be returned")
	flag.StringVar(&tlsCertFilepath, "cert", tlsCertFilepath, "filepath to the client TLS certificate")
	flag.StringVar(&tlsPrivateKeyFilepath, "key", tlsPrivateKeyFilepath, "filepath to the client TLS private key")
	flag.StringVar(&tlsCACertsFilepath, "cacert", tlsCACertsFilepath, "filepath to the CAs certs used to verify the server's certificate")
}

func main() {
	flag.Parse()

	clientOpts := client.ClientOptions{
		Uri:                   uri,
		Authority:             authority,
		Platform:              platform,
		AuthnMode:             authnMode,
		RequestFile:           requestFile,
		RequestYaml:           requestYaml,
		Jwt:                   jwt,
		ConfigFile:            configFile,
		MonitorInterval:       monitorInterval,
		Visualization:         visualization,
		FilterMode:            filterMode,
		FilterPattern:         filterPattern,
		TLSCertFilepath:       tlsCertFilepath,
		TLSCACertsFilepath:    tlsCACertsFilepath,
		TLSPrivateKeyFilepath: tlsPrivateKeyFilepath,
	}

	var c client.Client
	var err error
	switch apiVersion {
	case "v2":
		c, err = client_v2.New(clientOpts)
	case "v3":
		c, err = client_v3.New(clientOpts)
	default:
		log.Fatalf("Unsupported xDS API version: %v", apiVersion)
	}

	if err != nil {
		log.Fatal(err)
	}

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
