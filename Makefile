deployHelloPubSub:
	gcloud functions deploy HelloPubSub --runtime go113 --trigger-topic arrival_created

deleteHelloPubSub:
	gcloud functions delete HelloPubSub

deployPersonsHook:
	gcloud functions deploy PersonsHook --entry-point PersonsHook --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.create  --region us-east1 --trigger-resource "projects/epi-belize-staging/databases/(default)/documents/persons/{pushId}" --source . --env-vars-file env.yaml

deletePersonsHook:
	gcloud functions delete PersonsHook

deployPersonDeletedListener:
	gcloud functions deploy PersonDeletedListener --entry-point PersonDeletedListener --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.delete  --region us-east1 --trigger-resource "projects/epi-belize-staging/databases/(default)/documents/persons/{pushId}" --source . --env-vars-file env.yaml

deletePersonDeletedListener:
	gcloud functions delete PersonDeletedLIstener
