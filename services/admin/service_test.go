package admin_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/influxdb-cluster/services/admin"
)

// Ensure service can serve the root index page of the admin.
func TestService_Index(t *testing.T) {
	// Start service on random port.
	s := admin.NewService(admin.Config{BindAddress: "127.0.0.1:8787"})
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Duration(30)*time.Second)
	defer s.Close()

	// Request root index page.
	resp, err := http.Get("http://" + s.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Validate status code and body.
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	} else if _, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fatalf("unable to read body: %s", err)
	}
}
