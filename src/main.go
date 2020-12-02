package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	db := connectDb(dbConfigs.DbName)
	defer db.Close()

	results := getAllScheduledPlan(db, ctx)
	for i, res := range results {
		fmt.Println(i, res)
		fmt.Println(startJenkinsBuild("cq_test_external"))
	}

	//getTestcase(db, ctx, "TMlt00000208")
	//main2.generateExcel(db, ctx, "TMlt00000208")
}
