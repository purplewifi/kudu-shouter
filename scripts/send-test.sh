#!/bin/bash

# dodgy hack to ensure we are in the expected dir for the payload file
if [[ "$(basename "${PWD}")" != "scripts" ]]; then
  cd scripts || exit 2
fi

curl -v --data-binary @example-payload.json http://127.0.0.1:7890/capture