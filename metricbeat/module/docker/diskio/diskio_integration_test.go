<<<<<<< 2945c1437b846caac981cc9d67f0b8e40da6c938
// +build integration

=======
>>>>>>> add integration tests to diskio metricset
package diskio

import (
	"testing"

	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

<<<<<<< 2945c1437b846caac981cc9d67f0b8e40da6c938
func TestData(t *testing.T) {
	f := mbtest.NewEventsFetcher(t, getConfig())
	err := mbtest.WriteEvents(f, t)
	if err != nil {
		t.Fatal("write", err)
	}
=======
func TestFetch(t *testing.T) {
	f := mbtest.NewEventsFetcher(t, getConfig())
	event, err := f.Fetch()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(" module : %s metricset : %s event: %+v", f.Module().Name(), f.Name(), event)
>>>>>>> add integration tests to diskio metricset
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "docker",
		"metricsets": []string{"diskio"},
		"hosts":      []string{"localhost"},
<<<<<<< 2945c1437b846caac981cc9d67f0b8e40da6c938
		"socket":     "unix:///var/run/docker.sock",
=======
>>>>>>> add integration tests to diskio metricset
	}
}
