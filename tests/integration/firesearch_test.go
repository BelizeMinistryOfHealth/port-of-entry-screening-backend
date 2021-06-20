package integration

import (
	firesearch2 "bz.moh.epi/poebackend/repository/firesearch"
	"context"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
	"testing"
	"time"
)

func TestCreateIndex(t *testing.T) {
	ctx := context.Background()
	firesearchService := firesearch2.CreateFiresearchService(
		"Persons Index",
		"persons_index",
		"PGIA")
	err := firesearchService.CreateIndex(ctx)
	if err != nil {
		t.Errorf("CreateIndex failed: %v", err)
	}
}

func TestDeleteIndex(t *testing.T) {
	ctx := context.Background()
	indexPath := "persons_index"
	deleteIndexReq := firesearch.DeleteIndexRequest{IndexPath: indexPath}
	firesearchService := firesearch2.CreateFiresearchService("Persons Index", indexPath, "PGIA")
	_, err := firesearchService.IndexService.DeleteIndex(ctx, deleteIndexReq)
	if err != nil {
		t.Errorf("failed to delete index: %v", err)
	}
}

func TestPutDoc(t *testing.T) {

	firesearchService := firesearch2.CreateFiresearchService(
		"Persons Index",
		"persons_index",
		"PGIA")
	ctx := context.Background()

	//dateOfArrival, _ := time.Parse("2006-01-02", "2021-04-12")
	//person := models.Person{
	//	ID: "1",
	//	PersonalInfo: models.PersonalInfo{
	//		FirstName: "Roberto",
	//		LastName:  "Guerra",
	//		FullName:  "Roberto Uris Guerra",
	//	},
	//	Arrival: models.Arrival{
	//		TripID: "11111",
	//		ArrivalInfo: models.ArrivalInfo{
	//			DateOfArrival: dateOfArrival,
	//		},
	//	},
	//	Created:     time.Now(),
	//	Modified:    nil,
	//	PortOfEntry: "PGIA",
	//}
	//err := firesearchService.PutDoc(ctx, person)
	//if err != nil {
	//	t.Fatalf("Failed to put document to firesearch: %v", err)
	//}
	//
	//dateOfArrival, _ = time.Parse("2006-01-02", "2021-05-11")
	//person2 := models.Person{
	//	ID: "1112",
	//	PersonalInfo: models.PersonalInfo{
	//		FirstName: "Jamie",
	//		LastName:  "Xu",
	//		FullName:  "Jamie Xu",
	//	},
	//	Arrival: models.Arrival{
	//		TripID: "1212111",
	//		ArrivalInfo: models.ArrivalInfo{
	//			DateOfArrival: dateOfArrival,
	//		},
	//	},
	//	Created:     time.Now(),
	//	Modified:    nil,
	//	PortOfEntry: "PGIA",
	//}
	//err = firesearchService.PutDoc(ctx, person2)
	//if err != nil {
	//	t.Fatalf("Failed to put document to firesearch: %v", err)
	//}
	//
	//dateOfArrival, _ = time.Parse("2006-01-02", "2021-05-11")
	//person3 := models.Person{
	//	ID: "1111",
	//	PersonalInfo: models.PersonalInfo{
	//		FirstName: "Bill",
	//		LastName:  "Wang",
	//		FullName:  "Bill Wang",
	//	},
	//	Arrival: models.Arrival{
	//		TripID: "1212111",
	//		ArrivalInfo: models.ArrivalInfo{
	//			DateOfArrival: dateOfArrival,
	//		},
	//	},
	//	Created:     time.Now(),
	//	Modified:    nil,
	//	PortOfEntry: "Western Border",
	//}
	//err = firesearchService.PutDoc(ctx, person3)
	//if err != nil {
	//	t.Fatalf("Failed to put document to firesearch: %v", err)
	//}

	accessKeyService := firesearch.NewAccessKeyService(firesearchService.Client)
	keyReq := firesearch.GenerateKeyRequest{IndexPathPrefix: "firesearch/indexes/test_index"}
	keyResp, err := accessKeyService.GenerateKey(ctx, keyReq)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}
	accessKey := keyResp.AccessKey
	dateSearch, _ := time.Parse("2006-01-02", "2021-05-01")
	searchResult, err := firesearchService.SearchByDate(ctx, dateSearch, accessKey, "2021")
	if err != nil {
		t.Errorf("Failed to search: %v", err)
	}
	t.Logf("search result: %v", searchResult)
}

func TestSearch(t *testing.T) {
	ctx := context.Background()
	indexPath := "firesearch/indexes/persons_index"
	keyReq := firesearch.GenerateKeyRequest{IndexPathPrefix: "firesearch/indexes/persons_index"}
	firesearchService := firesearch2.CreateFiresearchService(
		"Persons Index",
		indexPath,
		"PGIA")
	accessKeyService := firesearch.NewAccessKeyService(firesearchService.Client)
	keyResp, err := accessKeyService.GenerateKey(ctx, keyReq)
	if err != nil {
		t.Fatalf("GenerateKey failed: %v", err)
	}
	accessKey := keyResp.AccessKey
	searchRequest := firesearch.SearchRequest{
		Query: firesearch.SearchQuery{
			IndexPath: indexPath,
			AccessKey: accessKey,
			Limit:     100,
			Text:      "2021",
			Filters: []firesearch.Field{
				//{
				//	Key: "portOfEntry",
				//	Value: "PGIA",
				//},
				{
					Key:   "year",
					Value: "2021",
				},
				{
					Key:   "month",
					Value: "6",
				},
				{
					Key:   "day",
					Value: 20,
				},
			},
			Select:       []string{"year", "fullName", "month", "middleName", "day", "nationality", "portOfEntry"},
			SearchFields: []string{"year"},
			Cursor:       "",
		},
	}
	resp, err := firesearchService.IndexService.Search(ctx, searchRequest)
	if err != nil {
		t.Fatalf("error searching index: %v", err)
	}
	//t.Logf("Hits: %v", resp.Hits)
	t.Logf("TOTAL HITS: %d", len(resp.Hits))
	var withPorts [][]firesearch.Field
	for _, r := range resp.Hits {
		t.Logf("hit: %v", r.Fields)
		if hasPoe(r.Fields) {
			withPorts = append(withPorts, r.Fields)
		}
	}
	t.Logf("withPorts: %v", withPorts)

}

func hasPoe(fields []firesearch.Field) bool {
	has := false

	for _, f := range fields {
		if f.Key == "portOfEntry" && f.Value != nil {
			has = true
			return has
		}
	}

	return has
}
