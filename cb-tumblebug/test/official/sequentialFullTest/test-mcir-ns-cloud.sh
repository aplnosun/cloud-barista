#!/bin/bash

function dozing()
{
	duration=$1
	printf "Dozing for %s : " $duration
	for (( i=1; i<=$duration; i++ ))
	do
		printf "%s " $i
		sleep 1
	done
	echo "(Back to work)"
}

source ../conf.env
source ../credentials.conf

echo "####################################################################"
echo "## Create MCIS from Zero Base"
echo "####################################################################"

CSP=${1}
REGION=${2:-1}
POSTFIX=${3:-developer}
if [ "${CSP}" == "aws" ]; then
	echo "[Test for AWS]"
	INDEX=1
elif [ "${CSP}" == "azure" ]; then
	echo "[Test for Azure]"
	INDEX=2
elif [ "${CSP}" == "gcp" ]; then
	echo "[Test for GCP]"
	INDEX=3
elif [ "${CSP}" == "alibaba" ]; then
	echo "[Test for Alibaba]"
	INDEX=4
else
	echo "[No acceptable argument was provided (aws, azure, gcp, alibaba, ...). Default: Test for AWS]"
	CSP="aws"
	INDEX=1
fi

../0.settingSpider/register-cloud.sh $CSP $REGION $POSTFIX
../0.settingTB/create-ns.sh $CSP $REGION $POSTFIX
../1.vNet/create-vNet.sh $CSP $REGION $POSTFIX
dozing 10
if [ "${CSP}" == "gcp" ]; then
	echo "[Test for GCP needs more preparation time]"
	dozing 20
fi
../2.securityGroup/create-securityGroup.sh $CSP $REGION $POSTFIX
dozing 10
../3.sshKey/create-sshKey.sh $CSP $REGION $POSTFIX
../4.image/register-image.sh $CSP $REGION $POSTFIX
../5.spec/register-spec.sh $CSP $REGION $POSTFIX


_self="${0##*/}"

echo ""
echo "[Logging to notify latest command history]"
echo "[CMD] ${_self} ${CSP} ${REGION} ${POSTFIX}" >> ./executionStatus
echo ""
echo "[Executed Command List]"
cat  ./executionStatus
echo ""
