<<<<<<< b418dbcfe557622aa22c477556bb7762e3ece5bf
// +build integration

=======
>>>>>>> add integration tests to cpu metricset
package cpu

import (
	"testing"

	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
)

<<<<<<< b418dbcfe557622aa22c477556bb7762e3ece5bf
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
>>>>>>> add integration tests to cpu metricset
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "docker",
		"metricsets": []string{"cpu"},
		"hosts":      []string{"localhost"},
<<<<<<< b418dbcfe557622aa22c477556bb7762e3ece5bf
		"socket":     "unix:///var/run/docker.sock",
=======
>>>>>>> add integration tests to cpu metricset
	}
}
