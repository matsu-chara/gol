package usecase

import (
	"github.com/matsu-chara/gol/operations"
	"net/http"
)

// Delete remove key
func Delete(filepath string, key string, w http.ResponseWriter, r *http.Request) {
	err := operations.RunRm(filepath, key)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	return
}
