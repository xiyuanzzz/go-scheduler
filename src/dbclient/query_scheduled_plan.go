package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type ScheduledResult struct {
	Id int
	PlanId string
	OsIp string
	OsUser string
	OsPasswd string
	BmcIp string
	BmcUser string
	BmcPasswd string
}

func queryAllNonExecPlan(db *sql.DB, ctx context.Context) *sql.Rows {
	fmt.Println("\n... 正在查询所有待执行的测试计划")
	queryStr := "SELECT id, plan_id, os_ip, os_user, os_passwd, bmc_ip, bmc_user, bmc_passwd " +
		"FROM tab_project " +
		"WHERE task_auto_all > 0 AND flag_action = false;"

	stmt, err := db.PrepareContext(ctx, queryStr)
	defer stmt.Close()
	if err != nil {
		log.Printf("Prepare statement of all non execution plans error: %v", err)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("Query all non execution plans error: %v", err)
	}

	return rows
}

func queryAllFlyPlan(db *sql.DB, ctx context.Context) *sql.Rows {
	fmt.Println("\n... 正在查询所有冲突的测试计划")
	queryStr := "SELECT id, os_ip " +
		"FROM tab_project " +
		"WHERE task_auto_all > 0 AND flag_action = true AND flag_finish = false;"

	stmt, err := db.PrepareContext(ctx, queryStr)
	defer stmt.Close()
	if err != nil {
		log.Printf("Prepare statement of all ongoing plans error: %v", err)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("Query all ongoing plans error: %v", err)
	}

	return rows
}

func getAllScheduledPlan(db *sql.DB, ctx context.Context) []ScheduledResult {
	allNonExecPlan := queryAllNonExecPlan(db, ctx)
	defer allNonExecPlan.Close()

	var scheduledRows []ScheduledResult
	for allNonExecPlan.Next() {
		var row = ScheduledResult{}
		if err := allNonExecPlan.Scan(&row.Id, &row.PlanId, &row.OsIp, &row.OsUser, &row.OsPasswd, &row.BmcIp, &row.BmcUser, &row.BmcPasswd); err != nil {
			log.Printf("Scan non exection row error: %v", err)
		}
		scheduledRows = append(scheduledRows, row)
	}
	fmt.Printf("所有待执行的测试计划:\n%+v\n", scheduledRows)

	allFlyPlan := queryAllFlyPlan(db, ctx)
	defer allFlyPlan.Close()

	var flyRows = map[string]int{}
	for allFlyPlan.Next() {
		var (
			Id int
			OsIp string
		)
		if err := allFlyPlan.Scan(&Id, &OsIp); err != nil {
			log.Printf("Scan ongoing row error: %v", err)
		}
		flyRows[OsIp] = Id
	}
	fmt.Printf("所有冲突的测试计划:\n%+v\n", flyRows)

	for i,j := 0,len(scheduledRows); i < j; {
		row := scheduledRows[i]
		if _, ok := flyRows[row.OsIp]; ok {
			scheduledRows = append(scheduledRows[:i], scheduledRows[i+1:]...)
			j--
		} else {
			i++
		}
	}
	fmt.Printf("\n过滤冲突后的测试计划:\n%+v\n", scheduledRows)
	return scheduledRows
}

/*func main() {
        ctx, cancelFunc := context.WithCancel(context.Background())
        defer cancelFunc()
        db := connectDb(dbConfigs.DbName)
        defer db.Close()

        results := getAllScheduledPlan(db, ctx)
        fmt.Printf("\n%+v\n", results)
}*/