package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rleibl/brauspielhaus/config"
	"github.com/rleibl/brauspielhaus/db"
	"net/http"
	"strconv"
)

type Context struct {
	ActivePage string
}

func RunServer() {

	db.Init()

	r := gin.Default()

	c := config.GetConfig()
	r.LoadHTMLGlob(c.TemplatePath)
	r.Static("/static/", c.StaticPath)
	r.GET("/beers/*id", beersHandler)
	r.GET("/", defaultHandler)

	r.Run(c.ServerAdress)
}

func defaultHandler(c *gin.Context) {

	g := gin.H{
		"beers":   db.GetBeers(),
		"context": Context{ActivePage: "home"},
	}

	c.HTML(http.StatusOK, "index.tmpl", g)
}

func beersHandler(c *gin.Context) {

	id := c.Param("id")
	fmt.Println(id)

	// FIXME cleanly separate between 'beers/' and 'beers/[id]'
	//       Use different templates, don't use index.tmpl, which
	//       may change. Create beerlist.tmpl and include in index
	//       and beers/ templates
	if id == "" || id == "/" {
		g := gin.H{
			"beers":   db.GetBeers(),
			"context": Context{ActivePage: "beers"},
		}

		c.HTML(http.StatusOK, "index.tmpl", g)
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "notfound.tmpl", nil)
		return
	}

	b, err := db.GetBeer(i)
	if err != nil {
		c.HTML(http.StatusNotFound, "notfound.tmpl", nil)
		return
	}

	g := gin.H{
		"beer":    b,
		"context": Context{ActivePage: "beers"},
	}
	c.HTML(http.StatusOK, "beer.tmpl", g)
}
