#!/usr/bin/env bash

if [[ -z "${SWAGGER_FILE}" ]]; then
    echo -e "$0: `SWAGGER_FILE` is empty"
    exit 1
fi

if [[ -z "${VERSION}" ]]; then
    echo -e "$0: `VERSION` is empty"
    exit 1
fi


sed -i.bak s/^\s*version:.*/version:\ ${VERSION}/g ${SWAGGER_FILE}
rm ${SWAGGER_FILE}.bak || true