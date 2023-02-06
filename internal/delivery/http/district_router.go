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

func (h *Handler) initDistrictRoutes() chi.Router {
	r := chi.NewRouter()
	r.Route("/{districtID}", func(r chi.Router) {
		r.Get("/", h.GetDistrictInformation)
		r.Put("/", h.UpdateDistrict)
	})
	r.With(httpin.NewInput(ListInput{})).Get("/", h.GetDistrictsInformation)
	r.Delete("/", h.DistrictDelete)
	r.Route("/create", func(r chi.Router) {
		r.Post("/", h.CreateDistrict)
	})
	return r
}

// @Summary CreateDistrict
// @Description Сreating a new district entry
// @Tags Districts
// @Accept json
// @Produce json
// @Param input body dto.DistrictDTO true "json information District"
// @Success 200 {object} dto.DistrictDTO
// @Failure 400 {object} ErrResponse
// @Router /districts/create [post]
func (h *Handler) CreateDistrict(w http.ResponseWriter, r *http.Request) {
	err := error(nil)
	districtCreate := dto.DistrictDTO{}
	if err = render.Decode(r, &districtCreate); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err = h.districtsService.CreateDistrict(districtCreate); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, districtCreate)
}

func (h *Handler) UpdateDistrictID(districtID, newID int) (int, string, error) {
	err := error(nil)
	districtNewID := dto.DistrictDTO{}
	messageResponse := ""
	if districtNewID, err = h.districtsService.GetInformationDistrict(newID); err != nil {
		err = h.districtsService.UpdateIDDistrict(districtID, newID)
		if err != nil {
			messageResponse += fmt.Sprintf("District ID: %s\n", err.Error())
		} else {
			messageResponse += fmt.Sprintf("District ID = %d OldID = %d \n", newID, districtID)
			districtID = newID
		}
	} else {
		messageResponse += fmt.Sprintf("District with ID = %d "+
			"Name = %s already exists\n", districtNewID.ID, districtNewID.Name)
		err = fmt.Errorf(messageResponse)
	}
	return districtID, messageResponse, err
}

func (h *Handler) UpdateDistrictName(districtID int, newName, name string) string {
	err := error(nil)
	messageResponse := ""
	err = h.districtsService.UpdateNameDistrict(districtID, newName)
	if err != nil {
		messageResponse += fmt.Sprintf("Name: %s\n", err.Error())
	} else {
		messageResponse += fmt.Sprintf("Name = %s OldName = %s\n", newName, name)
	}
	return messageResponse
}

// @Summary UpdateDistrict
// @Description District information update
// @Tags Districts
// @Accept json
// @Produce html
// @Param id path int true "District ID"
// @Param updateDistrictRequest body UpdateRequest true "json update information District"
// @Success 200 {string} string
// @Failure 400 {object} ErrResponse
// @Router /districts/{id} [put]
func (h *Handler) UpdateDistrict(w http.ResponseWriter, r *http.Request) {
	var (
		err             error
		districtID      int
		districtDTO     dto.DistrictDTO
		dataRequest     UpdateRequest
		messageResponse string = "Updated:\n"
	)
	districtID, err = strconv.Atoi(chi.URLParam(r, "districtID"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if districtDTO, err = h.districtsService.GetInformationDistrict(districtID); err != nil {
		messageResponse += fmt.Sprintf("District: %s\n", err.Error())
	} else {
		render.Decode(r, &dataRequest)
		if dataRequest.NewID != 0 {
			messageResponseCityID := ""
			districtID, messageResponseCityID, err = h.UpdateDistrictID(districtID, dataRequest.NewID)
			messageResponse += messageResponseCityID
			if err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}
		}
		if dataRequest.NewName != "" {
			messageResponseName := h.UpdateDistrictName(districtID, dataRequest.NewName, districtDTO.Name)
			messageResponse += messageResponseName
		}
	}
	render.Status(r, http.StatusOK)
	render.HTML(w, r, messageResponse)
}

// @Summary GetDistrictInformation
// @Description Getting information about the district
// @Tags Districts
// @Accept json
// @Produce json
// @Param id path int true "District ID"
// @Success 200 {object} dto.DistrictDTO
// @Failure 400 {object} ErrResponse
// @Router /districts/{id} [get]
func (h *Handler) GetDistrictInformation(w http.ResponseWriter, r *http.Request) {
	var (
		err         error
		districtID  int
		districtDTO dto.DistrictDTO
	)
	districtID, err = strconv.Atoi(chi.URLParam(r, "districtID"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	districtDTO, err = h.districtsService.GetInformationDistrict(districtID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, districtDTO)
}

// @Summary GetDistrictsInformation
// @Description Getting information about districts
// @Tags Districts
// @Accept json
// @Produce json
// @Param collection query ListInput false "string collection" collectionFormat(multi)
// @Success 200 {object} dto.DistrictsDTO
// @Router /districts/ [get]
func (h *Handler) GetDistrictsInformation(w http.ResponseWriter, r *http.Request) {
	var (
		districtsDTO dto.DistrictsDTO
	)
	districtsDTO = h.districtsService.GetInformationDistricts()
	input := r.Context().Value(httpin.Input).(*ListInput)
	if input.ID != nil {
		districtsDTO = districtsDTO.SelectByID(input.ID...)
	}
	if input.Name != nil {
		districtsDTO = districtsDTO.SelectByName(input.Name...)
	}
	if input.SortBy != "" {
		districtsDTO.SortByField(input.SortBy)
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, districtsDTO)
}

// @Summary DistrictDelete
// @Description Delete district information
// @Tags Districts
// @Accept json
// @Produce json
// @Param requestDelete body RequestDelete true "json delete targetID District"
// @Success 200 {object} dto.DistrictDTO
// @Failure 400 {object} ErrResponse
// @Router /districts/ [delete]
func (h *Handler) DistrictDelete(w http.ResponseWriter, r *http.Request) {
	var (
		deleteCityID int
		err          error
		dataRequest  RequestDelete
		dataResponse dto.DistrictDTO
	)
	render.Decode(r, &dataRequest)
	deleteCityID = dataRequest.TargetID
	dataResponse, err = h.districtsService.GetInformationDistrict(dataRequest.TargetID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	_, err = h.districtsService.DeleteDistrict(deleteCityID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, dataResponse)
}
