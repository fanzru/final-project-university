package converter

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CtxToInt64(ctx echo.Context, key string) (int64, error) {
	v := ctx.Get(key)
	s := fmt.Sprintf("%v", v)
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}
