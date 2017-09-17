package veneur

type ProxyConfig struct {
	ConsulForwardServiceName string `yaml:"consul_forward_service_name"`
	ConsulRefreshInterval    string `yaml:"consul_refresh_interval"`
	ConsulTraceServiceName   string `yaml:"consul_trace_service_name"`
	SrvForwardServiceName    string `yaml:"srv_forward_service_name"`
	SrvRefreshInterval       string `yaml:"srv_refresh_interval"`
	Debug                    bool   `yaml:"debug"`
	EnableProfiling          bool   `yaml:"enable_profiling"`
	ForwardAddress           string `yaml:"forward_address"`
	HTTPAddress              string `yaml:"http_address"`
	SentryDsn                string `yaml:"sentry_dsn"`
	StatsAddress             string `yaml:"stats_address"`
	TraceAddress             string `yaml:"trace_address"`
	TraceAPIAddress          string `yaml:"trace_api_address"`
}
