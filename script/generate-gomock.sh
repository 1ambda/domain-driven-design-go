#!/usr/bin/env bash

TAG="[$(basename -- "$0")]"

MOCK_PREFIX="mock"

USAGE="Usage: PKG_DIR={RELATIVE_PATH} $0"

if [ "${PKG_DIR}" == "" ]; then
    echo ${USAGE} >&2
    exit 1
fi

WORKSPACE=$(echo ${PKG_DIR} | tr -d " ")

GO_FILES=$(ls ${PKG_DIR}/*.go | xargs -n 1 basename)
for GO_FILE in ${GO_FILES[@]}
do

    if [[ ${GO_FILE} == *"test.go" ]]; then
        continue
    fi

    PKG=$(basename ${PKG_DIR})
    MOCK_FILE=${PKG_DIR}/${MOCK_PREFIX}_${GO_FILE}
    mockgen -package ${PKG} -source ${PKG_DIR}/${GO_FILE} -destination ${MOCK_FILE}

    if [ -f ${MOCK_FILE} ];
    then

      # remove if the generated file doesn't include `struct` (useless mock files)
      if cat ${MOCK_FILE} | grep -q --line-buffered struct;
      then
          echo -e "\t$TAG Generated: ${MOCK_FILE}"
      else
          echo -e "\t$TAG Removing: ${MOCK_FILE}"
          rm ${MOCK_FILE}
      fi
    fi

done