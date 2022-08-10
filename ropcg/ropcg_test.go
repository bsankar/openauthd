package ropcg

import (
	"testing"
)

func TestVerifyCredentials(t *testing.T) {

	err := VerifyCredentials("bsankar", "bsankar")
	if err != nil {
		t.Error("Error verifying credentials")
	}

	err = VerifyCredentials("bsankar", "bsankar1")
	if err == nil {
		t.Error("Error verifying credentials")
	}

	err = VerifyCredentials("bsankar1", "bsankar")
	if err == nil {
		t.Error("Error verifying credentials")
	}

	err = VerifyCredentials("bsankar1", "bsankar1")
	if err != nil {
		t.Error("Error verifying credentials")
	}

}
