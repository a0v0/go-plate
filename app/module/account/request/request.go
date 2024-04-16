package request

import (
	"time"
)

// AccountUpdateRequest is the request body for updating an account
type AccountUpdateRequest struct {
	DisplayName string `json:"display_name"` // The display name associated with this account.
	Username    string `json:"username"`     // The username associated with this account.
	Bio         string `json:"bio"`          // The bio associated with this account.
	Location    string `json:"location"`     // The location set by the user
	Timezone    string `josn:"timezone"`     // The timezone of the user
	LangTag     string `json:"lang_tag"`     // The language set by the user

}

// GetSelfAccountResponse is the response body for getting an account
type GetSelfAccountResponse struct {
	// The ID of the account
	ID string `json:"id"`
	// The ID of the user
	UserID string `json:"user_id"`
	// The profile of the account
	Profile Profile `json:"profile"`
}

type Profile struct {
	// The Profile ID.
	ID string `json:"id"`
	// The Usernames associated with this account.
	Username string `json:"username"`
	// The Display Names associated with this account.
	DisplayName string `json:"display_name"`
	// The Bio associated with this account.
	Bio string `json:"bio"`
	// Language associated with this account. This is a BCP-47 language tag.
	LangTag string `json:"lang_tag"`
	// Location associated with this account.
	Location string `json:"location"`
	// Timezone associated with this account.
	Timezone string `json:"timezone"`
	// Time when the profile was last updated.
	UpdateTime time.Time `json:"update_time"`
	// Time when the account was created.
	CreateTime time.Time `json:"create_time"`
}
