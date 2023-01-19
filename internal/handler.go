package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func HandlerWhen(c *gin.Context) {
	var answer int
	var answerStr string

	//parse request path
	year, err := strconv.Atoi(c.Params.ByName("year"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid year")
		return
	}
	timeReq := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	//calculate difference
	if timeReq.After(time.Now()) {
		answer = (timeReq.Year() - time.Now().Year()) * 365
		answerStr = "Day left: " + strconv.Itoa(answer)

		//answer = time.Until(timer)
		//Until and Since return struct Duration
		//that can't be longer than 290 years

	} else {
		answer = (time.Now().Year() - timeReq.Year()) * 365
		answerStr = "Day gone: " + strconv.Itoa(answer)
		//answer = time.Since(timer)
	}

	c.String(http.StatusOK, answerStr)
}
