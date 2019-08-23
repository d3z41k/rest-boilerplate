package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/d3z41k/rest-boilerplate/models"
	u "github.com/d3z41k/rest-boilerplate/utils"
)

// CreateContact - create user contact
var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserID = userID
	resp := contact.Create()
	u.Respond(w, resp)
}

// GetContacts - get user contacts
var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	data := models.GetContacts(userID)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}