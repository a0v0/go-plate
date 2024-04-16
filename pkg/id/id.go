package id

import (
	"github.com/jaevor/go-nanoid"
)

var (
	// AccountIDPrefix is the prefix for Account IDs
	AccountIDPrefix = "acc"
	// MediaIDPrefix is the prefix for Media IDs
	MediaIDPrefix = "med"
	// ProfileIDPrefix is the prefix for Profile IDs
	ProfileIDPrefix = "prf"
)

// NewAccountID returns a new Account ID of type NanoID.
func NewAccountID() string {
	// The canonic NanoID is nanoid.Standard(21).
	canonicID, _ := nanoid.Standard(21)
	return AccountIDPrefix + canonicID()
}

// IsValidAccountID checks if the given ID is a valid Account ID.
func IsValidAccountID(id string) bool {
	return len(id) == 24 && id[:3] == AccountIDPrefix
}

// NewMediaID returns a new Media ID of type NanoID.
func NewMediaID() string {
	// The canonic NanoID is nanoid.Standard(21).
	canonicID, _ := nanoid.Standard(27)
	return MediaIDPrefix + canonicID()
}

// IsValidMediaID checks if the given ID is a valid Media ID.
func IsValidMediaID(id string) bool {
	return len(id) == 30 && id[:3] == MediaIDPrefix
}

// NewProfileID returns a new Profile ID of type NanoID.
func NewProfileID() string {
	// The canonic NanoID is nanoid.Standard(21).
	canonicID, _ := nanoid.Standard(21)
	return ProfileIDPrefix + canonicID()
}

// IsValidProfileID checks if the given ID is a valid Profile ID.
func IsValidProfileID(id string) bool {
	return len(id) == 24 && id[:3] == ProfileIDPrefix
}
