# Build and push the annotate Orange function

if [ $# -eq 0 ]
  then
    echo "No arguments supplied. defaulting to version 1.0.0"
    docker build orange-otel-annotate/Dockerfile -t marcinziolkowski/orange-otel-annotate:1.0.0
    docker push marcinziolkowski/orange-otel-annotate:1.0.0
  else
    echo "Building and tagging for TAG: $1"
fi

docker build orange-otel-annotate/Dockerfile -t marcinziolkowski/orange-otel-annotate:$1
docker push marcinziolkowski/orange-otel-annotate:$1