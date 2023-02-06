package http

import (
	"CitiesService/internal/domain/dto"
	"fmt"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type UpdateCityRequest struct {
	UpdateRequest
	NewRegionID   int `json:"new region_id,omitempty"`
	NewDistrictID int `json:"new district_id,omitempty"`
	NewPopulation int `json:"new population,omitempty"`
	NewFoundation int `json:"new foundation,omitempty"`
}

type ListCitiesInput struct {
	ListInput
	RegionsID       []int `json:"regionID" in:"query=regionID[],regionID"`
	DistrictsID     []int `json:"districtsID" in:"query=districtID[],districtID"`
	LowerPopulation int   `json:"population_min" in:"query=population_min"`
	UpperPopulation int   `json:"population_max" in:"query=population_max"`
	LowerFoundation int   `json:"foundation_min" in:"query=foundation_min"`
	UpperFoundation int   `json:"foundation_max" in:"query=foundation_max"`
}

func (h *Handler) initCitiesRoutes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{cityID}", func(r chi.Router) {
		r.Get("/", h.GetCityInformation)
		r.Put("/", h.UpdateCity)
	})
	r.With(httpin.NewInput(ListCitiesInput{})).Get("/", h.GetCitiesInformation)
	r.Delete("/", h.CityDelete)
	r.Route("/create", func(r chi.Router) {
		r.Post("/", h.CreateCity)
	})
	return r
}

// @Summary CreateCity
// @Description Сreating a new city entry
// @Tags Cities
// @Accept json
// @Produce json
// @Param input body dto.CityDTO true "json information City"
// @Success 200 {object} dto.CityDTO
// @Failure 400 {object} ErrResponse
// @Router /cities/create [post]
func (h *Handler) CreateCity(w http.ResponseWriter, r *http.Request) {
	err := error(nil)
	cityCreate := dto.CityDTO{}
	if err = render.Decode(r, &cityCreate); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if _, err = h.regionsService.GetInformationRegion(cityCreate.RegionID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if _, err = h.districtsService.GetInformationDistrict(cityCreate.DistrictID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if err = h.citiesService.CreateCity(cityCreate); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, cityCreate)
}

func (h *Handler) UpdateCityID(cityID, newID int) (int, string, error) {
	err := error(nil)
	cityNewID := dto.CityDTO{}
	messageResponse := ""
	if cityNewID, err = h.citiesService.GetInformationCity(newID); err != nil {
		err = h.citiesService.UpdateIDCity(cityID, newID)
		if err != nil {
			messageResponse += fmt.Sprintf("City ID: %s\n", err.Error())
		} else {
			messageResponse += fmt.Sprintf("City ID = %d OldID = %d \n", newID, cityID)
			cityID = newID
		}
	} else {
		messageResponse += fmt.Sprintf("City with ID = %d "+
			"Name = %s already exists\n", cityNewID.ID, cityNewID.Name)
		err = fmt.Errorf(messageResponse)
	}
	return cityID, messageResponse, err
}

func (h *Handler) UpdateCityRegionID(cityID, newRegionID, regionID int) string {
	err := error(nil)
	messageResponse := ""
	if _, err = h.regionsService.GetInformationRegion(newRegionID); err != nil {
		messageResponse += fmt.Sprintf("RegionID: %s\n", err.Error())
	} else {
		err = h.citiesService.UpdateRegionIDCity(cityID, newRegionID)
		if err != nil {
			messageResponse += fmt.Sprintf("RegionID: %s\n", err.Error())
		} else {
			messageResponse += fmt.Sprintf("RegionID = %d "+
				"OldRegionID = %d\n", newRegionID, regionID)
		}
	}
	return messageResponse
}

func (h *Handler) UpdateCityDistrictID(cityID, newDistrictID, districtID int) string {
	err := error(nil)
	messageResponse := ""
	if _, err = h.districtsService.GetInformationDistrict(newDistrictID); err != nil {
		messageResponse += fmt.Sprintf("DistrictID: %s\n", err.Error())
	} else {
		err = h.citiesService.UpdateDistrictIDCity(cityID, newDistrictID)
		if err != nil {
			messageResponse += fmt.Sprintf("DistrictID: %s\n", err.Error())
		} else {
			messageResponse += fmt.Sprintf("DistrictID = %d "+
				"OldDistrictID = %d\n", newDistrictID, districtID)
		}
	}
	return messageResponse
}

func (h *Handler) UpdateCityName(cityID int, newName, name string) string {
	err := error(nil)
	messageResponse := ""
	err = h.citiesService.UpdateNameCity(cityID, newName)
	if err != nil {
		messageResponse += fmt.Sprintf("Name: %s\n", err.Error())
	} else {
		messageResponse += fmt.Sprintf("Name = %s OldName = %s\n", newName, name)
	}
	return messageResponse
}

func (h *Handler) UpdateCityFoundation(cityID, newFoundation, foundation int) string {
	err := error(nil)
	messageResponse := ""
	err = h.citiesService.UpdateFoundationCity(cityID, newFoundation)
	if err != nil {
		messageResponse += fmt.Sprintf("Foundation: %s\n", err.Error())
	} else {
		messageResponse += fmt.Sprintf("Foundation = %d "+
			"OldFoundation = %d\n", newFoundation, foundation)
	}
	return messageResponse
}

func (h *Handler) UpdateCityPopulation(cityID, newPopulation, population int) string {
	err := error(nil)
	messageResponse := ""
	err = h.citiesService.UpdatePopulationCity(cityID, newPopulation)
	if err != nil {
		messageResponse += fmt.Sprintf("Population: %s\n", err.Error())
	} else {
		messageResponse += fmt.Sprintf("Population = %d "+
			"OldPopulation = %d\n", newPopulation, population)
	}
	return messageResponse
}

// @Summary UpdateCity
// @Description Сity information update
// @Tags Cities
// @Accept json
// @Produce html
// @Param id path int true "City ID"
// @Param updateCityRequest body UpdateCityRequest false "json update information City"
// @Success 200 {string} string
// @Failure 400 {object} ErrResponse
// @Router /cities/{id} [put]
func (h *Handler) UpdateCity(w http.ResponseWriter, r *http.Request) {
	var (
		err             error
		cityID          int
		city            dto.CityDTO
		dataRequest     UpdateCityRequest
		messageResponse string = "Updated:\n"
	)
	cityID, err = strconv.Atoi(chi.URLParam(r, "cityID"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if city, err = h.citiesService.GetInformationCity(cityID); err != nil {
		messageResponse += fmt.Sprintf("City: %s\n", err.Error())
	} else {
		render.Decode(r, &dataRequest)
		if dataRequest.NewID != 0 {
			messageResponseCityID := ""
			cityID, messageResponseCityID, err = h.UpdateCityID(cityID, dataRequest.NewID)
			messageResponse += messageResponseCityID
			if err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}
		}
		if dataRequest.NewRegionID != 0 {
			messageResponseRegionID := h.UpdateCityRegionID(cityID, dataRequest.NewRegionID, city.RegionID)
			messageResponse += messageResponseRegionID
		}
		if dataRequest.NewDistrictID != 0 {
			messageResponseDistrictID := h.UpdateCityDistrictID(cityID, dataRequest.NewDistrictID, city.DistrictID)
			messageResponse += messageResponseDistrictID
		}
		if dataRequest.NewName != "" {
			messageResponseName := h.UpdateCityName(cityID, dataRequest.NewName, city.Name)
			messageResponse += messageResponseName
		}
		if dataRequest.NewFoundation != 0 {
			messageResponseFoundation := h.UpdateCityFoundation(cityID, dataRequest.NewFoundation, city.Foundation)
			messageResponse += messageResponseFoundation
		}
		if dataRequest.NewPopulation != 0 {
			messageResponsePopulation := h.UpdateCityPopulation(cityID, dataRequest.NewPopulation, city.Population)
			messageResponse += messageResponsePopulation
		}
	}
	render.Status(r, http.StatusOK)
	render.HTML(w, r, messageResponse)
}

// @Summary GetCityInformation
// @Description Getting information about the city
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "City ID"
// @Success 200 {object} dto.CityDTO
// @Failure 400 {object} ErrResponse
// @Router /cities/{id} [get]
func (h *Handler) GetCityInformation(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		cityID  int
		cityDTO dto.CityDTO
	)
	cityID, err = strconv.Atoi(chi.URLParam(r, "cityID"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	cityDTO, err = h.citiesService.GetInformationCity(cityID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, cityDTO)
}

// @Summary GetCitiesInformation
// @Description Getting information about cities
// @Tags Cities
// @Accept json
// @Produce json
// @Param collection query ListCitiesInput false "string collection" collectionFormat(multi)
// @Success 200 {object} dto.CitiesDTO
// @Router /cities/ [get]
func (h *Handler) GetCitiesInformation(w http.ResponseWriter, r *http.Request) {
	var (
		citiesDTO dto.CitiesDTO
	)
	citiesDTO = h.citiesService.GetInformationCities()
	input := r.Context().Value(httpin.Input).(*ListCitiesInput)
	if input.ID != nil {
		citiesDTO = citiesDTO.SelectByID(input.ID...)
	}
	if input.Name != nil {
		citiesDTO = citiesDTO.SelectByName(input.Name...)
	}
	if input.DistrictsID != nil {
		citiesDTO = citiesDTO.SelectByDistrictID(input.DistrictsID...)
	}
	if input.RegionsID != nil {
		citiesDTO = citiesDTO.SelectByRegionID(input.RegionsID...)
	}
	if input.LowerPopulation < input.UpperPopulation {
		citiesDTO = citiesDTO.SelectRangePopulation(input.LowerPopulation, input.UpperPopulation)
	}
	if input.LowerFoundation < input.UpperFoundation {
		citiesDTO = citiesDTO.SelectRangeFoundation(input.LowerFoundation, input.UpperFoundation)
	}
	if input.SortBy != "" {
		citiesDTO.SortByField(input.SortBy)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, citiesDTO)
}

// @Summary CityDelete
// @Description Delete city information
// @Tags Cities
// @Accept json
// @Produce json
// @Param requestDelete body RequestDelete true "json delete targetID City"
// @Success 200 {object} dto.CityDTO
// @Failure 400 {object} ErrResponse
// @Router /cities/ [delete]
func (h *Handler) CityDelete(w http.ResponseWriter, r *http.Request) {
	var (
		deleteCityID int
		err          error
		dataRequest  RequestDelete
		dataResponse dto.CityDTO
	)
	render.Decode(r, &dataRequest)
	deleteCityID = dataRequest.TargetID
	dataResponse, err = h.citiesService.GetInformationCity(dataRequest.TargetID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	_, err = h.citiesService.DeleteCity(deleteCityID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, dataResponse)
}
