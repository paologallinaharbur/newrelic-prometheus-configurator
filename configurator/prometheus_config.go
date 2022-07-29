package configurator

import (
	"net/url"
	"time"

	"github.com/alecthomas/units"
)

// PrometheusExtraConfig represents some configuration which will be included in prometheus as it is.
type PrometheusExtraConfig any

// TLSConfig represents tls configuration, `prometheusCommonConfig.TLSConfig` cannot be used directly
// because it does not Marshal to yaml properly.
type TLSConfig struct {
	CAFile             string `yaml:"ca_file,omitempty" json:"ca_file,omitempty"`
	CertFile           string `yaml:"cert_file,omitempty" json:"cert_file,omitempty"`
	KeyFile            string `yaml:"key_file,omitempty" json:"key_file,omitempty"`
	ServerName         string `yaml:"server_name,omitempty" json:"server_name,omitempty"`
	InsecureSkipVerify bool   `yaml:"insecure_skip_verify" json:"insecure_skip_verify"`
	MinVersion         string `yaml:"min_version,omitempty" json:"min_version,omitempty"`
}

// Authorization holds prometheus authorization information.
type Authorization struct {
	Type            string `yaml:"type,omitempty"`
	Credentials     string `yaml:"credentials,omitempty"`
	CredentialsFile string `yaml:"credentials_file,omitempty"`
}

// BasicAuth defines the config for the `Authorization` header on every scrape request.
type BasicAuth struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password,omitempty"`
	PasswordFile string `yaml:"password_file,omitempty"`
}

// OAuth2 defines the config for prometheus to gather a token from the endpoint.
type OAuth2 struct {
	ClientID         string            `yaml:"client_id"`
	ClientSecret     string            `yaml:"client_secret,omitempty"`
	ClientSecretFile string            `yaml:"client_secret_file,omitempty"`
	Scopes           []string          `yaml:"scopes,omitempty"`
	TokenURL         string            `yaml:"token_url"`
	EndpointParams   map[string]string `yaml:"endpoint_params,omitempty"`
	TLSConfig        *TLSConfig        `yaml:"tls_config,omitempty"`
	ProxyURL         string            `yaml:"proxy_url,omitempty"`
}

// Job holds fields which do not change from input and output jobs.
type Job struct {
	JobName               string           `yaml:"job_name"`
	HonorLabels           bool             `yaml:"honor_labels,omitempty"`
	HonorTimestamps       *bool            `yaml:"honor_timestamps,omitempty"`
	Params                url.Values       `yaml:"params,omitempty"`
	Scheme                string           `yaml:"scheme,omitempty"`
	BodySizeLimit         units.Base2Bytes `yaml:"body_size_limit,omitempty"`
	SampleLimit           uint             `yaml:"sample_limit,omitempty"`
	TargetLimit           uint             `yaml:"target_limit,omitempty"`
	LabelLimit            uint             `yaml:"label_limit,omitempty"`
	LabelNameLengthLimit  uint             `yaml:"label_name_length_limit,omitempty"`
	LabelValueLengthLimit uint             `yaml:"label_value_length_limit,omitempty"`
	MetricsPath           string           `yaml:"metrics_path,omitempty"`
	ScrapeInterval        time.Duration    `yaml:"scrape_interval,omitempty"`
	ScrapeTimeout         time.Duration    `yaml:"scrape_timeout,omitempty"`
	TLSConfig             *TLSConfig       `yaml:"tls_config,omitempty"`
	BasicAuth             *BasicAuth       `yaml:"basic_auth,omitempty"`
	Authorization         Authorization    `yaml:"authorization,omitempty"`
	OAuth2                OAuth2           `yaml:"oauth2,omitempty"`
}

// StaticConfig defines each of the static_configs for the prometheus config.
type StaticConfig struct {
	Targets []string          `yaml:"targets"`
	Labels  map[string]string `yaml:"labels,omitempty"`
}

// GlobalConfig configures values that are used across other configuration
// objects.
type GlobalConfig struct {
	// How frequently to scrape targets by default.
	ScrapeInterval time.Duration `yaml:"scrape_interval,omitempty"`
	// The default timeout when scraping targets.
	ScrapeTimeout time.Duration `yaml:"scrape_timeout,omitempty"`
	// The labels to add to any timeseries that this Prometheus instance scrapes.
	ExternalLabels map[string]string `yaml:"external_labels,omitempty"`
}

// RelabelConfig defines relabel config rules which can be used in other configuration objects.
type RelabelConfig struct {
	SourceLabels []string `yaml:"source_labels,omitempty"`
	Separator    string   `yaml:"separator,omitempty"`
	TargetLabel  string   `yaml:"target_label,omitempty"`
	Regex        string   `yaml:"regex,omitempty"`
	Modulus      int      `yaml:"modulus,omitempty"`
	Replacement  string   `yaml:"replacement,omitempty"`
	Action       string   `yaml:"action,omitempty"`
}