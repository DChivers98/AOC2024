#!/usr/bin/env bash
# run.sh
#
# Usage:  ./run.sh <day>
#   <day>  – 1 … 25

set -euo pipefail
DAY=$(printf "%02d" "$1")

RANGES=(
	"1-12   go"
	"13-25  python"
)

for entry in "${RANGES[@]}"; do
	range=$(awk '{print $1}' <<<"$entry")
	lang=$(awk '{print $2}' <<<"$entry")

	start=${range%-*}
	end=${range#*-}

	if ((10#$DAY >= 10#$start && 10#$DAY <= 10#$end)); then
		case $lang in
		go)
			(cd src/go && go run ./day"$DAY"/main.go)
			;;
		python)
			python src/python/day"$DAY".py
			;;
		esac
		exit 0
	fi
done

echo "❌ Day $DAY not implemented yet" >&2
exit 1
