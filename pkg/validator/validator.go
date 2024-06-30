package validator

import (
	"errors"
	"regexp"
	"time"
)

func ValidateGmail(gmail string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail.com$`).MatchString(gmail)
}
func ValidatePhone(phone string) bool {
	regex := `^\+998\d{9}$`
	return regexp.MustCompile(regex).MatchString(phone)
}

func CheckDeadline(timestamp string) (float64, error) {
	layout := time.RFC3339

	date, err := time.Parse(layout, timestamp)
	if err != nil {
		return -1, errors.New("wrong timestamp format")
	}

	now := time.Now()

	hoursUntil := date.Sub(now).Hours()

	if hoursUntil < 0 {
		return 0, nil
	}

	return hoursUntil, nil
}
