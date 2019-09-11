#!/bin/sh

GITHUB_USER=santiagopoli
GITHUB_REPO=k8s-meetup-20191112

tar -czvf bundle.tar.gz -C bundle .
github-release upload --replace --user $GITHUB_USER --repo $GITHUB_REPO --tag latest --name "bundle.tar.gz" --file bundle.tar.gz
rm -rf bundle.tar.gz