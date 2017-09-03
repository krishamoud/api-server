// Package deployments handles all requests related to server deployments
package deployments

import (
	"net/http"

	"github.com/krishamoud/game-server/app/common/controller"
)

// Controller struct
type Controller struct {
	common.Controller
}

// Create saves a new user to the database
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var err error
	token := r.FormValue("token")
	maxPlayers := r.FormValue("maxPlayers")
	name := r.FormValue("name")
	service, err := createDSTService(token, maxPlayers, name)
	if c.CheckError(err, http.StatusBadRequest, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		service,
		http.StatusOK,
	)
}

// Test is a testing endpoint
func (c *Controller) Test(w http.ResponseWriter, r *http.Request) {
	panic("Fucked")
}
