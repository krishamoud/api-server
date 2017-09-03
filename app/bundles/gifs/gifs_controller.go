package gifs

import (
	"net/http"

	"github.com/krishamoud/game-server/app/common/controller"
)

// Controller struct
type Controller struct {
	common.Controller
}

// Create a gif
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	gif := NewGif(r)
	c.SendJSON(
		w,
		r,
		gif,
		http.StatusOK,
	)
}
