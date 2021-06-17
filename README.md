# Port of Entry COVID10 Screening App
The Backend of the COVID19 Port Of Entry Screening App

## Generating Secrets
We use GCP's KMS for managing secrets that some Cloud Functions might need. Generate encrypted secret file:

```
gcloud kms encrypt --ciphertext-file=env.yml.enc --plaintext-file=env.yaml --key=entry-screening-func --keyr
ing=functions-deployment --location=global --verbosity=debug
```

Decrypt the file:

```
gcloud kms decrypt --ciphertext-file=<ENCRYPTED_FILE> --plaintext-file=<OUTPUT_FILE> --location=global --keyring=functions-deployment --key=entry-screening-func --verbosity=debug 
```


## Deploying
The project uses [Google's Cloud Build](https://cloud.google.com/build/docs/quickstarts) for building and deploying.
There are a series of cloud build files per `feature`. This is to avoid having to redeploy the entire project if only
a few functions were changed:

- cloudbuild.screening-updates.yaml deploys all cloud functions related to firestore triggers when the screening collection changes.
- cloudbuild.persons.yaml deploys all functions related to the firestore triggers when the persons collection changes.
- cloudbuild.registration.yaml deploys cloud functions for the registration app
- cloudbuild.firesearch.yaml deploys cloud functions for firesearch utilities

### Deploying from the cli

Persons Functions
```
gcloud builds submit --config=cloudbuild.persons.yaml --async --format=json
```


Screenings Functions
```
gcloud builds submit --config=cloudbuild.screening-updates.yaml --async --format=json
```

Firesearch Functions
```
gcloud builds submit --config=cloudbuild.firesearch.yaml --async --format=json
```

Registration Functions
```
gcloud builds submit --config=cloudbuild.registration.yaml --async --format=json
```
