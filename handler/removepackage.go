package handler

import (
	"net/http"
	"strconv"
)

func RemovePackage(w http.ResponseWriter, r *http.Request) {

	removeSize, err := strconv.Atoi(r.URL.Query().Get("remove"))
	if err != nil {
		http.Error(w, "Invalid package size", http.StatusBadRequest)
		return
	}

	if removeSize != 250 {
		for i, size := range packSizes {
			if size == removeSize {
				packSizes = append(packSizes[:i], packSizes[i+1:]...)
				break
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
