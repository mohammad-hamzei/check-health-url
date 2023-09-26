package utils_test

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"

	"git.tashilcar.com/core/tashilcar/core/utils"
	errs "git.tashilcar.com/core/tashilcar/core/yerrors"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Age       uint   `json:"age,omitempty"`
}

func TestTags(t *testing.T) {
	u := &User{}
	tags, err := utils.Tags(u, "json")
	if errs.IsNotNil(err) {
		t.Errorf("error while converting")
	}

	conc := strings.Join(tags, "|")
	h := sha256.New()
	hashed := h.Sum([]byte(conc))
	fmt.Println(conc)
	fmt.Printf("%x\n", hashed)
}
