#!/bin/bash

source ../conf.env

echo "####################################################################"
echo "## 5. spec: List"
echo "####################################################################"


curl -sX GET http://$TumblebugServer/tumblebug/ns/$NS_ID/resources/spec | json_pp #|| return 1
