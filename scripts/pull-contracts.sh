#!/bin/sh

CONTRACT_REPO_NAME="$1"
CONTRACT_REPO_BRANCH="$2"

# update or clone contracts repo
if [ -d "$CONTRACT_REPO_NAME" ]; then
	cd "$CONTRACT_REPO_NAME"
	git fetch
	git checkout "$CONTRACT_REPO_BRANCH"
	git pull
else
	git clone --branch "$CONTRACT_REPO_BRANCH" https://github.com/forta-network/"$CONTRACT_REPO_NAME"
	cd "$CONTRACT_REPO_NAME"
fi

# generate compatible ABIs
yarn install
npx hardhat run scripts/abis-without-custom-errors.js
