package profile2link

import (
	config "Fly2Links/Config"

	"golang.org/x/exp/slices"
)

type LinkType config.ProfileType

func (v *LinkType) GetLink() (string, error) {
	var result string
	var err error
	switch v.Protocol {
	case config.Vmess:
		result, err = vmessProfileLink(v)
	case config.Vless, config.Trojan:
		result, err = trojanVlessProfileLink(v)
	}
	return result, err
}
func (v *LinkType) SetRemarkPrefix(prf string) {
	v.Remark = prf + v.Remark
}
func (v *LinkType) FilterTag(tags *[]string) bool {
	for _, filt := range *tags {
		for _, t_ := range v.Tags {
			if filt == t_ {
				return true
			}
		}
	}
	return false
}
func (v *LinkType) FilterZone(zone *string) bool {
	return slices.Contains(v.Zone, *zone)
}
