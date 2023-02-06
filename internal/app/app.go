package app

import (
	"CitiesService/internal/config"
	httpDelivery "CitiesService/internal/delivery/http"
	"CitiesService/internal/repository"
	"CitiesService/internal/repository/file_storage"
	"CitiesService/internal/repository/file_storage/csv_file"
	"CitiesService/internal/repository/file_storage/json_file"
	"CitiesService/internal/repository/storages"
	"CitiesService/internal/server"
	"CitiesService/internal/service"
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
)

func Run(configPath string) {
	var (
		err error
		cfg *config.Config
	)
	cfg, err = config.Init(configPath)
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()
	wg := sync.WaitGroup{}
	citiesCVSNameFile := "data/cities.csv"
	citiesJSONNameFile := "data/cities_storage.json"
	districtsNameFile := "data/districts_storage.json"
	regionsNameFile := "data/regions_storage.json"

	citiesCSVFile := csv_file.CSVFile{citiesCVSNameFile}
	citiesJSONFile := json_file.JSONFile{citiesJSONNameFile}
	districtsJSONFile := json_file.JSONFile{districtsNameFile}
	regionsJSONFile := json_file.JSONFile{regionsNameFile}

	citiesCSVStorage := csv_file.NewCitiesCSVStorage(citiesCSVFile)
	citiesJSONStorage := json_file.NewCitiesJSONStorage(citiesJSONFile)
	districtsJSONStorage := json_file.NewDistrictsJSONStorage(districtsJSONFile)
	regionsJSONStorage := json_file.NewRegionsJSONStorage(regionsJSONFile)
	fileStorage := file_storage.NewFileStorage(citiesJSONStorage, districtsJSONStorage, regionsJSONStorage)

	citiesStorage := new(storages.CitiesStorage)
	regionsStorage := new(storages.RegionsStorage)
	districtsStorage := new(storages.DistrictsStorage)

	repositoryAll := repository.NewRepositories(citiesStorage, regionsStorage, districtsStorage)
	repositoryAll.Init()
	citiesLoader := service.NewCitiesLoader(citiesCSVStorage, fileStorage, repositoryAll)
	citiesUploader := service.NewCitiesUploaderService(citiesCSVStorage, fileStorage, repositoryAll)

	if err = citiesLoader.Load(ctx); err != nil {
		log.Printf("Loader:%s", err.Error())
		cancel()
	} else {
		log.Println("Load completed successfully ")
	}

	serviceDep := service.NewServicesDeps(*repositoryAll)
	serviceApp := service.NewServices(*serviceDep)
	handlers := httpDelivery.NewHandler(serviceApp.Cities, serviceApp.Regions, serviceApp.Districts)

	//HTTP Server
	srv := server.NewServer(cfg, handlers.Init(cfg.HTTP.Host, cfg.HTTP.Port))
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = srv.Run(); err != nil {
			log.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	log.Print("Server started")
	<-ctx.Done()

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err = srv.Stop(context.Background()); err != nil {
			log.Printf("error occured on server shutting down: %s", err.Error())
		}
	}()
	go func() {
		defer wg.Done()
		if err = citiesUploader.Upload(ctx); err != nil {
			log.Printf("Uploader:%s", err.Error())
		} else {
			log.Println("Upload completed successfully ")
		}
	}()
	wg.Wait()
	log.Print("Server stopped")
}
