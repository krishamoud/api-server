package users

import (
	"github.com/gorilla/mux"
	"github.com/krishamoud/game-server/app/common/controller"
	"net/http"
)

// UsersController struct
type UsersController struct {
	common.Controller
}

// Index func return all users
func (c *UsersController) Index(w http.ResponseWriter, r *http.Request) {
	var err error
	result, err := indexUsers()
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		result,
		http.StatusOK,
	)
}

// New shows the new user page
func (c *UsersController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
	c.SendJSON(
		w,
		r,
		[]int{0, 1, 2},
		http.StatusOK,
	)
}

// Create saves a new user to the database
func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := createUser(email, password)
	if c.CheckError(err, http.StatusBadRequest, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		user,
		http.StatusOK,
	)
}

// Show a single user
func (c *UsersController) Show(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	userId := vars["userId"]
	ctx := r.Context()
	reqUser := ctx.Value("userId").(string)
	result, err := showUser(userId, reqUser)

	if c.CheckError(err, http.StatusNotFound, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		result,
		http.StatusOK,
	)
}

// Edit page
func (c *UsersController) Edit(w http.ResponseWriter, r *http.Request) {}

// Update user doc and save to database
func (c *UsersController) Update(w http.ResponseWriter, r *http.Request) {}

// Destroy a user
func (c *UsersController) Destroy(w http.ResponseWriter, r *http.Request) {}

// Auth a user
func (c *UsersController) Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := authenticateUser(email, password)
	if c.CheckError(err, http.StatusUnauthorized, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		user,
		http.StatusOK,
	)
}
