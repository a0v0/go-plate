package data

import (
	"context"
	"fmt"

	"go_plate/internal/database"
	"go_plate/internal/ent"
	"time"
)

type Data struct {
	DB *database.Database
}

type NewProfileData struct {
	DisplayName string
	Bio         string
	Username    string
	LangTag     string
	Location    string
	Timezone    string
}

func NewData(database *database.Database) *Data {
	return &Data{
		DB: database,
	}
}

// GetAccount returns an account by its user ID.
func (s *Data) GetAccount(uid string) (*ent.Account, error) {
	return s.DB.Ent.Account.
		Query().

		// `Only` fails if no user found,
		// or more than 1 account returned.
		Only(context.Background())
}

// GetProfile returns a user profile by account.
func (s *Data) GetProfile(account *ent.Account) (*ent.Profile, error) {
	return s.DB.Ent.Account.
		QueryProfile(account).Only(context.Background())
}

// GetProfileByUserID returns a user profile by userID.
func (s *Data) GetProfileByUserID(userID string) (*ent.Profile, error) {
	return nil, nil
	// return s.DB.Ent.Profile.Query().Where(profile.HasAccountWith(account.UserID(userID))).Only(context.Background())
}

// UsernameAvailable checks if username is available
func (s *Data) IsUsernameAvailable(username string) bool {
	// _, err := s.DB.Ent.Profile.Query().Where(profile.UsernameEQ(username)).Only(context.Background())
	// return ent.IsNotFound(err)
	return false
}

// CreateAccount creates a new account when a user signs up at superTokens.
func (s *Data) CreateAccount(uid string, createdTime time.Time) error {
	ctx := context.Background()
	tx, err := s.DB.Ent.Tx(ctx)
	if err != nil {
		return fmt.Errorf("error starting a transaction: %w", err)
	}

	// Create Account
	account, err := tx.Account.
		Create().
		SetUserID(uid).
		Save(ctx)
	if err != nil {
		return database.Rollback(tx, fmt.Errorf("failed creating account: %w", err))
	}

	// Create Empty Profile
	_, err = tx.Profile.
		Create().
		SetAccount(account).
		Save(ctx)
	if err != nil {
		return database.Rollback(tx, fmt.Errorf("failed creating user: %w", err))
	}

	// Commit the transaction.
	return tx.Commit()
}

func (s *Data) UpdateProfile(userid string, r *NewProfileData) error {
	ctx := context.Background()
	// Start the transaction
	tx, err := s.DB.Ent.Tx(ctx)
	if err != nil {
		return fmt.Errorf("error starting a transaction: %w", err)
	}

	_, err = tx.Profile.Update().
		SetUsername(r.Username).
		SetDisplayName(r.DisplayName).
		SetBio(r.Bio).
		SetLangTag(r.LangTag).
		SetLocation(r.Location).
		SetTimezone(r.Timezone).
		Save(ctx)
	if err != nil {
		return database.Rollback(tx, fmt.Errorf("failure updating account %w", err))
	}

	// Commit the transaction.
	return tx.Commit()
}
