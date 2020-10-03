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
	r.GET("/blog/*id", blogHandler)
	r.GET("/", defaultHandler)
	r.NoRoute(notFoundHandler)

	r.Run(c.ServerAdress)
}

func defaultHandler(c *gin.Context) {

	g := gin.H{
		// XXX return only recent beers here (last 2 or 3)
		"beers":   db.GetBeers(),
		"blogentries": db.GetBlogEntries(),
		"context": Context{ActivePage: "home"},
	}

	c.HTML(http.StatusOK, "index.tmpl", g)
}

func notFoundHandler(c *gin.Context) {

	g := gin.H{
		"context": Context{ActivePage: "None"},
	}

	c.HTML(http.StatusNotFound, "notfound.tmpl", g)
}

func blogHandler(c *gin.Context) {

	id := c.Param("id")

	if id == "" || id == "/" {
		g := gin.H{
			"blogentries":   db.GetBlogEntries(),
			"context": Context{ActivePage: "blog"},
		}

		c.HTML(http.StatusOK, "blog_index.tmpl", g)
		return
	}
	id = id[1:] // strip leading '/'
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Atoi failed for '%s'\n", id)
		c.HTML(http.StatusNotFound, "notfound.tmpl",
			gin.H{"context":Context{ActivePage: "None"}} )
		return
	}

	b, err := db.GetBlogEntry(i)
	if err != nil {
		c.HTML(http.StatusNotFound, "notfound.tmpl",
			gin.H{"context": Context{ActivePage: "None"}} )
		return
	}

	g := gin.H{
		"blog":    b,
		"context": Context{ActivePage: "blog"},
	}
	c.HTML(http.StatusOK, "blog.tmpl", g)
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

		c.HTML(http.StatusOK, "beers_index.tmpl", g)
		return
	}
	id = id[1:] // strip leading '/'
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Atoi failed for '%s'\n", id)
		c.HTML(http.StatusNotFound, "notfound.tmpl",
			gin.H{"context":Context{ActivePage: "None"}} )
		return
	}

	b, err := db.GetBeer(i)
	if err != nil {
		c.HTML(http.StatusNotFound, "notfound.tmpl",
			gin.H{"context":Context{ActivePage: "None"}} )
		return
	}

	g := gin.H{
		"beer":    b,
		"context": Context{ActivePage: "beers"},
	}
	c.HTML(http.StatusOK, "beer.tmpl", g)
}
