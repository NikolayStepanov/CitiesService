package main

import "CitiesService/internal/app"

const configPath = "configs/main"

// @title Cities Service
// @version 1.0
// @description API Server for Cities Application
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	app.Run(configPath)
}
