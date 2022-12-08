#!/usr/bin/bash

expected='3'
result=$(go run ../main.go ./sum.pom)
[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

expected='7'
result=$(go run ../main.go ./func.pom)
[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

echo "PASS"