package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// return c.Render()
	// user := &models.User{Name: "Kiswono Prayogo"}
	// c.Tx.NewRecord(user)
	// c.Tx.Create(user)
	// return c.Render(user)
	var i interface{} = 24
	fmt.Printf("%v\n", i)
	c.Response.Status = 201
	return c.Render()
}
