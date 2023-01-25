package server

import (
	config "Fly2Links/Config"
	p2l "Fly2Links/Profile2Link"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var cfg *config.SettingsType
var prf *[]config.ProfileType

func Serve() {
	route := gin.Default()
	route.GET(cfg.PathPrefix+"/:id", func(c *gin.Context) {
		var req ProfileRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		result := new([]ProfileResultType)
		addrs_ := c.Query("address")
		remarkPrefix_ := c.Query("remark-prefix")
		filterStr_ := c.Query("filter")
		filter_ := strings.Split(filterStr_, ",")
		type_ := c.Query("as")
		fmt.Printf("hey %s", type_)
		for _, pr := range *prf {
			pr_ := p2l.LinkType{Profile: pr}
			if filterStr_ != "" {
				if !pr_.FilterTag(&filter_) {
					continue
				}
			}
			pr_.SetId(req.Id)
			if addrs_ != "" {
				pr_.SetAddress(addrs_)
			}
			if remarkPrefix_ != "" {
				pr_.SetRemarkPrefix(remarkPrefix_)
			}
			rs := new(ProfileResultType)
			rs.FromLinker(&pr_)
			*result = append(*result, *rs)
		}
		switch type_ {
		case "json":
			c.JSON(200, result)
		default:
			ret := ""
			for _, r_ := range *result {
				ret += r_.Link + "\n"
			}
			c.String(200, ret)
		}
	})
	route.Run()
}
func init() {
	cfg = config.Config()
	prf = config.Profile()
}
