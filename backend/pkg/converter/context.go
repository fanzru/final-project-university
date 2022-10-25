package converter

import (
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CtxToInt64(ctx echo.Context, key string) (int64, error) {
	v := ctx.Get(key)
	log.Println("------------------------------ user id : ", v)
	s := fmt.Sprintf("%v", v)
	log.Println("------------------------------ string user id : ", s)
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	log.Println("------------------------------ int64 user id : ", n)
	return n, nil
}
