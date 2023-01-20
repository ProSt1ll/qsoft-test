package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func HandlerWhen(c *gin.Context) {
	//parse request path
	year, err := strconv.Atoi(c.Params.ByName("year"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid year")
		return
	}
	timeReq := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	//calculate difference
	var answer int
	var answerStr string
	if timeReq.After(time.Now()) {
		answer = (timeReq.Year() - time.Now().Year()) * 365
		answerStr = fmt.Sprintf("Day left: %s", strconv.Itoa(answer))
		//Until and Since from pkg time returns struct Duration
		//that can't be longer than 290 years
	} else {
		answer = (time.Now().Year() - timeReq.Year()) * 365
		answerStr = fmt.Sprintf("Day gone: %s", strconv.Itoa(answer))
	}

	c.String(http.StatusOK, answerStr)
}
