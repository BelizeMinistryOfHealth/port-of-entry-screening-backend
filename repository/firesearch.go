package repository

import (
	"bz.moh.epi/poebackend/models"
	"context"
	"fmt"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
	"os"
)

// FiresearchService is an instance of a service that allows us to do operations on Firesearch.
type FiresearchService struct {
	IndexName    string
	IndexPath    string
	PortOfEntry  string
	IndexService *firesearch.IndexService
	Client       *firesearch.Client
}

// CreateFiresearchService creates an instance of the FiresearchService
func CreateFiresearchService(indexName, indexPath, portOfEntry string) FiresearchService {
	host := os.Getenv("FIRESEARCH_HOST")
	api := os.Getenv("FIRESEARCH_API_KEY")
	client := firesearch.NewClient(
		host,
		api,
	)
	indexService := firesearch.NewIndexService(client)
	return FiresearchService{
		IndexName:    indexName,
		IndexPath:    fmt.Sprintf("firesearch/indexes/%s", indexPath),
		PortOfEntry:  portOfEntry,
		IndexService: indexService,
		Client:       client,
	}
}

// CreateIndex creates an index in Firesearch
func (f FiresearchService) CreateIndex(ctx context.Context) error {
	createIndexReq := firesearch.CreateIndexRequest{
		Index: firesearch.Index{
			IndexPath:     f.IndexPath,
			Name:          f.IndexName,
			Language:      "english",
			KeepStopWords: false,
			CaseSensitive: false,
			NoStem:        false,
		},
	}
	_, err := f.IndexService.CreateIndex(ctx, createIndexReq)
	if err != nil {
		return fmt.Errorf("firesearch: IndexService.CreateIndex %w", err)
	}
	return nil
}

// PutDoc creates a document in the Firesearch index
func (f FiresearchService) PutDoc(ctx context.Context, person models.Person) error {
	personalInfo := person.PersonalInfo
	putDocReq := firesearch.PutDocRequest{
		IndexPath: f.IndexPath,
		Doc: firesearch.Doc{
			ID: person.ID,
			SearchFields: []firesearch.SearchField{
				{
					Key:   "firstName",
					Value: personalInfo.FirstName,
					Store: true,
				},
				{
					Key:   "lastName",
					Value: personalInfo.LastName,
					Store: true,
				},
				{
					Key:   "fullName",
					Value: personalInfo.FullName,
					Store: true,
				},
				{
					Key:   "tripID",
					Value: person.Arrival.TripID,
					Store: true,
				},
				{
					Key:   "portOfEntry",
					Value: person.PortOfEntry,
					Store: true,
				},
				{
					Key:   "dateOfArrival",
					Value: person.Arrival.ArrivalInfo.DateOfArrival.Format("2006-01-02"),
					Store: true,
				},
			},
			Fields: []firesearch.Field{
				{
					Key:   "ID",
					Value: person.ID,
				},
				{
					Key:   "personalInfo",
					Value: personalInfo,
				},
				{
					Key:   "arrivalInfo",
					Value: person.Arrival.ArrivalInfo,
				},
			},
		},
	}

	_, err := f.IndexService.PutDoc(ctx, putDocReq)
	if err != nil {
		return fmt.Errorf("firesearch: PutDoc error: %w", err)
	}
	return nil
}

// SearchDocs searches for a document that has the specified searchText
func (f FiresearchService) SearchDocs(ctx context.Context, accessKey, searchText string) ([]firesearch.SearchResult, error) {
	searchReq := firesearch.SearchRequest{
		Query: firesearch.SearchQuery{
			IndexPath: f.IndexPath,
			AccessKey: accessKey,
			Limit:     50,
			Text:      searchText,
			Filters: []firesearch.Field{
				{
					Key:   "portOfEntry",
					Value: f.PortOfEntry,
				},
			},
			Select: []string{
				"ID",
				"personalInfo",
				"arrivalInfo",
				"tripID",
				"portOfEntry",
				"dateOfArrival",
			},
			SearchFields: []string{},
		},
	}
	searchResp, err := f.IndexService.Search(ctx, searchReq)
	if err != nil {
		return []firesearch.SearchResult{}, fmt.Errorf("firesearch: Search failed: %w", err)
	}

	return searchResp.Hits, nil
}
