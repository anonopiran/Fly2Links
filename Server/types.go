package server

import (
	p2l "Fly2Links/Profile2Link"
	"encoding/json"
)

type ProfileRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}
type ProfileResultType struct {
	Link   string   `json:"link"`
	Remark string   `json:"remark"`
	Tags   []string `json:"tags"`
}

func (v *ProfileResultType) FromLinker(lnk p2l.Linker) {
	v.Link, _ = lnk.GetLink()
	v.Remark = lnk.GetRemark()
	t := lnk.GetTags()
	if t != nil {
		v.Tags = lnk.GetTags()
	} else {
		v.Tags = []string{}
	}

}
func (v *ProfileResultType) AsJson() ([]byte, error) {
	return json.Marshal(v)
}
