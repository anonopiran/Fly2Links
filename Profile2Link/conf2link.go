package profile2link

import (
	config "Fly2Links/Config"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func vmessProfileLink(profile *LinkType) (string, error) {
	profileData := map[string]string{
		"add":  profile.Address,
		"port": fmt.Sprint(profile.Port),
		"aid":  fmt.Sprint(profile.AltId),
		"id":   profile.Id,
		"ps":   profile.Remark,
		"scy":  "auto",
		"v":    "2",
		"sni":  profile.Transport.Sni,
		"tls":  profile.Transport.Tls,
		"net":  string(profile.Transport.Net),
	}
	switch profile.Transport.Net {
	case config.Tcp:
		profileData["type"] = profile.Transport.HeaderType
		profileData["host"] = profile.Transport.Host
		profileData["path"] = profile.Transport.Path
	case config.Kcp:
		profileData["type"] = profile.Transport.HeaderType
		profileData["path"] = profile.Transport.Seed
	case config.Ws, config.H2:
		profileData["host"] = profile.Transport.Host
		profileData["path"] = profile.Transport.Path
	case config.Quic:
		profileData["type"] = profile.Transport.HeaderType
		profileData["host"] = profile.Transport.QuicSecurity
		profileData["path"] = profile.Transport.Key
	case config.Grpc:
		profileData["type"] = profile.Transport.Mode
		profileData["path"] = profile.Transport.ServiceName
	}

	jsonData, err := json.Marshal(profileData)
	if err != nil {
		log.Error(err)
		return "", err
	}
	jsonEnc := b64.StdEncoding.EncodeToString(jsonData)
	return "vmess://" + jsonEnc, nil
}
func trojanVlessProfileLink(profile *LinkType) (string, error) {
	link := fmt.Sprintf("%s://%s@%s:%d", profile.Protocol, profile.Id, profile.Address, profile.Port)
	url_, _ := url.Parse(link)
	query := url_.Query()
	query.Add("flow", profile.Flow)
	query.Add("type", string(profile.Transport.Net))
	query.Add("headerType", profile.Transport.HeaderType)
	query.Add("host", profile.Transport.Host)
	query.Add("path", profile.Transport.Path)
	query.Add("seed", profile.Transport.Seed)
	query.Add("quicSecurity", profile.Transport.QuicSecurity)
	query.Add("key", profile.Transport.Key)
	query.Add("mode", profile.Transport.Mode)
	query.Add("serviceName", profile.Transport.ServiceName)
	query.Add("security", profile.Transport.Tls)
	query.Add("sni", profile.Transport.Sni)
	query.Add("fp", string(profile.Transport.FingerPrint))
	if profile.Transport.AllowInsecure {
		query.Add("allowInsecure", "1")
	}
	url_.RawQuery = query.Encode()
	url_.Fragment = profile.Remark
	return url_.String(), nil
}
