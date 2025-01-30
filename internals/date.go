package internals

import "time"

func IsDateValue(stringDate string, layout string) bool {
	_, err := time.Parse(layout, stringDate)
	return err == nil
}

func CreateDateAsDateTime() string {
	return time.Now().Format(time.DateTime)
}
