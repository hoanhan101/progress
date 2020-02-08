package handler

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

func getParams(c echo.Context, names map[string]bool) (map[string]string, error) {
	values := map[string]string{}

	for name, required := range names {
		v := c.Param(name)

		if v == "" && required == true {
			return nil, errors.New(fmt.Sprintf("`%v` parameter is not specified in the URI", name))
		}

		values[name] = v
	}

	return values, nil
}

func getFormValues(c echo.Context, names map[string]bool) (map[string]string, error) {
	values := map[string]string{}

	for name, required := range names {
		v := c.FormValue(name)

		if v == "" && required == true {
			return nil, errors.New(fmt.Sprintf("`%v` value is not specified in the request body", name))
		}

		values[name] = v
	}

	return values, nil
}
