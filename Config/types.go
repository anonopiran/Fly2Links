package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type TransportEnumType string

const (
	Tcp  TransportEnumType = "tcp"
	Kcp  TransportEnumType = "kcp"
	Ws   TransportEnumType = "ws"
	H2   TransportEnumType = "http"
	Quic TransportEnumType = "quic"
	Grpc TransportEnumType = "grpc"
)

type ProtocolEnumType string

const (
	Vmess  ProtocolEnumType = "vmess"
	Vless  ProtocolEnumType = "vless"
	Trojan ProtocolEnumType = "trojan"
)

type FilePathType string
type LogLevelType string

type SettingsType struct {
	ProfilePath FilePathType `env:"PROFILE_PATH" env-required:"true"`
	LogLevel    LogLevelType `env:"LOG_LEVEL" env-default:"warning"`
	PathPrefix  string       `env:"PATH_PREFIX" env-default:""`
}

type TransportType struct {
	Net          TransportEnumType `json:"network"`
	HeaderType   string            `json:"headerType"`
	Host         string            `json:"host"`
	Path         string            `json:"path"`
	Seed         string            `json:"seed"`
	QuicSecurity string            `json:"security"`
	Key          string            `json:"key"`
	Mode         string            `json:"mode"`
	ServiceName  string            `json:"serviceName"`
	Tls          string            `json:"tls"`
	Sni          string            `json:"serverName"`
}
type ProfileType struct {
	Protocol  ProtocolEnumType `json:"protocol"`
	Address   string           `json:"address"`
	Port      int64            `json:"port"`
	Id        string           `json:"-"`
	AltId     int64            `json:"alterId"`
	Flow      string           `json:"flow"`
	Transport TransportType    `json:"transport"`
	Remark    string           `json:"remark"`
	Tags      []string         `json:"tags"`
	Zone      string           `json:"zone"`
}

// .............

func (f *FilePathType) AsString() string {
	return string(*f)
}

func (f *FilePathType) SetValue(s string) error {
	_, err := os.Stat(s)
	if os.IsNotExist(err) {
		log.WithField("path", s).WithError(err)
		return err
	}
	*f = FilePathType(s)
	return nil
}
func (f *LogLevelType) SetValue(s string) error {
	ll, err := log.ParseLevel(s)
	if err != nil {
		log.WithField("value", s).WithError(err).Error("error while parsing config")
		return err
	}
	log.SetLevel(ll)
	*f = LogLevelType(s)
	return nil
}
