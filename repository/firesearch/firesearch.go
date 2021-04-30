package firesearch

import (
	"bz.moh.epi/poebackend/models"
	"context"
	"fmt"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
	"os"
	"time"
)

// Service is an instance of a service that allows us to do operations on Firesearch.
type Service struct {
	IndexName    string
	IndexPath    string
	PortOfEntry  string
	IndexService *firesearch.IndexService
	Client       *firesearch.Client
}

// CreateFiresearchService creates an instance of the Service
func CreateFiresearchService(indexName, indexPath, portOfEntry string) Service {
	host := os.Getenv("FIRESEARCH_HOST")
	api := os.Getenv("FIRESEARCH_API_KEY")
	client := firesearch.NewClient(
		host,
		api,
	)
	indexService := firesearch.NewIndexService(client)
	return Service{
		IndexName:    indexName,
		IndexPath:    fmt.Sprintf("firesearch/indexes/%s", indexPath),
		PortOfEntry:  portOfEntry,
		IndexService: indexService,
		Client:       client,
	}
}

// CreateIndex creates an index in Firesearch
func (f Service) CreateIndex(ctx context.Context) error {
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
func (f Service) PutDoc(ctx context.Context, person models.Person) error {
	personalInfo := person.PersonalInfo
	dateOfArrival := person.Arrival.ArrivalInfo.DateOfArrival
	yearOfArrival := dateOfArrival.Year()
	monthOfArrival := dateOfArrival.Month()
	dayOfArrival := dateOfArrival.Day()
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
				{
					Key:   "yearOfArrival",
					Value: fmt.Sprintf("%d", yearOfArrival),
					Store: true,
				},
				{
					Key:   "monthOfArrival",
					Value: fmt.Sprintf("%d", monthOfArrival),
					Store: true,
				},
				{
					Key:   "dayOfArrival",
					Value: fmt.Sprintf("%d", dayOfArrival),
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
func (f Service) SearchDocs(ctx context.Context, accessKey, searchText string) ([]firesearch.SearchResult, error) {
	filters := []firesearch.Field{
		{
			Key:   "portOfEntry",
			Value: f.PortOfEntry,
		},
		{
			Key:   "yearOfArrival",
			Value: "2021",
		},
		{
			Key:   "monthOfArrival",
			Value: "5",
		},
	}
	searchReq := firesearch.SearchRequest{
		Query: firesearch.SearchQuery{
			IndexPath: f.IndexPath,
			AccessKey: accessKey,
			Limit:     50,
			Text:      searchText,
			Filters:   filters,
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

// SearchByDate searches for a document that has the specified searchText
func (f Service) SearchByDate(ctx context.Context, date time.Time, accessKey, searchText string) ([]firesearch.SearchResult, error) {
	filters := []firesearch.Field{
		{
			Key:   "portOfEntry",
			Value: f.PortOfEntry,
		},
		{
			Key:   "yearOfArrival",
			Value: fmt.Sprintf("%d", date.Year()),
		},
		{
			Key:   "monthOfArrival",
			Value: fmt.Sprintf("%d", date.Month()),
		},
	}
	searchReq := firesearch.SearchRequest{
		Query: firesearch.SearchQuery{
			IndexPath: f.IndexPath,
			AccessKey: accessKey,
			Limit:     50,
			Text:      searchText,
			Filters:   filters,
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

// CreatePerson inserts a person into firesearch
func (f Service) CreatePerson(ctx context.Context, person models.Person) error {
	return f.PutDoc(ctx, person)
}

// UpdatePerson updates a person in firesearch
func (f Service) UpdatePerson(ctx context.Context, person models.Person) error {
	return f.PutDoc(ctx, person)
}
