#!/bin/sh

tar -czvf bundle.tar.gz -C bundle .
aws s3 cp bundle.tar.gz s3://opa-talk/bundle.tar.gz --acl public-read
rm -rf bundle.tar.gz