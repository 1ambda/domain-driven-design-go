# Terraform: GCP

Configure gcloud and `GOOGLE_PROJECT`.

```bash
gcloud auth application-default login
test -z DEVSHELL_GCLOUD_CONFIG && gcloud auth application-default login

gcloud config set project {YOUR_PROJECT}
export GOOGLE_PROJECT=$(gcloud config get-value project)
```

Then apply terraform to bring GKE cluster.

```bash
terraform init
terraform apply

# get kubectl context
gcloud container clusters get-credentials g-street # or your cluster name
```