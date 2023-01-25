package profile2link

import config "Fly2Links/Config"

type Linker interface {
	GetLink() (string, error)
	GetRemark() string
	GetTags() []string
	SetRemark(string)
	SetRemarkPrefix(string)
	SetId(string)
	SetAddress(string)
	FilterTag(*[]string) bool
}
type LinkType struct {
	Profile config.ProfileType
}

func (v *LinkType) GetLink() (string, error) {
	var result string
	var err error
	switch v.Profile.Protocol {
	case config.Vmess:
		result, err = vmessProfileLink(&v.Profile)
	case config.Vless, config.Trojan:
		result, err = trojanVlessProfileLink(&v.Profile)
	}
	return result, err
}
func (v *LinkType) GetRemark() string {
	return v.Profile.Remark
}
func (v *LinkType) GetTags() []string {
	return v.Profile.Tags
}
func (v *LinkType) SetRemark(rmrk string) {
	v.Profile.Remark = rmrk
}
func (v *LinkType) SetRemarkPrefix(prf string) {
	v.Profile.Remark = prf + v.Profile.Remark
}
func (v *LinkType) SetId(id string) {
	v.Profile.Id = id
}
func (v *LinkType) SetAddress(addrs string) {
	v.Profile.Address = addrs
}
func (v *LinkType) FilterTag(tags *[]string) bool {
	for _, filt := range *tags {
		for _, t_ := range v.Profile.Tags {
			if filt == t_ {
				return true
			}
		}
	}
	return false
}
