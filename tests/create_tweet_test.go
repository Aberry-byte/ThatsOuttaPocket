package tests

import (
	"testing"

	"github.com/Aberry-byte/ThatsOuttaPocket/helpers"
)

func TestPost(t *testing.T) {
	helpers.Post("This is a test tweet, please ignore")
}
