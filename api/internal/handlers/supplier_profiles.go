package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"receipt-wrangler/api/internal/commands"
	"receipt-wrangler/api/internal/constants"
	"receipt-wrangler/api/internal/permissions"
	"receipt-wrangler/api/internal/services"
	"receipt-wrangler/api/internal/structs"
	"receipt-wrangler/api/internal/utils"
)

func ListSupplierProfiles(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")

	handler := structs.Handler{
		ErrorMessage:     "Error retrieving supplier profiles",
		Writer:           w,
		Request:          r,
		GroupId:          groupId,
		GroupPermissions: []string{permissions.GroupReceiptsCreate},
		ResponseType:     constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			uintGroupId, err := utils.StringToUint(groupId)
			if err != nil {
				return http.StatusBadRequest, err
			}

			token := structs.GetClaims(r)
			profiles, err := services.NewSupplierProfileService(nil).List(token.UserId, uintGroupId)
			if err != nil {
				return http.StatusInternalServerError, err
			}

			bytes, err := utils.MarshalResponseData(profiles)
			if err != nil {
				return http.StatusInternalServerError, err
			}

			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
			return 0, nil
		},
	}

	HandleRequest(handler)
}

func GetSupplierProfile(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")

	handler := structs.Handler{
		ErrorMessage:     "Error retrieving supplier profile",
		Writer:           w,
		Request:          r,
		GroupId:          groupId,
		GroupPermissions: []string{permissions.GroupReceiptsCreate},
		ResponseType:     constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			return writeSupplierProfile(w, r, groupId, func(userId uint, uintGroupId uint, profileId uint) (any, error) {
				return services.NewSupplierProfileService(nil).Get(userId, uintGroupId, profileId)
			})
		},
	}

	HandleRequest(handler)
}

func CreateSupplierProfile(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")

	handler := structs.Handler{
		ErrorMessage:     "Error creating supplier profile",
		Writer:           w,
		Request:          r,
		GroupId:          groupId,
		GroupPermissions: []string{permissions.GroupReceiptsUpdate},
		ResponseType:     constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			command := commands.UpsertSupplierProfileCommand{}
			err := command.LoadDataFromRequest(w, r)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			vErr := command.Validate()
			if len(vErr.Errors) > 0 {
				structs.WriteValidatorErrorResponse(w, vErr, http.StatusBadRequest)
				return 0, nil
			}

			uintGroupId, err := utils.StringToUint(groupId)
			if err != nil {
				return http.StatusBadRequest, err
			}

			token := structs.GetClaims(r)
			profile, serviceErr, err := services.NewSupplierProfileService(nil).Create(token.UserId, uintGroupId, command)
			return writeSupplierMutation(w, profile, serviceErr, err)
		},
	}

	HandleRequest(handler)
}

func UpdateSupplierProfile(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")

	handler := structs.Handler{
		ErrorMessage:     "Error updating supplier profile",
		Writer:           w,
		Request:          r,
		GroupId:          groupId,
		GroupPermissions: []string{permissions.GroupReceiptsUpdate},
		ResponseType:     constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			command := commands.UpsertSupplierProfileCommand{}
			err := command.LoadDataFromRequest(w, r)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			vErr := command.Validate()
			if len(vErr.Errors) > 0 {
				structs.WriteValidatorErrorResponse(w, vErr, http.StatusBadRequest)
				return 0, nil
			}

			uintGroupId, err := utils.StringToUint(groupId)
			if err != nil {
				return http.StatusBadRequest, err
			}
			profileId, err := utils.StringToUint(chi.URLParam(r, "profileId"))
			if err != nil {
				return http.StatusBadRequest, err
			}

			token := structs.GetClaims(r)
			profile, serviceErr, err := services.NewSupplierProfileService(nil).Update(token.UserId, uintGroupId, profileId, command)
			return writeSupplierMutation(w, profile, serviceErr, err)
		},
	}

	HandleRequest(handler)
}

func DeleteSupplierProfile(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")

	handler := structs.Handler{
		ErrorMessage:     "Error deleting supplier profile",
		Writer:           w,
		Request:          r,
		GroupId:          groupId,
		GroupPermissions: []string{permissions.GroupReceiptsUpdate},
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			uintGroupId, err := utils.StringToUint(groupId)
			if err != nil {
				return http.StatusBadRequest, err
			}
			profileId, err := utils.StringToUint(chi.URLParam(r, "profileId"))
			if err != nil {
				return http.StatusBadRequest, err
			}

			err = services.NewSupplierProfileService(nil).Delete(uintGroupId, profileId)
			if errors.Is(err, services.ErrSupplierProfileNotFound) {
				return http.StatusNotFound, err
			}
			if err != nil {
				return http.StatusInternalServerError, err
			}

			w.WriteHeader(http.StatusOK)
			return 0, nil
		},
	}

	HandleRequest(handler)
}

func ResolveSupplierProfile(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "groupId")

	handler := structs.Handler{
		ErrorMessage:     "Error resolving supplier profile",
		Writer:           w,
		Request:          r,
		GroupId:          groupId,
		GroupPermissions: []string{permissions.GroupReceiptsCreate},
		ResponseType:     constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			command := commands.ResolveSupplierProfileCommand{}
			err := command.LoadDataFromRequest(w, r)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			vErr := command.Validate()
			if len(vErr.Errors) > 0 {
				structs.WriteValidatorErrorResponse(w, vErr, http.StatusBadRequest)
				return 0, nil
			}

			uintGroupId, err := utils.StringToUint(groupId)
			if err != nil {
				return http.StatusBadRequest, err
			}

			token := structs.GetClaims(r)
			profile, err := services.NewSupplierProfileService(nil).Resolve(token.UserId, uintGroupId, command.Name)
			if err != nil {
				return http.StatusInternalServerError, err
			}

			response := commands.ResolveSupplierProfileResponse{Profile: profile}
			bytes, err := utils.MarshalResponseData(response)
			if err != nil {
				return http.StatusInternalServerError, err
			}

			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
			return 0, nil
		},
	}

	HandleRequest(handler)
}

func writeSupplierMutation(w http.ResponseWriter, profile any, vErr structs.ValidatorError, err error) (int, error) {
	if len(vErr.Errors) > 0 {
		structs.WriteValidatorErrorResponse(w, vErr, http.StatusBadRequest)
		return 0, nil
	}
	if errors.Is(err, services.ErrSupplierProfileNotFound) {
		return http.StatusNotFound, err
	}
	if errors.Is(err, services.ErrSupplierCategoryTagMissing) {
		structs.WriteValidatorErrorResponse(w, structs.ValidatorError{Errors: map[string]string{
			"defaults": "One or more categories or tags do not exist",
		}}, http.StatusBadRequest)
		return 0, nil
	}
	if errors.Is(err, services.ErrSupplierCategoryTagDenied) {
		utils.WriteCustomErrorResponse(w, "One or more categories or tags are not permitted", http.StatusForbidden)
		return 0, nil
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	bytes, err := utils.MarshalResponseData(profile)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return 0, nil
}

func writeSupplierProfile(w http.ResponseWriter, r *http.Request, groupId string, load func(uint, uint, uint) (any, error)) (int, error) {
	uintGroupId, err := utils.StringToUint(groupId)
	if err != nil {
		return http.StatusBadRequest, err
	}
	profileId, err := utils.StringToUint(chi.URLParam(r, "profileId"))
	if err != nil {
		return http.StatusBadRequest, err
	}

	token := structs.GetClaims(r)
	profile, err := load(token.UserId, uintGroupId, profileId)
	if errors.Is(err, services.ErrSupplierProfileNotFound) {
		return http.StatusNotFound, err
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}

	bytes, err := utils.MarshalResponseData(profile)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return 0, nil
}
