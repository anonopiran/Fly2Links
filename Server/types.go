package server

import (
	p2l "Fly2Links/Profile2Link"
	"encoding/json"
)

type ProfileRequest struct {
	Id   string `uri:"id" binding:"required,uuid"`
	Zone string `uri:"zone"`
}
type ProfileResultType struct {
	Link   string   `json:"link"`
	Remark string   `json:"remark"`
	Tags   []string `json:"tags"`
}

func (v *ProfileResultType) FromLink(lnk p2l.LinkType) {
	v.Link, _ = lnk.GetLink()
	v.Remark = lnk.Remark
	t := lnk.Tags
	if t != nil {
		v.Tags = lnk.Tags
	} else {
		v.Tags = []string{}
	}

}
func (v *ProfileResultType) AsJson() ([]byte, error) {
	return json.Marshal(v)
}
