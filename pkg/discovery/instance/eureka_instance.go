package instance

import (
	"fmt"
	"github.com/penitence1992/go-gin-server/pkg/discovery/configs"
	"github.com/penitence1992/go-gin-server/pkg/utils/network"
	"math"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func CreateEurekaRegister(config configs.EurekaConfig) (*EurekaInstanceCreate, error) {
	var err error
	port := int(config.Port)
	if port == 0 {
		port = 8080
	}
	appName := config.AppName
	ip := config.IpAddress
	if ip == "" {
		ip = ipSupplier()
	}
	instanceId, err := createInstanceId(ip, appName, port)
	if err != nil {
		return nil, err
	}
	base := ""
	if config.PreferIpAddress {
		base = ip
	} else {
		base = config.Hostname
	}
	return &EurekaInstanceCreate{
		Instance: EurekaInstance{
			InstanceId:       instanceId,
			App:              strings.ToUpper(appName),
			HostName:         config.Hostname,
			IpAddr:           ip,
			Status:           UP,
			OverriddenStatus: UNKNOWN,
			Port: EurekaPort{
				Port:    port,
				Enabled: true,
			},
			SecurePort: EurekaPort{
				Port:    443,
				Enabled: false,
			},
			CountryId: 1,
			DataCenterInfo: EurekaDataCenterInfo{
				Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
				Name:  "MyOwn",
			},
			LeaseInfo: EurekaLeaseInfo{
				RenewalIntervalInSecs: 5,
				DurationInSecs:        20,
				RegistrationTimestamp: FindCurrentTimestampToMillisecond(),
				ServiceUpTimestamp:    FindCurrentTimestampToMillisecond(),
			},
			Metadata:                      config.Metadata,
			HomePageUrl:                   createHomePageUrl(base, port),
			StatusPageUrl:                 createStatusPageUrl(base, port),
			HealthCheckUrl:                createHealthCheckUrl(base, port),
			VipAddress:                    appName,
			SecureVipAddress:              appName,
			IsCoordinatingDiscoveryServer: "false",
			LastUpdatedTimestamp:          FindCurrentTimestampToMillisecond(),
			LastDirtyTimestamp:            FindCurrentTimestampToMillisecond(),
		},
	}, nil
}

func createInstanceId(ip, appName string, port int) (string, error) {
	var (
		err      error
		hostname string
	)
	if hostname, err = os.Hostname(); err == nil {
		return fmt.Sprintf("%s:%s:%d", hostname, ip, port), nil
	} else {
		return fmt.Sprintf("%s:%s:%d", appName, ip, port), nil
	}
}

func FindCurrentTimestampToMillisecond() int64 {
	return time.Now().UnixNano() / int64(math.Pow10(6))
}

func createHomePageUrl(ip string, port int) string {
	return fmt.Sprintf("http://%s:%d/", ip, port)
}

func createStatusPageUrl(ip string, port int) string {
	return fmt.Sprintf("http://%s:%d/actuator/info", ip, port)
}

func createHealthCheckUrl(ip string, port int) string {
	return fmt.Sprintf("http://%s:%d/actuator/health", ip, port)
}

func ipSupplier() string {
	ip, err := network.FindCurrentIp()
	if err != nil {
		log.Warningf("获取当前ip失败:%v, 使用127.0.0.1这个ip", err)
		ip = "127.0.0.1"
	}
	return ip

}
