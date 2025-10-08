package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password123"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	PrettyPrint(hashed)

	if hashed == "" {
		t.Error("expected non-empty hashed password")
	}

	if hashed == password {
		t.Error("expected hashed password to differ from plain text")
	}

	// Verifica que se pueda validar correctamente
	if err := VerifyHashedPassword(hashed, password); err != nil {
		t.Errorf("expected password to match hash, got error: %v", err)
	}
}

// TestVerifyHashedPassword_Error verifica que falle con contraseñas incorrectas.
func TestVerifyHashedPassword_Error(t *testing.T) {
	password := "password456"
	wrongPassword := "otherpassword"

	hashed, err := HashPassword(password)

	PrettyPrint(hashed)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = VerifyHashedPassword(hashed, wrongPassword)
	if err == nil {
		t.Error("expected error for wrong password, got nil")
	}
}
