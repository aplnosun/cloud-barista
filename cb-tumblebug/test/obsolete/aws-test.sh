#!/bin/bash

export CONN_CONFIG=aws-us-east-1-config
export IMAGE_NAME=ami-085925f297f89fce1
export SPEC_NAME=t3.micro

export NS_ID=253ba8d2-7315-4cc4-adc9-a4ee3b56ba75

./full_test.sh
