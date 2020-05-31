.PHONY: deploy delete

#gomodgen:
#	GO111MODULE=on go mod init

deploy:
	gcloud functions deploy arrivals --entry-point Arrivals --runtime go113 --trigger-http --env-vars-file env.yaml --region us-east1 --allow-unauthenticated --source .

deployQueryByPort:
	gcloud functions deploy byPortOfEntry --entry-point FindByPortOfEntry --runtime go113 --trigger-http --env-vars-file env.yaml --region us-east1--allow-unauthenticated --source .



delete:
	gcloud functions delete ingest --entry-point Ingest --runtime go113 --trigger-http
