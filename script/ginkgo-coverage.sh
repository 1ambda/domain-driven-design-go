#!/usr/bin/env bash

set -eo pipefail

covermode=${COVERMODE:-atomic}
coverdir=$(mktemp -d /tmp/coverage.XXXXXXXXXX)
profile="${coverdir}/cover.out"
coveragetxt="coverage.merged"

hash goveralls 2>/dev/null || go get github.com/mattn/goveralls
hash godir 2>/dev/null || go get github.com/Masterminds/godir

remove_mock_from_cover_data() {
  for file in $(find . -type f -name "*.coverprofile")
  do
    grep -v "mock_" ${file} > ${file}.tmp;
    rm -rf ${file}
    mv ${file}.tmp ${file}
  done
}

generate_cover_data() {
  # assume that coverage report is already generated
  echo "" > ${coveragetxt}
  find . -type f -name "*.coverprofile" | while read -r file; do cat $file >> ${coveragetxt} && mv $file ${coverdir}; done
  echo "mode: $covermode" >"$profile"
  grep -h -v "^mode:" "$coverdir"/*.coverprofile >>"$profile"
}

push_to_coveralls() {
  goveralls -coverprofile="${profile}" -service=circle-ci -repotoken $COVERALLS_REPO_TOKEN || echo "push to coveralls failed"
}

push_to_codecov() {
  bash <(curl -s https://codecov.io/bash) || echo "push to codecov failed"
}

remove_mock_from_cover_data
generate_cover_data
go tool cover -func "${profile}"

case "${1-}" in
  --html)
    go tool cover -html "${profile}"
    ;;
  --coveralls)
		if [ -z $COVERALLS_REPO_TOKEN ]; then
			echo '$COVERALLS_REPO_TOKEN not set. Skipping pushing coverage report to coveralls.io'
			exit
		fi
    push_to_coveralls
    ;;
  --codecov)
    push_to_codecov
    ;;
esac