package helpers

import (
	"time"
)

func ReformatDate(date string) string {
	layout := "02/01/2006"
	formatLayout := "2006-01-02"

	t, _ := time.Parse(layout, date)
	return t.Format(formatLayout)
}

func ReformatDateExcel(date string) string {
	layout := "02-01-06"
	formatLayout := "02/01/2006"

	t, _ := time.Parse(layout, date)
	return t.Format(formatLayout)
}
