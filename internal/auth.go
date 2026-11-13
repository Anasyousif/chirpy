package auth

import "github.com/alexedwards/argon2id"

// HashPassword hashes a plain-text password using Argon2id
func HashPassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

// CheckPasswordHash compares a plain-text password with an Argon2id hash
func CheckPasswordHash(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}