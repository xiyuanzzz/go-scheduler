package main

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"log"
)

func startJenkinsBuild(jobName string) int64 {
	fmt.Println("\n... 正在连接Jenkins")
	jenkins := gojenkins.CreateJenkins(nil, "http://10.151.12.99:8081/", "root", "ips@2019")
	_, err := jenkins.Init()
	if err != nil {
		log.Printf("Jenkins init error: %v", err)
	}
	fmt.Println("*** Jenkins已连接 ***")

	fmt.Printf("... 即将触发任务 %s\n", jobName)
	job, err := jenkins.GetJob(jobName)
	if err != nil {
		log.Printf("Jenkins get job error: %v", err)
	}

	build, err := job.GetLastBuild()
	if err != nil {
		log.Printf("Jenkins get build error: %v", err)
	}
	buildNum := build.GetBuildNumber() + 1

	_, err = jenkins.BuildJob(jobName)
	if err != nil {
		log.Printf("Jenkins build job error: %v", err)
	}
	fmt.Printf("*** 任务 %s 构建 #%d 已触发 ***\n", jobName, buildNum)

	return buildNum
}

/*func main() {
	fmt.Println(startJenkinsBuild("cq_test_external"))
}*/