
deployPersonsHook:
	gcloud functions deploy PersonsHook --entry-point PersonsHook --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.create  --region us-east1 --trigger-resource "projects/epi-belize/databases/(default)/documents/persons/{pushId}" --source . --env-vars-file env.yaml

deletePersonsHook:
	gcloud functions delete PersonsHook

deployPersonDeletedListener:
	gcloud functions deploy PersonDeletedListener --entry-point PersonDeletedListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.delete  --region us-east1 --trigger-resource "projects/epi-belize/databases/(default)/documents/persons/{pushId}" --source . --env-vars-file env.yaml

deletePersonDeletedListener:
	gcloud functions delete PersonDeletedListener

deployPersonUpdatedListener:
	gcloud functions deploy PersonUpdatedListener --entry-point PersonUpdatedListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.update  --region us-east1 --trigger-resource "projects/epi-belize/databases/(default)/documents/persons/{pushId}" --source . --env-vars-file env.yaml

deletePersonUpdatedListener:
	gcloud functions delete PersonUpdatedListener

deployFiresearchAccessKey:
	gcloud alpha functions deploy FiresearchAccessKey --entry-point AccessKeyFn --runtime go113 --trigger-http --env-vars-file env.yaml --region us-east1 --allow-unauthenticated --security-level secure-always --source .

deleteFiresearchAccessKey:
	gcloud functions delete FiresearchAccessKey

deployFiresearchArrivalStatAccessKey:
	gcloud alpha functions deploy FiresearchArrivalsStatAccessKey --entry-point ArrivalsStatAccessKeyFn --runtime go113 --trigger-http --env-vars-file env.yaml --region us-east1 --allow-unauthenticated --security-level secure-always --source .

deployScreeningCreatedListener:
	gcloud functions deploy ScreeningCreatedListener --entry-point ScreeningListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.create  --region us-east1 --trigger-resource "projects/epi-belize/databases/(default)/documents/screenings/{pushId}" --source . --env-vars-file env.yaml

deleteScreeningCreatedListener:
	gcloud functions delete ScreeningCreatedListener

deployScreeningUpdatedListener:
	gcloud functions deploy ScreeningUpdatedListener --entry-point ScreeningListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.update  --region us-east1 --trigger-resource "projects/epi-belize/databases/(default)/documents/screenings/{pushId}" --source . --env-vars-file env.yaml

deleteScreeningUpdatedListener:
	gcloud functions delete ScreeningUpdatedListener

deployRegistration:
	gcloud alpha functions deploy RegistrationFn --entry-point RegistrationFn --runtime go113 --trigger-http --env-vars-file env.yaml --region us-east1 --allow-unauthenticated --security-level secure-always --source .

deleteRegistration:
	gcloud functions delete RegistrationFn



############### Staging ######################
deployStagingHelloPubSub:
	gcloud functions deploy HelloPubSub --runtime go113 --trigger-topic arrival_created

deleteStagingHelloPubSub:
	gcloud functions delete HelloPubSub

deployStagingPersonsHook:
	gcloud functions deploy PersonsHook --entry-point PersonsHook --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.create  --region us-east1 --trigger-resource "projects/epi-belize-staging/databases/(default)/documents/persons/{pushId}" --source . --env-vars-file env.yaml

deleteStagingPersonsHook:
	gcloud functions delete PersonsHook

deployStagingPersonDeletedListener:
	gcloud functions deploy PersonDeletedListener --entry-point PersonDeletedListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.delete  --region us-east1 --trigger-resource "projects/epi-belize-staging/databases/(default)/documents/persons/{pushId}" --source . --env-vars-file env.staging.yaml

deleteStagingPersonDeletedListener:
	gcloud functions delete PersonDeletedListener

deployStagingPersonUpdatedListener:
	gcloud functions deploy PersonUpdatedListener --entry-point PersonUpdatedListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.update  --region us-east1 --trigger-resource "projects/epi-belize-staging/databases/(default)/documents/persons/{pushId}" --source . --env-vars-file env.yaml

deleteStagingPersonUpdatedListener:
	gcloud functions delete PersonUpdatedListener

deployStagingFiresearchAccessKey:
	gcloud alpha functions deploy FiresearchAccessKey --entry-point AccessKeyFn --runtime go113 --trigger-http --env-vars-file env.yaml --region us-east1 --allow-unauthenticated --security-level secure-always --source .

deleteStagingFiresearchAccessKey:
	gcloud functions delete FiresearchAccessKey

deployStagingScreeningCreatedListener:
	gcloud functions deploy ScreeningCreatedListener --entry-point ScreeningListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.create  --region us-east1 --trigger-resource "projects/epi-belize-staging/databases/(default)/documents/screenings/{pushId}" --source . --env-vars-file env.yaml

deleteStagingScreeningCreatedListener:
	gcloud functions delete ScreeningCreatedListener

deployStagingScreeningUpdatedListener:
	gcloud functions deploy ScreeningUpdatedListener --entry-point ScreeningListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.update  --region us-east1 --trigger-resource "projects/epi-belize-staging/databases/(default)/documents/screenings/{pushId}" --source . --env-vars-file env.yaml

deleteStagingScreeningUpdatedListener:
	gcloud functions delete ScreeningUpdatedListener

deployStagingRegistration:
	gcloud alpha functions deploy RegistrationFn --entry-point RegistrationFn --runtime go113 --trigger-http --env-vars-file env.staging.yaml --region us-east1 --allow-unauthenticated --security-level secure-always --source .

deleteStagingRegistration:
	gcloud functions delete RegistrationFn
