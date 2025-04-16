package services

import (
	"testing"
)

func TestCountAlphanumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Empty string", "", 0},
		{"Only letters", "Target", 6},
		{"Mixed characters", "M&M Corner Market", 14},
		{"With numbers", "123 Main St", 9},
		{"Only special chars", "!@#$%^&*()", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := countAlphanumeric(test.input)
			if result != test.expected {
				t.Errorf("countAlphanumeric(%s) = %d; want %d", test.input, result, test.expected)
			}
		})
	}
}

func TestCalculatePointsFromTotal(t *testing.T) {
	tests := []struct {
		name     string
		total    float64
		expected int
	}{
		{"Even dollar amount", 100.00, 75}, // 50 (round) + 25 (multiple of 0.25)
		{"Quarter amount", 25.75, 25},      // Multiple of 0.25 but not round
		{"Dime amount", 10.10, 0},          // Neither round nor multiple of 0.25
		{"Penny amount", 12.01, 0},         // Neither round nor multiple of 0.25
		{"Zero amount", 0.00, 75},          // 0 is both round and multiple of 0.25
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculatePointsFromTotal(test.total)
			if result != test.expected {
				t.Errorf("calculatePointsFromTotal(%f) = %d; want %d", test.total, result, test.expected)
			}
		})
	}
}

func TestCalculatePointsFromItems(t *testing.T) {
	tests := []struct {
		name     string
		items    []Item
		expected int
	}{
		{
			"Empty items",
			[]Item{},
			0,
		},
		{
			"One item, no length bonus",
			[]Item{
				{ShortDescription: "Milk", Price: "2.50"},
			},
			0,
		},
		{
			"Two items, one length bonus",
			[]Item{
				{ShortDescription: "Milk", Price: "2.50"},
				{ShortDescription: "Bread", Price: "3.50"},
			},
			5,
		},
		{
			"Three items, one length bonus",
			[]Item{
				{ShortDescription: "Milk", Price: "2.50"},
				{ShortDescription: "Eggs", Price: "3.00"},
				{ShortDescription: "Pot", Price: "10.00"}, // length bonus: 2
			},
			7,
		},
		{
			"Four items, 2 length bonuses",
			[]Item{
				{ShortDescription: "Milk", Price: "2.50"},
				{ShortDescription: "Egg", Price: "3.00"},  // length bonus: 1
				{ShortDescription: "Pot", Price: "10.00"}, // length bonus: 2
				{ShortDescription: "Rice", Price: "5.75"},
			},
			13,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculatePointsFromItems(test.items)
			if result != test.expected {
				t.Errorf("calculatePointsFromItems() = %d; want %d", result, test.expected)
			}
		})
	}
}

func TestCalculatePointsFromPurchaseDate(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected int
	}{
		{"Odd day", "2022-01-01", 6},
		{"Even day", "2022-01-02", 0},
		{"Another odd day", "2022-01-23", 6},
		{"Another even day", "2022-01-30", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculatePointsFromPurchaseDate(test.date)
			if err != nil {
				t.Fatalf("calculatePointsFromPurchaseDate(%s) returned error: %v", test.date, err)
			}
			if result != test.expected {
				t.Errorf("calculatePointsFromPurchaseDate(%s) = %d; want %d", test.date, result, test.expected)
			}
		})
	}
}

func TestCalculatePointsFromPurchaseTime(t *testing.T) {
	tests := []struct {
		name     string
		time     string
		expected int
	}{
		{"Before 2PM", "13:59", 0},  // Before 2PM
		{"At 2PM", "14:00", 0},      // 2PM exactly
		{"After 2PM", "14:01", 10},  // After 2PM
		{"Before 4PM", "15:59", 10}, // Before 4PM
		{"At 4PM", "16:00", 0},      // 4PM exactly
		{"After 4PM", "16:01", 0},   // After 4PM
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculatePointsFromPurchaseTime(test.time)
			if err != nil {
				t.Fatalf("calculatePointsFromPurchaseTime(%s) returned error: %v", test.time, err)
			}
			if result != test.expected {
				t.Errorf("calculatePointsFromPurchaseTime(%s) = %d; want %d", test.time, result, test.expected)
			}
		})
	}
}
