package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type TrainResource struct {
	ID              int    `json:"id"`
	DriverName      string `json:"driver_name"`
	OperatingStatus bool   `json:"operating_status"`
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetTrain() revel.Result {
	var train TrainResource

	id := c.Params.Route.Get("train-id")

	train.ID, _ = strconv.Atoi(id)
	train.DriverName = "Logan"
	train.OperatingStatus = true

	c.Response.Status = http.StatusOK

	return c.RenderJSON(train)
}

func (c App) CreateTrain() revel.Result {
	var train TrainResource
	c.Params.BindJSON(&train)
	train.ID = 2
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(train)
}

func (c App) RemoveTrain() revel.Result {
	id := c.Params.Route.Get("train-id")
	log.Println("Successfully deleted the resource:", id)
	c.Response.Status = http.StatusOK

	return c.RenderText("")
}
