package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Receipt JSON structure
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item in the receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Validate checks if the receipt conforms to the schema
func (r Receipt) Validate() error {
	if !regexp.MustCompile(`^[\w\s\-&]+$`).MatchString(r.Retailer) {
		return fmt.Errorf("invalid retailer format")
	}

	if _, err := time.Parse("2006-01-02", r.PurchaseDate); err != nil {
		return fmt.Errorf("invalid purchaseDate format")
	}

	if _, err := time.Parse("15:04", r.PurchaseTime); err != nil {
		return fmt.Errorf("invalid purchaseTime format")
	}

	if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(r.Total) {
		return fmt.Errorf("invalid total format")
	}

	if len(r.Items) == 0 {
		return fmt.Errorf("items array cannot be empty")
	}

	// Validate each item
	for _, item := range r.Items {
		if !regexp.MustCompile(`^[\w\s\-]+$`).MatchString(item.ShortDescription) {
			return fmt.Errorf("invalid shortDescription format")
		}
		if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(item.Price) {
			return fmt.Errorf("invalid price format")
		}
	}

	return nil
}

func generateUUID() string {
	return uuid.New().String()
}

func calculatePoints(receipt Receipt) int {
	points := 0

	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(re.FindAllString(receipt.Retailer, -1))

	total, _ := strconv.ParseFloat(receipt.Total, 64)

	if total == math.Floor(total) {
		points += 50
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	start, _ := time.Parse("15:04", "14:00")
	end, _ := time.Parse("15:04", "16:00")
	if purchaseTime.After(start) && purchaseTime.Before(end) {
		points += 10
	}

	return points
}
