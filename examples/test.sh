#!/bin/bash

expected='3'
result=$(go run ../main.go ./sum.pom)
[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

expected='7'
result=$(go run ../main.go ./func.pom)
[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

expected='1'
result=$(go run ../main.go ./array.pom)
[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

expected='Привет мир!'
result=$(go run ../main.go ./print.pom)
[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}" && exit 1)

echo "PASS"