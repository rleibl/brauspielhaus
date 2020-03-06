package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rleibl/brauspielhaus/config"
	"github.com/rleibl/brauspielhaus/db"
	"net/http"
	"strconv"
)

func RunServer() {

	config.Init()
	db.Init()

	r := gin.Default()

	c := config.GetConfig()
	r.LoadHTMLGlob(c.TemplatePath)
	r.Static("/static/", c.StaticPath)
	r.GET("/beers/:id", beersHandler)
	r.GET("/", defaultHandler)

	r.Run(":8080")
}

func defaultHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.tmpl", gin.H{"beers": db.GetBeers()})
}

func beersHandler(c *gin.Context) {

	id := c.Param("id")
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
	c.HTML(http.StatusOK, "beer.tmpl", gin.H{"beer": b})
}
