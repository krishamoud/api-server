// Package deployments handles all requests related to server deployments
package deployments

import (
	"github.com/krishamoud/game-server/app/common/controller"
	"net/http"
)

// DeploymentsController struct
type DeploymentsController struct {
	common.Controller
}

// Create saves a new user to the database
func (c *DeploymentsController) Create(w http.ResponseWriter, r *http.Request) {
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
func (c *DeploymentsController) Test(w http.ResponseWriter, r *http.Request) {
	panic("Fucked")
}
