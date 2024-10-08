#!/bin/bash

export KEYCLOAK_URL="http://keycloak.internal"

grep KEYCLOAK_TF_CLIENT_SECRET ~/.keycloak.secrets  || echo "KEYCLOAK_TF_CLIENT_SECRET=$(openssl rand -base64 32)" >> ~/.keycloak.secrets
. ~/.keycloak.secrets

grep OIDC_CLIENT_SECRET ~/.iconrepo.secrets || echo "OIDC_CLIENT_SECRET=$(openssl rand -base64 32)" >> ~/.iconrepo.secrets
. ~/.iconrepo.secrets
ICONREPO_CLIENT_SECRET="$OIDC_CLIENT_SECRET"

bash ./wait-for-local-keycloak.sh
bash ./create-terraform-client.sh

# export TF_LOG=DEBUG
cd realm-client &&
  terraform init &&
  terraform apply -auto-approve \
    -var=keycloak_url=${KEYCLOAK_URL} \
    -var="tf_client_secret=$KEYCLOAK_TF_CLIENT_SECRET" \
    -var="client_secret=$ICONREPO_CLIENT_SECRET"

cd ../groups-users &&
  terraform init &&
  terraform apply -auto-approve \
    -var=keycloak_url=${KEYCLOAK_URL} \
    -var="tf_client_secret=$KEYCLOAK_TF_CLIENT_SECRET"
