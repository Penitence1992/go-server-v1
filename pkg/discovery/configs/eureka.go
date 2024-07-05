package configs

type EurekaConfig struct {
	ZoneUrl string `json:"zoneUrl"`

	AppName string `json:"appName"`

	IpAddress string `json:"ipAddress"`

	Hostname string `json:"hostname"`

	Port uint `json:"port"`

	PreferIpAddress bool `json:"preferIpAddress"`

	Metadata map[string]string `json:"metadata"`
}
