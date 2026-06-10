package network_test

import (
	"testing"

	"github.com/labring/aiproxy/core/common/network"
	"github.com/smartystreets/goconvey/convey"
)

func TestIsIpInSubnet(t *testing.T) {
	ip1 := "192.168.0.5"
	ip2 := "125.216.250.89"
	subnet := "192.168.0.0/24"
	convey.Convey("TestIsIpInSubnet", t, func() {
		if ok, err := network.IsIPInSubnet(ip1, subnet); err != nil {
			t.Errorf("failed to check ip in subnet: %s", err)
		} else {
			convey.So(ok, convey.ShouldBeTrue)
		}

		if ok, err := network.IsIPInSubnet(ip2, subnet); err != nil {
			t.Errorf("failed to check ip in subnet: %s", err)
		} else {
			convey.So(ok, convey.ShouldBeFalse)
		}
	})
}
