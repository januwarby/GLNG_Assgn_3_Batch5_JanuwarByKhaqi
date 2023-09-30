package utils

func GetStatus(value int, category string) string {
	if category == "water" {
		if value < 5 {
			return "aman"
		} else if value >= 6 && value <= 8 {
			return "siaga"
		} else {
			return "bahaya"
		}
	} else if category == "wind" {
		if value < 6 {
			return "aman"
		} else if value >= 7 && value <= 15 {
			return "siaga"
		} else {
			return "bahaya"
		}
	} else {
		return "tidak sesuai"
	}
}
