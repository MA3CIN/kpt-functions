# kpt-functions

My own kpt-functions, developed for the Nephio R3 Observability POC.

## Getting started

Head over to the ``` /smf-package ``` to view the SMF Free5GC package. Then, run ``` kpt eval ``` to run the kpt package declarative pipeline, and observe the OpenTelemetry Operator annotations appearing in the Network Function yaml definition.

## Running the functions
Use  ```kpt fn eval ./testdata/test1/resources.yaml --image marcinziolkowski/orange-otel-annotate:1.0.0 -- by-path='spec.**.containerPort' ``` to run the kpt functions. The values to pass should be as follow:  "First annotation" "Second annotation"


## Rebuilding docker images for functions

Use ``` ./build-and-push-images.sh ``` to re-build the images and push them to a registry of your choice.