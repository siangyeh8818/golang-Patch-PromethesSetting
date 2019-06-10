package main

import (
	//	"flag"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	//	"log"

	"os"
	//	"path/filepath"
	yaml "gopkg.in/yaml.v3"
)

var prom_new_job string

func main() {
	Init()
	flag.Parse()
	fin, err := os.Open("prometheus.yml")
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	bytes, err := ioutil.ReadAll(fin)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(bytes))
	yaml_content := string(bytes)

	test := PrometheusConfig{}

	err = yaml.Unmarshal([]byte(yaml_content), &test)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Prometheus_YAML:\n%v\n\n", test)
	fmt.Println(test.ScrapeConfigs[0].Params.Match[0])
	fmt.Println(len(test.ScrapeConfigs[0].Params.Match))
	fmt.Println(test.ScrapeConfigs[0].Params.Match[len(test.ScrapeConfigs[0].Params.Match)-1])

	(&test).AddOpenfaasStruct("{job=\"test12345\"}")
	fmt.Println(test.ScrapeConfigs[0].Params.Match[len(test.ScrapeConfigs[0].Params.Match)-1])

	yamloutput, err := yaml.Marshal(&test)

	WriteWithIoutil("prometheus.yml.new", string(yamloutput))
}

func Init() {
	flag.StringVar(&prom_new_job, "job", "", "job name")
}
