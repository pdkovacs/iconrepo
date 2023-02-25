#!/bin/bash

echo "Executing curl...."

bash ./wait-for-local-keycloak.sh
bash ./create-terraform-client.sh

# export TF_LOG=DEBUG
cd && \
  terraform init && \
	terraform apply -auto-approve