package integration

import (
	firesearch2 "bz.moh.epi/poebackend/repository/firesearch"
	"context"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	ctx := context.Background()
	firesearchService := firesearch2.CreateFiresearchService(
		"Test Index",
		"test_index",
		"PGIA")
	err := firesearchService.CreateIndex(ctx)
	if err != nil {
		t.Errorf("CreateIndex failed: %v", err)
	}
}

func TestPutDoc(t *testing.T) {
	//host := os.Getenv("FIRESEARCH_HOST")
	//api := os.Getenv("FIRESEARCH_API_KEY")
	//client := firesearch.NewClient(
	//	host,
	//	api,
	//)
	//indexService := firesearch.NewIndexService(client)
	firesearchService := firesearch2.CreateFiresearchService(
		"Test Index",
		"test_index",
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
	//err := repository.PutDoc(ctx, indexService, "test_index", person)
	//if err != nil {
	//	t.Fatalf("Failed to put document to firesearch: %v", err)
	//}
	//
	//dateOfArrival, _ := time.Parse("2006-01-02", "2021-05-11")
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
	//err := firesearchService.PutDoc(ctx, person2)
	//if err != nil {
	//	t.Fatalf("Failed to put document to firesearch: %v", err)
	//}

	//dateOfArrival, _ := time.Parse("2006-01-02", "2021-05-11")
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
	//err := repository.PutDoc(ctx, indexService, "test_index", person3)
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
	searchResult, err := firesearchService.SearchDocs(ctx, accessKey, "202105")
	if err != nil {
		t.Errorf("Failed to search: %v", err)
	}
	t.Logf("search result: %v", searchResult)
}
