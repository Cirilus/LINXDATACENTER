package main

import (
	"LINXDATACENTER/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
)

func main() {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		logrus.Errorf("Error opening the file, err= %s", err)
		return
	}

	defer file.Close()

	var products []model.Product

	if filePath[len(filePath)-4:] == "json" {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&products)
		if err != nil {
			logrus.Errorf("Something going wrong in decode json, err= %s", err)
		}
	} else if filePath[len(filePath)-3:] == "csv" {
		reader := csv.NewReader(file)
		for {
			read, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				logrus.Errorf("Something going wrong in read csv, err= %s", err)
				return
			}

			price, err := strconv.ParseFloat(read[1], 64)
			if err != nil {
				logrus.Errorf("Cannot parse price, err= %s", err)
				return
			}
			rating, err := strconv.ParseFloat(read[2], 64)
			if err != nil {
				logrus.Errorf("Cannot price the ratign, err= %s", err)
				return
			}

			product := model.Product{
				Name:   read[0],
				Price:  price,
				Rating: rating,
			}
			products = append(products, product)
		}
	} else {
		logrus.Error("The file must have format json or csv")
		return
	}
	var highestPriceProduct, highestRatingProduct model.Product
	highestPrice := -1.0
	highestRating := -1.0

	for _, product := range products {
		if product.Price > highestPrice {
			highestPrice = product.Price
			highestPriceProduct = product
		}
		if product.Rating > highestRating {
			highestRating = product.Rating
			highestRatingProduct = product
		}
	}

	fmt.Printf("The most expensive product: %s, Its price: %.2f\n", highestPriceProduct.Name, highestPrice)
	fmt.Printf("The biggest rating product: %s, Its rating: %.2f\n", highestRatingProduct.Name, highestRating)
}
