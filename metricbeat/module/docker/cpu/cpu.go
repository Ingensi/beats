package cpu

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
	dc"github.com/fsouza/go-dockerclient"

	"github.com/elastic/beats/metricbeat/module/docker"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	if err := mb.Registry.AddMetricSet("docker", "cpu", New); err != nil {
		panic(err)
	}
}

// MetricSet type defines all fields of the MetricSet
// As a minimum it must inherit the mb.BaseMetricSet fields, but can be extended with
// additional entries. These variables can be used to persist data or configuration between
// multiple fetch calls.
type MetricSet struct {
	mb.BaseMetricSet
	cpuService *CPUService
	dockerClient *dc.Client
}

// New create a new instance of the MetricSet
// Part of new is also setting up the configuration by processing additional
// configuration entries if needed.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	//get the configuration
	config := docker.GetDefaultConf()
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}
	return &MetricSet{
		BaseMetricSet: base,
		dockerClient: docker.CreateDockerCLient(config),
		cpuService: NewCpuService(),
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
// It returns the event which is then forward to the output. In case of an error, a
// descriptive error must be returned.
func (m *MetricSet) Fetch() ([]common.MapStr, error) {

	rawStats,err:= docker.FetchDockerStats(m.dockerClient)

	if err == nil {
		formatedStats := m.cpuService.GetCPUstatsList(rawStats)
		return eventsMapping(formatedStats), nil
	}
	return nil, nil


}
