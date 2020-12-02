package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func queryTestSet(ctx context.Context, db *sql.DB, planId string) *sql.Rows {
	fmt.Printf("\n... 正在查询测试计划 %s 的所有测试用例\n", planId)

	queryStr := "SELECT task_id " +
		"FROM tab_task " +
		"WHERE plan_id = ?"
	stmt, err := db.PrepareContext(ctx, queryStr)
	defer stmt.Close()
	if err != nil {
		log.Printf("Prepare statement of all test cases error: %v", err)
	}

	rows, err := stmt.QueryContext(ctx, planId)
	if err != nil {
		log.Printf("Query all test cases error: %v", err)
	}

	return rows
}

func prepareTestSet(rows sql.Rows) []string {
	var testSet []string

	for rows.Next() {
		testId := ""
		if err := rows.Scan(&testId); err == nil {
			testSet = append(testSet, testId)
		} else {
			log.Printf("Scan test case row error: %v", err)
		}
	}

	fmt.Printf("共 %d 个测试用例：%+v\n", len(testSet), testSet)
	return testSet
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	db := connectDb(dbConfigs.DbName)
	defer db.Close()

	rows := queryTestSet(ctx, db, "TMlt00003008")
}
