package services

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

// in-memory store for points.
// would replace with a database in a production app
var pointsStore = map[string]int{
	"abc": 100,
}
var ERR_RECEIPT_NOT_FOUND = fmt.Errorf("No receipt found for that ID.")

func ProcessReceipt(receipt Receipt) (string, error) {
	points, err := calculatePoints(receipt)
	if err != nil {
		return "", err
	}
	id := generateID()
	pointsStore[id] = points
	return id, nil
}

func calculatePoints(receipt Receipt) (int, error) {
	points := 0
	points += countAlphanumeric(receipt.Retailer)
	points += calculatePointsFromTotal(receipt.Total)
	points += calculatePointsFromItems(receipt.Items)
	datePoints, err := calculatePointsFromPurchaseDate(receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	points += datePoints
	timePoints, err := calculatePointsFromPurchaseTime(receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	points += timePoints

	return points, nil
}

func countAlphanumeric(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	return count
}

func calculatePointsFromTotal(total float64) int {
	points := 0
	if math.Mod(total, 1) == 0 {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	return points
}

func calculatePointsFromItems(items []Item) int {
	points := 0
	points += len(items) / 2 * 5 // 5 points for every 2 items

	for _, item := range items {
		if len(strings.Trim(item.ShortDescription, " "))%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}
	return points
}

func calculatePointsFromPurchaseDate(date string) (int, error) {
	points := 0
	dayStr := date[len(date)-2:] // expected date format: YYYY-MM-DD
	day, err := strconv.Atoi(dayStr)
	if err != nil && day%2 == 1 {
		points += 6
	}
	return points, err
}

func calculatePointsFromPurchaseTime(timeStr string) (int, error) {
	points := 0
	purchaseTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return 0, err
	}
	twoPm, err := time.Parse("15:04", "14:00")
	if err != nil {
		return 0, err
	}
	fourPm, err := time.Parse("15:04", "16:00")
	if err != nil {
		return 0, err
	}

	if purchaseTime.After(twoPm) && purchaseTime.Before(fourPm) {
		points += 10
	}
	return points, err
}

func generateID() string {
	return uuid.NewString()
}

func GetPointsData(id string) (PointsResponse, error) {
	var err error
	points, ok := pointsStore[id]
	if !ok {
		err = ERR_RECEIPT_NOT_FOUND
	}
	return PointsResponse{Points: points}, err
}
