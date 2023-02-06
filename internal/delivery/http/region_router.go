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

func (h *Handler) initRegionRoutes() chi.Router {
	r := chi.NewRouter()
	r.Route("/{regionID}", func(r chi.Router) {
		r.Get("/", h.GetRegionInformation)
		r.Put("/", h.UpdateRegion)
	})
	r.With(httpin.NewInput(ListInput{})).Get("/", h.GetRegionsInformation)
	r.Delete("/", h.RegionDelete)
	r.Route("/create", func(r chi.Router) {
		r.Post("/", h.CreateRegion)
	})
	return r
}

// @Summary CreateRegion
// @Description Сreating a new region entry
// @Tags Regions
// @Accept json
// @Produce json
// @Param input body dto.RegionDTO true "json information Region"
// @Success 200 {object} dto.RegionDTO
// @Failure 400 {object} ErrResponse
// @Router /regions/create [post]
func (h *Handler) CreateRegion(w http.ResponseWriter, r *http.Request) {
	err := error(nil)
	regionCreate := dto.RegionDTO{}
	if err = render.Decode(r, &regionCreate); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err = h.regionsService.CreateRegion(regionCreate); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, regionCreate)
}

func (h *Handler) UpdateRegionID(regionID, newID int) (int, string, error) {
	err := error(nil)
	regionNewID := dto.RegionDTO{}
	messageResponse := ""
	if regionNewID, err = h.regionsService.GetInformationRegion(newID); err != nil {
		err = h.regionsService.UpdateIDRegion(regionID, newID)
		if err != nil {
			messageResponse += fmt.Sprintf("Region ID: %s\n", err.Error())
		} else {
			messageResponse += fmt.Sprintf("Region ID = %d OldID = %d \n", newID, regionID)
			regionID = newID
		}
	} else {
		messageResponse += fmt.Sprintf("Region with ID = %d "+
			"Name = %s already exists\n", regionNewID.ID, regionNewID.Name)
		err = fmt.Errorf(messageResponse)
	}
	return regionID, messageResponse, err
}

func (h *Handler) UpdateRegionName(districtID int, newName, name string) string {
	err := error(nil)
	messageResponse := ""
	err = h.regionsService.UpdateNameRegion(districtID, newName)
	if err != nil {
		messageResponse += fmt.Sprintf("Name: %s\n", err.Error())
	} else {
		messageResponse += fmt.Sprintf("Name = %s OldName = %s\n", newName, name)
	}
	return messageResponse
}

// @Summary UpdateRegion
// @Description Region information update
// @Tags Regions
// @Accept json
// @Produce html
// @Param id path int true "Region ID"
// @Param updateDistrictRequest body UpdateRequest true "json update information Region"
// @Success 200 {string} string
// @Failure 400 {object} ErrResponse
// @Router /regions/{id} [put]
func (h *Handler) UpdateRegion(w http.ResponseWriter, r *http.Request) {
	var (
		err             error
		regionID        int
		regionDTO       dto.RegionDTO
		dataRequest     UpdateRequest
		messageResponse string = "Updated:\n"
	)
	regionID, err = strconv.Atoi(chi.URLParam(r, "regionID"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if regionDTO, err = h.regionsService.GetInformationRegion(regionID); err != nil {
		messageResponse += fmt.Sprintf("Region: %s\n", err.Error())
	} else {
		render.Decode(r, &dataRequest)
		if dataRequest.NewID != 0 {
			messageResponseCityID := ""
			regionID, messageResponseCityID, err = h.UpdateRegionID(regionID, dataRequest.NewID)
			messageResponse += messageResponseCityID
			if err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}
		}
		if dataRequest.NewName != "" {
			messageResponseName := h.UpdateRegionName(regionID, dataRequest.NewName, regionDTO.Name)
			messageResponse += messageResponseName
		}
	}
	render.Status(r, http.StatusOK)
	render.HTML(w, r, messageResponse)
}

// @Summary GetRegionInformation
// @Description Getting information about the region
// @Tags Regions
// @Accept json
// @Produce json
// @Param id path int true "Region ID"
// @Success 200 {object} dto.RegionDTO
// @Failure 400 {object} ErrResponse
// @Router /regions/{id} [get]
func (h *Handler) GetRegionInformation(w http.ResponseWriter, r *http.Request) {
	var (
		err       error
		regionID  int
		regionDTO dto.RegionDTO
	)
	regionID, err = strconv.Atoi(chi.URLParam(r, "regionID"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	regionDTO, err = h.regionsService.GetInformationRegion(regionID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, regionDTO)
}

// @Summary GetRegionsInformation
// @Description Getting information about regions
// @Tags Regions
// @Accept json
// @Produce json
// @Param collection query ListInput false "string collection" collectionFormat(multi)
// @Success 200 {object} dto.RegionsDTO
// @Router /regions/ [get]
func (h *Handler) GetRegionsInformation(w http.ResponseWriter, r *http.Request) {
	var (
		regionsDTO dto.RegionsDTO
	)
	regionsDTO = h.regionsService.GetInformationRegions()
	input := r.Context().Value(httpin.Input).(*ListInput)
	if input.ID != nil {
		regionsDTO = regionsDTO.SelectByID(input.ID...)
	}
	if input.SortBy != "" {
		regionsDTO.SortByField(input.SortBy)
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, regionsDTO)
}

// @Summary RegionDelete
// @Description Delete region information
// @Tags Regions
// @Accept json
// @Produce json
// @Param requestDelete body RequestDelete true "json delete targetID Region"
// @Success 200 {object} dto.RegionDTO
// @Failure 400 {object} ErrResponse
// @Router /regions/ [delete]
func (h *Handler) RegionDelete(w http.ResponseWriter, r *http.Request) {
	var (
		deleteCityID int
		err          error
		dataRequest  RequestDelete
		dataResponse dto.RegionDTO
	)
	render.Decode(r, &dataRequest)
	deleteCityID = dataRequest.TargetID
	dataResponse, err = h.regionsService.GetInformationRegion(dataRequest.TargetID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	_, err = h.regionsService.DeleteRegion(deleteCityID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, dataResponse)
}
