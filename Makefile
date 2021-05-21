deployHelloPubSub:
	gcloud functions deploy HelloPubSub --runtime go113 --trigger-topic arrival_created

deleteHelloPubSub:
	gcloud functions delete HelloPubSub

deployPersonsHook:
	gcloud functions deploy PersonsHook --entry-point PersonsHook --runtime go113 --trigger-event providers/cloud.firestore/eventTypes/document.create  --region us-east1 --trigger-resource "projects/epi-belize-staging/databases/(default)/documents/persons/{pushId}" --source .

deletePersonsHook:
	gcloud functions delete PersonsHook
