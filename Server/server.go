package server

import (
	config "Fly2Links/Config"
	p2l "Fly2Links/Profile2Link"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Serve() {
	cfg := config.Config()
	prf := config.Profile()
	route := gin.Default()
	route.GET(cfg.PathPrefix+"/:zone/:id", func(c *gin.Context) {
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
		for _, pr := range *prf {
			pr_ := p2l.LinkType(pr)
			if filterStr_ != "" {
				if !pr_.FilterTag(&filter_) {
					continue
				}
			}
			if !pr_.FilterZone(&req.Zone) {
				continue
			}
			pr_.Id = req.Id
			if addrs_ != "" {
				pr_.Address = addrs_
			}
			if remarkPrefix_ != "" {
				pr_.SetRemarkPrefix(remarkPrefix_)
			}
			rs := new(ProfileResultType)
			rs.FromLink(pr_)
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
	route.Use(cors.Default())
	route.Run()
}
