package repository

import (
	"MyLog-M/internal/domain"
	"database/sql"
	"fmt"
	"time"
)

type LogRepo struct {
	db *sql.DB
}

func NewLogRepo(db *sql.DB) *LogRepo {
	return &LogRepo{db: db}
}

//GetLogTail shows the last n records
func (this *LogRepo) Tail(limit int64) (*[]domain.Data, error) {
	var result []domain.Data
	sql := "SELECT `RID`, `UnixTime`, `Type`, `Severity`, `Log` FROM Log "
	sql += "ORDER BY RID DESC "
	if limit > 0 {
		sql += fmt.Sprintf("LIMIT %d ", limit)
	}
	fmt.Println(sql)

	rows, err := this.db.Query(sql)
	if err != nil {
		return &result, err
	}

	defer rows.Close()
	for rows.Next() {
		item := domain.Data{}
		if err := rows.Scan(&item.RID, &item.UnixTime, &item.LogType, &item.LogSeverity, &item.LogText); err != nil {
			fmt.Println(err)
			return &result, err
		}
		item.LocalTime = time.Unix(int64(item.UnixTime), 0)
		fmt.Println((item))
		result = append(result, item)
	}
	return &result, nil
}
