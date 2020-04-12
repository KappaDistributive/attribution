# End-to-End Computation of GMV-based Shapley Values

```go
package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/KappaDistributive/attribution"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"sync"
)

const (
	bqProject = "marketing"
)

// dataRow represents a row in the raw data obtained from BigQuery.
type dataRow struct {
	ChannelPath    string
	SumGMV         float64
	SumTransaction int64
}

// String returns a simplified string represenation of a dataRow object.
func (row dataRow) String() string {
	return row.ChannelPath
}

// checkError provides naive error handling.
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// getContribution creates a Contribution object from a given dataRow object.
func getContribution(row dataRow, valueType string) attribution.Contribution {
	rawTouchpoints := strings.Split(row.ChannelPath, ">")
	touchpoints := make([]attribution.Touchpoint, len(rawTouchpoints))
	for index, rawTouchpoint := range rawTouchpoints {
		touchpoints[index] = attribution.Touchpoint{
			Name: strings.Trim(strings.ToLower(rawTouchpoint), " "),
		}
	}
	value := new(big.Float)
	switch valueType {
	case "gmv":
		value.SetFloat64(row.SumGMV)
	case "transaction":
		value.SetFloat64(float64(row.SumGMV))
	default:
		log.Fatalf("Unknown value type: %s", valueType)
	}
	return attribution.Contribution{
		Touchpoints: attribution.Touchpoints(touchpoints),
		Value:       *value,
	}
}

// getContributions transforms a given data table into a list of contributions.
func getContributions(rawData []dataRow, valueType string) []attribution.Contribution {
	var allContributions []attribution.Contribution

	for _, row := range rawData {
		allContributions = append(allContributions, getContribution(row, valueType))
	}

	return allContributions
}

// getQuery reads the relevant query string from a file.
func getQuery() string {
	data, err := ioutil.ReadFile("query.sql")
	checkError(err)

	query := string(data)

	return query
}

// getRawData reads the relevant raw data from BigQuery.
func getRawData(query string) []dataRow {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, bqProject)
	checkError(err)

	queryJob := client.Query(query)
	it, err := queryJob.Read(ctx)
	checkError(err)

	var rawData []dataRow

	for {
		var row dataRow
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		checkError(err)
		rawData = append(rawData, row)
	}

	return rawData
}

// gmvContributionSets transforms a list of Contribution objects into a list of corresponding ContributionSet objects.
func getContributionSets(gmvContributions []attribution.Contribution) []attribution.ContributionSet {
	var gmvContributionSets []attribution.ContributionSet
	for _, contribution := range gmvContributions {
		gmvContributionSets = append(gmvContributionSets, contribution.Set())
	}
	return gmvContributionSets
}

// main computes GMV-based Shapley values for all relevant marketing touchpoints.
func main() {
	query := getQuery()
	log.Printf("Query:\n%s\n", query)
	rawData := getRawData(query)
	log.Printf("Retrieved %d raw data rows.", len(rawData))
	gmvContributions := getContributions(rawData, "gmv")
	log.Printf("Retrieved %d GMV contributions.", len(gmvContributions))

	gmvContributionSets := getContributionSets(gmvContributions)
	touchpoints := attribution.GetAllTouchpoints(gmvContributionSets)
	log.Printf("%s", touchpoints)

	// compute Shapley values in parallel.
	var wg sync.WaitGroup
	for _, touchpoint := range touchpoints {
		wg.Add(1)
		go func(touchpoint attribution.Touchpoint, wg *sync.WaitGroup) {
			defer wg.Done()
			shapleyValue := attribution.GetShapleyValue(touchpoint, gmvContributionSets)
			log.Printf(
				"Shapley value for touchpoint %s wrt GMV: %s",
				touchpoint,
				shapleyValue.String())

		}(touchpoint, &wg)
	}
	wg.Wait()
}
```
