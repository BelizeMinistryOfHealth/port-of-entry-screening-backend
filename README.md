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
