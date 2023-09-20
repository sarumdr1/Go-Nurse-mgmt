package middleware

import (
	"fmt"
	"net/http"
	"testing"
)

func TestExtractTokenFromRequest(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	reqToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTUxMjU4NzksImlhdCI6MTY5NTEyMjI3OSwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJ1c2VyQHlvcG1haWwuY29tIn0.Ty5AB6G5buHOqr6IzXujZ4_2KLb7Z1dLPFV3V-6iDEQ"
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", reqToken))
	token := extractTokenFromRequest(req)

	if token != reqToken {
		t.Errorf("Token not matched")
	}
}
