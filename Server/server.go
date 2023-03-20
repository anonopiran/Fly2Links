package server

import (
	config "Fly2Links/Config"
	p2l "Fly2Links/Profile2Link"
	b64 "encoding/base64"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Serve() {
	cfg := config.Config()
	prf := config.Profile()
	route := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET"}
	route.Use(cors.New(corsConfig))
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
		ret := ""
		for _, r_ := range *result {
			ret += r_.Link + "\n"
		}
		switch type_ {
		case "json":
			c.JSON(200, result)
		case "base":
			c.String(200, b64.StdEncoding.EncodeToString([]byte(ret)))
		default:
			c.String(200, ret)
		}
	})
	route.Run()
}
