package service

import (
	"CitiesService/internal/domain/dto"
	"CitiesService/internal/repository"
	"CitiesService/internal/repository/file_storage"
	"CitiesService/internal/repository/file_storage/csv_file"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

type CitiesLoader struct {
	citiesCSV   *csv_file.CitiesCSVStorage
	storageFile *file_storage.FileStorage
	repository  *repository.Repositories
}

func NewCitiesLoader(citiesCSV *csv_file.CitiesCSVStorage, storageFile *file_storage.FileStorage, repository *repository.Repositories) *CitiesLoader {
	return &CitiesLoader{citiesCSV: citiesCSV, storageFile: storageFile, repository: repository}
}

func (c *CitiesLoader) LoadCityInStorage(city csv_file.CityStructCSV) error {
	err := error(nil)
	if _, err = c.repository.Cities.Create(city.ID, city.Name); err != nil {
		err = fmt.Errorf("city id=%d name=%s not created", city.ID, city.Name)
		log.Println(err)
	} else {
		errUpdateRegionID := error(nil)
		errUpdateDistrictID := error(nil)
		errUpdatePopulation := error(nil)
		errUpdateFoundation := error(nil)
		errSearchRegion := error(nil)
		errSearchDistrict := error(nil)

		cityRegionID := 0
		cityDistrictID := 0

		_, errUpdatePopulation = c.repository.Cities.UpdatePopulation(city.ID, city.Population)
		_, errUpdateFoundation = c.repository.Cities.UpdateFoundation(city.ID, int(city.Foundation))

		cityRegionID, errSearchRegion = c.repository.Regions.SearchByName(city.RegionName)
		cityDistrictID, errSearchDistrict = c.repository.Districts.SearchByName(city.DistrictName)

		if errSearchRegion == nil {
			_, errUpdateRegionID = c.repository.Cities.UpdateRegionID(city.ID, cityRegionID)
		}
		if errSearchDistrict == nil {
			_, errUpdateDistrictID = c.repository.Cities.UpdateDistrictID(city.ID, cityDistrictID)
		}

		if errSearchRegion != nil || errSearchDistrict != nil ||
			errUpdatePopulation != nil || errUpdateFoundation != nil ||
			errUpdateRegionID != nil || errUpdateDistrictID != nil {
			err = fmt.Errorf("city created but data updated incorrectly")
		}
	}
	return err
}

func (c *CitiesLoader) LoadDistrictInStorage(district dto.DistrictDTO) error {
	err := error(nil)
	districtID := 0

	districtID = district.ID
	if districtID != 0 {
		if _, err = c.repository.Districts.Create(districtID, district.Name); err != nil {
			err = fmt.Errorf("district id=%d name=%s not created", districtID, district.Name)
		}
	} else {
		err = fmt.Errorf("district id=%d not created", districtID)
	}

	return err
}

func (c *CitiesLoader) LoadRegionInStorage(region dto.RegionDTO) error {
	err := error(nil)
	regionID := 0

	regionID = region.ID
	if regionID != 0 {
		if _, err = c.repository.Regions.Create(regionID, region.Name); err != nil {
			err = fmt.Errorf("region id=%d name=%s not created", region.ID, region.Name)
		}
	} else {
		err = fmt.Errorf("region id=%d not created", regionID)
	}
	return err
}

func (c *CitiesLoader) LoadCities() error {
	err := error(nil)
	bSuccess := true

	citiesCSV := []csv_file.CityStructCSV{}
	if citiesCSV, err = c.citiesCSV.ReadAll(); err != nil {
		panic(err)
	}
	for _, city := range citiesCSV {
		if errLoadCity := c.LoadCityInStorage(city); errLoadCity != nil {
			log.Println(errLoadCity)
			bSuccess = false
		}
	}
	if !bSuccess {
		err = errors.New("load cities went wrong")
	}
	return err
}

func (c *CitiesLoader) LoadDistricts() error {
	err := error(nil)
	bSuccess := true
	districts := []dto.DistrictDTO{}

	if districts, err = c.storageFile.DistrictsFileStorage.ReadAll(); err != nil {
		panic(err)
	}
	for _, district := range districts {
		if errLoadDistrict := c.LoadDistrictInStorage(district); errLoadDistrict != nil {
			log.Println(errLoadDistrict)
			bSuccess = false
		}
	}
	if !bSuccess {
		err = errors.New("load districts went wrong")
	}
	return err
}

func (c *CitiesLoader) LoadRegions() error {
	err := error(nil)
	bSuccess := true
	regions := []dto.RegionDTO{}

	if regions, err = c.storageFile.RegionsFileStorage.ReadAll(); err != nil {
		panic(err)
	}
	for _, region := range regions {
		if errLoadRegion := c.LoadRegionInStorage(region); errLoadRegion != nil {
			log.Println(errLoadRegion)
			bSuccess = false
		}
	}
	if !bSuccess {
		err = errors.New("load districts went wrong")
	}
	return err
}

func (c *CitiesLoader) Load(ctx context.Context) error {
	err := error(nil)
	errs, _ := errgroup.WithContext(ctx)

	errs.Go(func() error {
		if err = c.LoadRegions(); err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	errs.Go(func() error {
		if err = c.LoadDistricts(); err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	err = errs.Wait()
	if err == nil {
		if err = c.LoadCities(); err != nil {
			log.Println(err)
		}
	}
	return err
}
