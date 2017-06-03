#!/usr/bin/env bash
VERBOSE=0
while getopts ":v" opt; do
  case $opt in
    v)
      VERBOSE=true
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done

cd development
echo "Checking dist files"
suffix=.dist

for parameters in $(find . -name \*.dist)
do
    yaml=${parameters%$suffix}
    echo "- $yaml"
    if [ ! -f $yaml ]; then
        echo "generated"
        cp $parameters $yaml
    fi
done

docker-compose up --build -d --remove-orphans
docker-compose run --service-ports web
