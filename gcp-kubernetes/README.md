# Application Deployment on GKE

- [x] GLBC Ingress
- [x] Let's Encrypt Cert using cert-manager 
- [x] Spinnaker

## Secrets

- [x] Database
- [x[ Session Store

```bash
cat <<EOF | kubectl create -f -
apiVersion: v1
kind: Secret
metadata:
  name: database
type: Opaque
data:
  DATABASE_USERNAME: $(echo -n "" | base64)
  DATABASE_PASSWORD: $(echo -n "" | base64)
  DATABASE_HOST: $(echo -n "" | base64)
  DATABASE_PORT: $(echo -n "" | base64)
  DATABASE_NAME: $(echo -n "" | base64)
EOF

# validate 
export FIELD_KEY=DATABASE_USERNAME; kubectl get secret database --template="{{.data.${FIELD_KEY}}}" | base64 --decode; echo "";
export FIELD_KEY=DATABASE_PASSWORD; kubectl get secret database --template="{{.data.${FIELD_KEY}}}" | base64 --decode; echo "";
export FIELD_KEY=DATABASE_HOST; kubectl get secret database --template="{{.data.${FIELD_KEY}}}" | base64 --decode; echo "";
export FIELD_KEY=DATABASE_PORT; kubectl get secret database --template="{{.data.${FIELD_KEY}}}" | base64 --decode; echo "";
export FIELD_KEY=DATABASE_NAME; kubectl get secret database --template="{{.data.${FIELD_KEY}}}" | base64 --decode; echo "";
```

```bash
cat <<EOF | kubectl create -f -
apiVersion: v1
kind: Secret
metadata:
  name: session-store 
type: Opaque
data:
  SESSION_SECRET: $(echo -n "" | base64)
EOF

# validate 
export FIELD_KEY=SESSION_SECRET; kubectl get secret session --template="{{.data.${FIELD_KEY}}}" | base64 --decode; echo "";
```

## Deployment (using kubectl)

## Deployment (using spinnaker)

TBD

## Secrets



``````