package helper

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func NullTimeStampToString(s sql.NullTime) string {
	if s.Valid {
		return s.Time.Format("2006-01-02 15:04:05")
	}

	return ""
}


func NullDateToString(s sql.NullTime) string  {
	if s.Valid {
		return s.Time.Format("2006-01-02")
	}

	return ""
}

func DeleteChecker(s sql.NullTime) error {
	if s.Valid {
		return nil
	}

	return errors.New("does not exist")
}

func TimeToSecond(time string) int {
	hour, _ := strconv.Atoi(time[:2])
	minute, _ := strconv.Atoi(time[3:])
	return 	hour*3600 + minute*60
}

func SecondToTime(second int) string {
	hours := second / 3600
	minutes := (second % 3600) / 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func DateSince(s sql.NullTime) int {
	if s.Valid {
		now := time.Now()
		then := s.Time

		years := now.Year() - then.Year()
		months := int(now.Month()) - int(then.Month())
		days := now.Day() - then.Day()

		if days < 0 {
			months -= 1
		}
		if months < 0 {
			years -= 1
			months += 12
		}

		totalMonths := years*12 + months
		return totalMonths 
	}

	return 0
}