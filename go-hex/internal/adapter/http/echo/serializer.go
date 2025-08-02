package echo

import (
	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
)

type JsonSerializer struct{}

func (s *JsonSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := sonic.ConfigDefault.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func (s *JsonSerializer) Deserialize(c echo.Context, i interface{}) error {
	return sonic.ConfigDefault.NewDecoder(c.Request().Body).Decode(i)
}
