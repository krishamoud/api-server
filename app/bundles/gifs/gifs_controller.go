package gifs

import (
	"github.com/krishamoud/game-server/app/common/controller"
	"net/http"
)

// UsersController struct
type GifsController struct {
	common.Controller
}

// Create a gif
func (c *GifsController) Create(w http.ResponseWriter, r *http.Request) {
	gif := NewGif(r)
	c.SendJSON(
		w,
		r,
		gif,
		http.StatusOK,
	)
}
