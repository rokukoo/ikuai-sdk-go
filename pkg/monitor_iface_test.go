package ikuai

import (
	"testing"
)

var (
	TestConfig = &Config{
		Username: "admin",
		Password: "rokuko1124",
		Url:      "http://172.16.0.1",
		Retry: RetryConfig{
			Enable: true,
		},
	}
)

func TestShowMonitorIface(t *testing.T) {
	SetDefaultConfig(TestConfig)

	iface, err := ShowMonitorIface(MonitorIfaceTypeStream)
	if err != nil {
		t.Fatalf("ShowMonitorIface failed: %v", err)
	}
	t.Logf("ShowMonitorIface: %v", iface)
}
