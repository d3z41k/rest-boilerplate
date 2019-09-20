package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/d3z41k/rest-boilerplate/models"
	u "github.com/d3z41k/rest-boilerplate/utils"
	"github.com/go-chi/chi"
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

// GetContact - get user contact
var GetContact = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	contactID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	data := models.GetContact(contactID, userID)
	if data == nil {
		u.Respond(w, u.Message(false, "Contact is not found"))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
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

// UpdateContact - update user contact
var UpdateContact = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	contactData := &models.Contact{}
	contactID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(contactData)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := models.UpdateContact(contactData, contactID, userID)
	if data == nil {
		u.Respond(w, u.Message(false, "Contact is not found"))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// DeleteContact - delete user contact
var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	contactID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	resp := models.DeleteContact(contactID, userID)
	if resp == nil {
		u.Respond(w, u.Message(false, "Contact is not found"))
		return
	}

	u.Respond(w, resp)
}
