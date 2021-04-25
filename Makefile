deployHelloPubSub:
	gcloud functions deploy HelloPubSub --runtime go113 --trigger-topic arrival_created

deleteHelloPubSub:
	gcloud functions delete HelloPubSub
