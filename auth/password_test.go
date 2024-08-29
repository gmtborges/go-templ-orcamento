package auth

import "testing"

func TestAuthService_Generate_Verify_Password(t *testing.T) {
	hash, err := GeneratePasswordHash("passwd123")
	if err != nil {
		t.Errorf("Error generating hash: %v", err)
	}
	isValid, err := VerifyPasswordHash("passwd123", hash)
	if err != nil {
		t.Errorf("Error verifying password: %v", err)
	}
	if !isValid {
		t.Errorf("Expected isValid to be true, got %v", isValid)
	}
}
