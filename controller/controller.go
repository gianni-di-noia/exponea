package controller

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	"exponea.com/core"
	"github.com/labstack/echo/v4"
)

// BuildWorks does the requstes and build the response
func BuildWorks(c echo.Context) ([]core.Work, error) {
	start := time.Now()
	works := []core.Work{}
	queue := make(chan core.Work, core.Count)
	var wg sync.WaitGroup

	timeout, err := parseTimeout(c)
	if err != nil {
		return nil, c.JSON(500, core.Message{Message: "timeout must be an integer"})
	}

	for i := 1; i <= core.Count; i++ {
		wg.Add(1)
		go fetchWork(queue, &wg)
	}
	wg.Wait()
	close(queue)

	for elem := range queue {
		log.Printf("elem %d", elem)
		if elem.Time == 0 {
			return nil, c.JSON(500, core.Message{Message: "one of requests failed"})
		}
		works = append(works, elem)
	}

	end := time.Now()
	usedTime := end.Sub(start).Milliseconds()
	if usedTime > timeout {
		return nil, c.JSON(500, core.Message{Message: "timeout exceeded"})
	}
	log.Printf("USED %v seconds\n", usedTime)
	return works, nil
}

func fetchWork(queue chan core.Work, wg *sync.WaitGroup) {
	defer wg.Done()
	body, _ := core.MyRequest(core.ExponeaURL, "GET", "")
	if body == nil {
		queue <- core.Work{Time: 0}
	} else {
		var data core.Work
		err := json.Unmarshal(body, &data)
		if err != nil {
			log.Fatal(err)
		}
		queue <- data
	}
}

func parseTimeout(c echo.Context) (int64, error) {
	timeoutQ := c.QueryParam("timeout")
	log.Printf("timeout %s", timeoutQ)
	timeout, err := strconv.ParseInt(timeoutQ, 10, 64)
	if err != nil {
		return 0, err
	}
	return timeout, nil
}
