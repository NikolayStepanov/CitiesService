package storages

import "C"
import (
	"CitiesService/internal/domain"
	"CitiesService/internal/domain/dto"
	"fmt"
	"log"
	"strings"
	"sync"
)

const regionsCap = 100

type RegionsStorage struct {
	sync.Mutex
	regionsMap map[int]*domain.Region
}

func (R *RegionsStorage) GetRegionInfo(ID int) (dto.RegionDTO, error) {
	err := error(nil)
	regionInfo := dto.RegionDTO{}
	if region, ok := R.regionsMap[ID]; !ok {
		err = fmt.Errorf("getRegionInfo: Region with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		regionInfo.ID = region.ID()
		regionInfo.Name = region.Name()
	}
	return regionInfo, err
}

func (R *RegionsStorage) Init() {
	R.regionsMap = make(map[int]*domain.Region, regionsCap)
}

func (R *RegionsStorage) Len() int {
	return len(R.regionsMap)
}

func (R *RegionsStorage) GetAll() dto.RegionsDTO {
	regionsDTO := make([]dto.RegionDTO, len(R.regionsMap))
	indexRegion := 0
	for _, region := range R.regionsMap {
		regionDTO := dto.RegionDTO{ID: region.ID(), Name: region.Name()}
		regionsDTO[indexRegion] = regionDTO
		indexRegion++
	}
	return regionsDTO
}

func (R *RegionsStorage) SearchByName(name string) (int, error) {
	err := error(nil)
	regionID := 0
	for keyID, region := range R.regionsMap {
		if strings.Contains(region.Name(), name) {
			regionID = keyID
			break
		}
	}
	if regionID == 0 {
		err = fmt.Errorf("region with name=%s was not found", name)
		log.Println(err)
	}
	return regionID, err
}

func (R *RegionsStorage) GetName(ID int) (string, error) {
	err := error(nil)
	regionName := ""

	if region, ok := R.regionsMap[ID]; !ok {
		err = fmt.Errorf("getName: Region with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		regionName = region.Name()
	}

	return regionName, err
}

func (R *RegionsStorage) Create(ID int, name string) (string, error) {
	err := error(nil)
	regionName := ""

	if regionVal, ok := R.regionsMap[ID]; ok {
		err = fmt.Errorf("create: Region with ID=%d Name=%s is in storages", regionVal.ID(), regionVal.Name())
		log.Println(err)
	} else {
		region := domain.Region{}
		region.SetID(ID)
		region.SetName(name)
		R.regionsMap[region.ID()] = &region
		regionName = region.Name()
	}
	return regionName, err
}

func (R *RegionsStorage) Delete(ID int) (string, error) {
	err := error(nil)
	regionName := ""

	if regionVal, ok := R.regionsMap[ID]; !ok {
		err = fmt.Errorf("delete: Region with ID=%d is not in storages", ID)
		log.Println(err)
	} else {
		regionName = regionVal.Name()
		delete(R.regionsMap, ID)
	}
	return regionName, err
}

func (R *RegionsStorage) UpdateID(ID int, newID int) (int, error) {
	err := error(nil)
	regionNewID := 0
	if regionVal, ok := R.regionsMap[ID]; !ok {
		err = fmt.Errorf("updata: Region with ID=%d is not in storages", ID)
		log.Println(err)
	} else if ID == newID {
		err = fmt.Errorf("updata: Region ID=%d equals newID=%d", ID, newID)
		log.Println(err)
	} else {
		regionVal.SetID(newID)
		R.regionsMap[newID] = regionVal
		regionNewID = newID
		delete(R.regionsMap, ID)
	}
	return regionNewID, err
}

func (R *RegionsStorage) UpdateName(ID int, newName string) (string, error) {
	err := error(nil)
	regionName := ""

	if regionVal, ok := R.regionsMap[ID]; !ok {
		err = fmt.Errorf("updataName: Region with ID=%d is not in storages", ID)
		log.Println(err)
	} else if regionVal.Name() == newName {
		err = fmt.Errorf("updata: Region ID=%d Name=%s equals new name", ID, regionVal.Name())
		log.Println(err)
	} else {
		regionName = regionVal.Name()
		regionVal.SetName(newName)
	}
	return regionName, err
}
