<<<<<<< b15c311f2fa9a3beb220fcf17cddcff370c7abcc
// +build integration

=======
>>>>>>> add integration tests to network metricset
package network

import (
	"testing"

	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

<<<<<<< b15c311f2fa9a3beb220fcf17cddcff370c7abcc
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
>>>>>>> add integration tests to network metricset
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "docker",
		"metricsets": []string{"network"},
		"hosts":      []string{"localhost"},
<<<<<<< b15c311f2fa9a3beb220fcf17cddcff370c7abcc
		"socket":     "unix:///var/run/docker.sock",
=======
>>>>>>> add integration tests to network metricset
	}
}
