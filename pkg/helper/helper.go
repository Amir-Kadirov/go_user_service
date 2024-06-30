package helper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NullTimeStampToString(s sql.NullTime) string {
	if s.Valid {
		return s.Time.Format("2006-01-02 15:04:05")
	}

	return ""
}

func NullDateToString(s sql.NullTime) string {
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
	return hour*3600 + minute*60
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

func NullNumberToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return "000001"
}

func GenerateLoginID(db *pgxpool.Pool, user string) (string, error) {
	var currentMaxID string

	if user == "Teacher" {
		query := `SELECT "LoginID"
			  FROM "Teacher"
			  ORDER BY "created_at" DESC
			  LIMIT 1;`
		var loginId sql.NullString
		err := db.QueryRow(context.Background(), query).Scan(&loginId)
		if err != nil {
			return "", err
		}

		logId := NullNumberToString(loginId)

		if logId == "000001" {
			return fmt.Sprintf("T%s", logId), nil
		} else {
			num, _ := strconv.Atoi(logId[1:])
			num = num + 1
			currentMaxID = "T" + fmt.Sprintf("%d", num)
		}
	} else if user == "SupportTeacher" {
		query := `SELECT "LoginID"
			  FROM "SupportTeacher"
			  ORDER BY "created_at" DESC
			  LIMIT 1;`
		var loginId sql.NullString
		err := db.QueryRow(context.Background(), query).Scan(&loginId)
		if err != nil {
			return "", err
		}

		logId := NullNumberToString(loginId)

		if logId == "000001" {
			return fmt.Sprintf("T%s", logId), nil
		} else {
			num, _ := strconv.Atoi(logId[2:])
			num = num + 1
			currentMaxID = "ST" + fmt.Sprintf("%d", num)
		}
	} else if user == "Manager" {
		query := `SELECT "LoginID"
			  FROM "Manager"
			  ORDER BY "created_at" DESC
			  LIMIT 1;`
		var loginId sql.NullString
		err := db.QueryRow(context.Background(), query).Scan(&loginId)
		if err != nil {
			return "", err
		}

		logId := NullNumberToString(loginId)

		if logId == "000001" {
			return fmt.Sprintf("T%s", logId), nil
		} else {
			num, _ := strconv.Atoi(logId[1:])
			num = num + 1
			currentMaxID = "M" + fmt.Sprintf("%d", num)
		}
	} else if user == "Student" {
		query := `SELECT "LoginID"
			  FROM "Student"
			  ORDER BY "created_at" DESC
			  LIMIT 1;`
		var loginId sql.NullString
		err := db.QueryRow(context.Background(), query).Scan(&loginId)
		if err != nil {
			return "", err
		}

		logId := NullNumberToString(loginId)

		if logId == "000001" {
			return fmt.Sprintf("T%s", logId), nil
		} else {
			num, _ := strconv.Atoi(logId[1:])
			num = num + 1
			currentMaxID = "S" + fmt.Sprintf("%d", num)
		}
	} else if user == "Admin" {
		query := `SELECT "LoginID"
			  FROM "Administration"
			  ORDER BY "created_at" DESC
			  LIMIT 1;`
		var loginId sql.NullString
		err := db.QueryRow(context.Background(), query).Scan(&loginId)
		if err != nil {
			return "", err
		}

		logId := NullNumberToString(loginId)

		if logId == "000001" {
			return fmt.Sprintf("A%s", logId), nil
		} else {
			num, _ := strconv.Atoi(logId[1:])
			num = num + 1
			currentMaxID = "A" + fmt.Sprintf("%d", num)
		}
	}

	return currentMaxID, nil
}
