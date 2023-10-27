package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

var packSizes = []int{5000, 2000, 1000, 500, 250}

func CalculatePacks(w http.ResponseWriter, r *http.Request) {

	orderQuantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	usedPacksCount := CalculatePacksNeeded(orderQuantity, packSizes)
	response := fmt.Sprintf("For %d items, the following packs were used:\n", orderQuantity)
	for packSize, count := range usedPacksCount {
		response += fmt.Sprintf("%d packs of %d\n", count, packSize)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func CalculatePacksNeeded(orderQuantity int, packSizes []int) map[int]int {
	remainingQuantity := orderQuantity
	usedPacksCount := make(map[int]int)
	var myFirstTime bool = true

	for _, packSize := range packSizes {
		if remainingQuantity == 0 {
			break
		}
		packCount := remainingQuantity / packSize
		if packCount > 0 {
			remainingQuantity %= packSize
			if packSize == 250 && (remainingQuantity > 0 && remainingQuantity <= 250) && myFirstTime {
				usedPacksCount[500] = packCount
				remainingQuantity = 0
				myFirstTime = false
			} else {
				usedPacksCount[packSize] = packCount
			}
		}
	}
	if remainingQuantity > 0 {
		usedPacksCount[packSizes[len(packSizes)-1]]++
	}
	return usedPacksCount
}
