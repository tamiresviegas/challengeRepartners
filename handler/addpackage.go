package handler

import (
	"net/http"
	"sort"
	"strconv"
)

func AddPackage(w http.ResponseWriter, r *http.Request) {
	addPackage, err := strconv.Atoi(r.URL.Query().Get("addPackage"))
	if err != nil {
		http.Error(w, "Invalid package size", http.StatusBadRequest)
		return
	}
	if addPackage <= 250 {
		http.Error(w, "Invalid package size, can't add less than 250", http.StatusBadRequest)
		return
	}

	packSizes = append(packSizes, addPackage)
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
