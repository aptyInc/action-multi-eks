package main

import (
	"log"
	"os"

	"github.com/actions-go/toolkit/core"
)

func runMain() {
	region := os.Getenv("REGION")
	branchName := os.Getenv("BRANCH_NAME")

	if region == "" || branchName == "" {
		log.Fatal("REGION and BRANCH_NAME must be set")
	}

	var EKS_NAME string
	CROSS_ACCOUNT := "false"

	switch region {
	case "us-east-1":
		EKS_NAME = branchName + "-eks"
	case "ap-southeast-2":
		EKS_NAME = "au-" + branchName + "-eks"
	case "eu-central-1":
		{
			EKS_NAME = "eu1-" + branchName + "-eks"
			CROSS_ACCOUNT = "true"
		}
	}

	core.SetOutput("EKS_NAME", EKS_NAME)
	core.SetOutput("CROSS_ACCOUNT", CROSS_ACCOUNT)
}

func main() {
	runMain()
}
