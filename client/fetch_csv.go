package client

import (
	"encoding/csv"
	"net/http"
	"strconv"
	"time"

	"github.com/Jopoleon/AtlantTest/logger"

	"github.com/Jopoleon/AtlantTest/models"
)

type CSVFetchClient struct {
	log *logger.LocalLogger
	c   *http.Client
}

type CSVFetcher interface {
	GetCSV(url string) ([]models.Product, error)
}

func NewCSVFetchClient(log *logger.LocalLogger) CSVFetcher {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	return &CSVFetchClient{
		c:   client,
		log: log,
	}
}

func (c *CSVFetchClient) GetCSV(url string) ([]models.Product, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.log.Error(err)
		return nil, err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		c.log.Error(err)
		return nil, err
	}
	//PRODUCT NAME;PRICE.
	file := csv.NewReader(resp.Body)
	file.Comment = '#'
	file.Comma = ';'
	file.FieldsPerRecord = 2
	records, err := file.ReadAll()
	if err != nil {
		c.log.Error(err)
		return nil, err
	}
	var res []models.Product

	//skipping first line of csv with column names
	for _, record := range records[1:] {

		price, err := strconv.ParseFloat(record[1], 32)
		if err != nil {
			c.log.Error(err)
			return nil, err
		}
		p := models.Product{
			Name:      record[0],
			LastPrice: float32(price),
		}
		res = append(res, p)
	}
	return res, nil
}
