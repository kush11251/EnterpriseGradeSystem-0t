package main

import (
	"fmt"
	"github.com/EnterpriseGradeSystem/pkg/config"
	"github.com/EnterpriseGradeSystem/pkg/controllers"
)

func main() {
	config.Init()
	controllers.RunServer()
	fmt.Println("Server started on port 8080")
}