package models

import (
	"fmt"

	"github.com/jinzhu/gorm"

	u "github.com/d3z41k/rest-boilerplate/utils"
)

// Contact is a struct to rep user contact
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
}

// Validate incoming contact details
func (c *Contact) Validate() (map[string]interface{}, bool) {
	if c.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}
	if c.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}
	if c.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

// Create new user contact
func (c *Contact) Create() map[string]interface{} {
	if resp, ok := c.Validate(); !ok {
		return resp
	}

	GetDB().Create(c)

	resp := u.Message(true, "success")
	resp["contact"] = c
	return resp
}

// GetContact return a user contact by ID
func GetContact(contactID int, userID uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ? AND user_id = ?", contactID, userID).First(contact).Error
	if err != nil {
		return nil
	}

	return contact
}

// GetContacts return user contacts by user ID
func GetContacts(userID uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", userID).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}

// UpdateContact update a user contact by ID
func UpdateContact(contactData *Contact, contactID int, userID uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ? AND user_id = ?", contactID, userID).First(contact).Error
	if err != nil {
		return nil
	}

	if contactData.Name != "" {
		contact.Name = contactData.Name
	}
	if contactData.Phone != "" {
		contact.Phone = contactData.Phone
	}

	db.Save(contact)

	return contact
}

// DeleteContact delete user contact by ID
func DeleteContact(contactID int, userID uint) map[string]interface{} {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ? AND user_id = ?", contactID, userID).First(contact).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	db.Delete(&contact)

	resp := u.Message(true, "success")
	return resp
}
