package service

import (
	"CitiesService/internal/domain/common"
	"CitiesService/internal/repository"
	"CitiesService/internal/repository/file_storage"
	"CitiesService/internal/repository/file_storage/csv_file"
	"context"
	"golang.org/x/sync/errgroup"
	"log"
)

type CitiesUploaderService struct {
	citiesCSV   *csv_file.CitiesCSVStorage
	storageFile *file_storage.FileStorage
	repository  *repository.Repositories
}

func NewCitiesUploaderService(citiesCSV *csv_file.CitiesCSVStorage, storageFile *file_storage.FileStorage, repository *repository.Repositories) *CitiesUploaderService {
	return &CitiesUploaderService{citiesCSV: citiesCSV, storageFile: storageFile, repository: repository}
}

func (c *CitiesUploaderService) UploadCitiesCSV() error {
	err := error(nil)
	cities := c.repository.Cities.GetAll()
	cities.SortByField(common.ID)
	citiesCSV := make([]csv_file.CityStructCSV, len(cities))
	for index, city := range cities {
		cityCSV := csv_file.CityStructCSV{}
		cityCSV.ID = city.ID
		cityCSV.Name = city.Name
		cityCSV.Foundation = city.Foundation
		cityCSV.Population = city.Population
		cityCSV.RegionName, err = c.repository.Regions.GetName(city.RegionID)
		if err != nil {
			log.Println(err)
		}
		cityCSV.DistrictName, err = c.repository.Districts.GetName(city.DistrictID)
		if err != nil {
			log.Println(err)
		}
		citiesCSV[index] = cityCSV
	}
	err = c.citiesCSV.WriteAll(citiesCSV)
	return err
}

func (c *CitiesUploaderService) UploadCities() error {
	err := error(nil)
	cities := c.repository.Cities.GetAllSortByField(common.NAME)
	err = c.storageFile.CitiesFileStorage.WriteAll(cities)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (c *CitiesUploaderService) UploadDistricts() error {
	err := error(nil)
	districts := c.repository.Districts.GetAll()
	districts.SortByField(common.ID)
	err = c.storageFile.DistrictsFileStorage.WriteAll(districts)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (c *CitiesUploaderService) UploadRegions() error {
	err := error(nil)
	regions := c.repository.Regions.GetAll()
	regions.SortByField(common.ID)
	err = c.storageFile.RegionsFileStorage.WriteAll(regions)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (c *CitiesUploaderService) Upload(ctx context.Context) error {
	err := error(nil)

	errs, _ := errgroup.WithContext(ctx)
	errs.Go(func() error {
		if err = c.UploadRegions(); err != nil {
			log.Println(err)
		}
		return err
	})
	errs.Go(func() error {
		if err = c.UploadDistricts(); err != nil {
			log.Println(err)
		}
		return err
	})
	errs.Go(func() error {
		if err = c.UploadCities(); err != nil {
			log.Println(err)
		}
		return err
	})
	errs.Go(func() error {
		if err = c.UploadCitiesCSV(); err != nil {
			log.Println(err)
		}
		return err
	})
	err = errs.Wait()
	return err
}
