#!/bin/bash

PROJECT=squat-cloud
CLUSTER=squat-cluster
ZONE=us-central1-a
DOCKER_IMAGE=squat-backend

set -e

docker build -t gcr.io/${PROJECT}/${DOCKER_IMAGE}:$TRAVIS_COMMIT .

echo $GCLOUD_SERVICE_KEY | base64 --decode -i > ${HOME}/gcloud-service-key.json
gcloud auth activate-service-account --key-file ${HOME}/gcloud-service-key.json

gcloud --quiet config set project ${PROJECT}
gcloud --quiet config set container/cluster ${CLUSTER}
gcloud --quiet config set compute/zone ${ZONE}
gcloud --quiet container clusters get-credentials ${CLUSTER}

gcloud docker -- push gcr.io/${PROJECT}/${DOCKER_IMAGE}

yes | gcloud beta container images add-tag gcr.io/${PROJECT}/${DOCKER_IMAGE}:$TRAVIS_COMMIT gcr.io/${PROJECT}/${DOCKER_IMAGE}:latest
