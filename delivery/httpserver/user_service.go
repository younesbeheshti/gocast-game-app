package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/younesbeheshti/gocast_game/dto"
	"github.com/younesbeheshti/gocast_game/pkg/httpmsg"
	"github.com/younesbeheshti/gocast_game/service/userservice"
	"net/http"
)

func (s Server) userRegisterHandler(c *echo.Context) error {

	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err, fieldErrors := s.userValidator.ValidateRegisterRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, map[string]interface{}{
			"message": msg,
			"error":   fieldErrors,
		})
	}

	resp, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

func (s Server) userLoginHandler(c *echo.Context) error {
	var req userservice.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println("req:", req)

	resp, err := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) userProfileHandler(c *echo.Context) error {

	authToken := c.Request().Header.Get("Authorization")
	claims, err := s.authSvc.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := s.userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)

}
