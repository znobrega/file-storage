package fileid

import "github.com/hashicorp/go-uuid"

type UUID struct{}

func  (u UUID) Generate() (string, error) {
	return uuid.GenerateUUID()
}
