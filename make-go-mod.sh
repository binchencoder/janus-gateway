#!/bin/bash

#bazel_output_base=$(bazel info | grep "^output_base: " | awk -F ": " '{print $2}')
#[[ -z $bazel_output_base ]] && {
#	echo "Error: no bazel output base from 'bazel info'"
#	exit 1
#}

declare -a lines=()
i=0
while IFS= read -r line; do
	lines[i]="$line"
	let "i++"
done <repositories.bzl

requires=""
replaces=""
in_block=false

for ((i = 0; i < ${#lines[@]}; i++)); do
	line=$(echo ${lines[$i]})
	if ! $in_block; then
		if [ "$line" == "go_repository(" ]; then
			in_block=true
			name=""
			importpath=""
			replace=""
			version=""
		fi
		continue
	fi

	if [ "$line" == ")" ]; then
		in_block=false

		if [ -z "$name" ] || [ -z "$importpath" ] || [ -z "$version" ]; then
			continue
		fi

		if [ -z "$replace" ]; then
			requires="$requires\t$importpath $version\n"
		else
			replaces="$replaces\t$importpath $version => $replace $version\n"
		fi
	else
		key=$(echo "$line" | awk -F ' = ' '{print $1}')
		value=$(echo "$line" | awk -F ' = \"' '{print $2}' | sed 's/",//g')

		case "$key" in
		"name")
			name="$value"
			;;
		"importpath")
			importpath="$value"
			;;
		"replace")
			replace="$value"
			;;
		"version")
			version="$value"
			;;
		esac
	fi
done

req=$(echo -e "$requires")
req=$(echo "$req" | sort)

rep=$(echo -e "$replaces")
rep=$(echo "$rep" | sort)

mirrors=$(cat repositories.txt)

cat <<EOF >go.mod
module github.com

go 1.13

require (
$req
)

replace (
$rep
$mirrors
)
EOF
