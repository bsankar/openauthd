package implgrant

import (
	"testing"
)

func TestVerifyCredentials(t *testing.T) {

	tests := []struct {
		loginpagedata LoginPageData
	}{
		{LoginPageData{"bsankar", "bsankar"}},
		{LoginPageData{"bsankar", "bsankar1"}},
		{LoginPageData{"bsankar1", "bsankar"}},
	}
	for _, test := range tests {
		err := VerifyCredentials(test.loginpagedata)
		if err != nil {
			t.Error("Error verifying credentials")
		}
	}

}
