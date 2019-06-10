package main

type PrometheusConfig struct {
	Global struct {
		ScrapeInterval     string `yaml:"scrape_interval"`
		EvaluationInterval string `yaml:"evaluation_interval"`
		ScrapeTimeout      string `yaml:"scrape_timeout"`
		ExternalLabels     struct {
			Server string `yaml:"server"`
		} `yaml:"external_labels"`
	} `yaml:"global"`

	ScrapeConfigs []ScrapeConfigs `yaml:"scrape_configs"`
	/*
		ScrapeConfigs []struct {
			JobName        string `yaml:"job_name"`
			ScrapeInterval string `yaml:"scrape_interval"`
			HonorLabels    bool   `yaml:"honor_labels,omitempty"`
			MetricsPath    string `yaml:"metrics_path,omitempty"`
			Params         struct {
				Match []string `yaml:"match[]"`
			} `yaml:"params,omitempty"`
			StaticConfigs []struct {
				Targets []string `yaml:"targets"`
			} `yaml:"static_configs"`
		} `yaml:"scrape_configs"`

	*/
}
type ScrapeConfigs struct {
	JobName        string `yaml:"job_name"`
	ScrapeInterval string `yaml:"scrape_interval"`
	HonorLabels    bool   `yaml:"honor_labels,omitempty"`
	MetricsPath    string `yaml:"metrics_path,omitempty"`
	/*
		Params struct {
			Match []string `yaml:"match[]"`
		} `yaml:"params,omitempty"`
	*/
	Params        Params `yaml:"params,omitempty"`
	StaticConfigs []struct {
		Targets []string `yaml:"targets"`
	} `yaml:"static_configs"`
}

type Params struct {
	Match []string `yaml:"match[]"`
}

func (s *PrometheusConfig) AddOpenfaasStruct(job string) {

	s.ScrapeConfigs[0].Params.Match = append(s.ScrapeConfigs[0].Params.Match, job)

}
