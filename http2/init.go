package http2

import (
	"winput/controller"
	"winput/db"
)

type Http struct{}

var c controller.Ctrl
var d db.MongoDB

var bol controller.ResBool

func init() {
	d = db.MongoDB{"mongodb://localhost:27021", "WServiceInput", "Input"}
	c = controller.Ctrl{}
}
