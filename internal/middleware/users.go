package middleware

import (
	"regexp"
	"todoapp/internal/handlers"
	"todoapp/internal/models"
	"todoapp/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		validates := validator.New()

		validates.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			pass := fl.Field().String()
			hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pass)
			hasNumber := regexp.MustCompile(`[0-9]`).MatchString(pass)
			hasSymbol := regexp.MustCompile(`[!@#~$%^&*()+|_.,<>?/{}\-]`).MatchString(pass)
			return len(pass) >= 8 && hasUpper && hasNumber && hasSymbol
		})

		user := new(models.User)
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, handlers.Response{Message: "Invalid user", Errors: err.Error()})
		}
		if err := validates.Struct(user); err != nil {
			code := int(0)
			errorsMap := map[string]string{}
			for _, e := range err.(validator.ValidationErrors) {
				code, errorsMap = utils.ValidationErrors(e, user)
			}
			return c.JSON(code, handlers.Response{
				Message: "validation failed",
				Errors:  errorsMap,
			})
		}

		c.Set("user", user)

		return next(c)
	}

}
