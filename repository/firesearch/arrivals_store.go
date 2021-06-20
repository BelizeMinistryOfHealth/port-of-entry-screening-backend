package firesearch

import (
	"bz.moh.epi/poebackend/models"
	"context"
	"fmt"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
)

// ArrivalsStore is the firesearch store that keeps statistics on arrival numbers
type ArrivalsStore struct {
	Service Service
}

// PutDoc saves an arrival stat in the arrivals index
func (c *ArrivalsStore) PutDoc(ctx context.Context, arrival models.ArrivalStat) error {
	doc := firesearch.Doc{
		ID: arrival.ID,
		SearchFields: []firesearch.SearchField{
			{
				Key:   "date",
				Value: arrival.Date,
				Store: true,
			},
			{
				Key:   "year",
				Value: fmt.Sprintf("%d", arrival.Year),
				Store: true,
			},
			{
				Key:   "portOfEntry",
				Value: arrival.PortOfEntry,
				Store: true,
			},
		},
		Fields: []firesearch.Field{
			{
				Key:   "id",
				Value: arrival.ID,
			},
			{
				Key:   "countryOfEmbarkation",
				Value: arrival.CountryOfEmbarkation,
			},
			{
				Key:   "purposeOfTrip",
				Value: arrival.PurposeOfTrip,
			},
		},
	}
	req := firesearch.PutDocRequest{
		IndexPath: c.Service.IndexPath,
		Doc:       doc,
	}
	_, err := c.Service.IndexService.PutDoc(ctx, req)
	if err != nil {
		return fmt.Errorf("failure saving to arrival stat index: %w", err)
	}
	return nil
}

func (c *ArrivalsStore) getDoc(ctx context.Context, ID, accessKey string) (models.ArrivalStat, error) {
	req := firesearch.SearchRequest{Query: firesearch.SearchQuery{
		IndexPath: c.Service.IndexPath,
		AccessKey: accessKey,
		Limit:     1,
		Text:      ID,
		Filters:   nil,
		Select: []string{
			"id",
			"year",
			"date",
			"portOfEntry",
			"countryOfEmbarkation",
			"purposeOfTrip",
		},
		SearchFields: []string{"id"},
	}}

	searchResp, err := c.Service.IndexService.Search(ctx, req)
	if err != nil {
		return models.ArrivalStat{}, fmt.Errorf("failed to retrieve a document: %w", err)
	}
	hits := searchResp.Hits
	var arrivalStat models.ArrivalStat
	for _, h := range hits {
		fields := h.Fields
		arrivalStat = models.ArrivalStat{
			ID:                   GetField(fields, "id").(string),
			Year:                 GetField(fields, "year").(int),
			Date:                 GetField(fields, "date").(string),
			PortOfEntry:          GetField(fields, "portOfEntry").(string),
			CountryOfEmbarkation: GetField(fields, "countryOfEmbarkation").(string),
			PurposeOfTrip:        GetField(fields, "purposeOfTrip").(string),
		}
	}
	return arrivalStat, nil

}

// DeleteDoc deletes an arrival document from the arrival stat index
func (c *ArrivalsStore) DeleteDoc(ctx context.Context, ID string) error {
	deleteDocReq := firesearch.DeleteDocRequest{
		IndexPath: c.Service.IndexPath,
		ID:        ID,
	}
	_, err := c.Service.IndexService.DeleteDoc(ctx, deleteDocReq)
	if err != nil {
		return fmt.Errorf("DeleteDoc() failed: %w", err)
	}
	return nil
}
