#!/usr/bin/env bash

PACKAGE_NAME=resources
GENERATOR_IMAGE=openapitools/openapi-generator-cli:v7.2.0
OPENAPI_DIR="$(pwd)/docs/web_deploy"
GENERATED="$(pwd)/resources"

function printHelp {
    echo "usage: ./generate.sh [<flags>]
            script to generate Go model bindings based on docs

            Flags:
                  --package PACKAGE        generated package name (default - $PACKAGE_NAME)
                  --image IMAGE            generator docker image (default - $GENERATOR_IMAGE)

              -h, --help                   show this message
              -p, --path-to-generate PATH  path to put generated code (default - $PACKAGE_NAME)
              -i, --input OPENAPI_DIR      path to dir where openapi.yaml is stored (default - $OPENAPI_DIR)"
}

function parseArgs {
    while [[ -n "$1" ]]
    do
        case "$1" in
            -h | --help)
                printHelp && exit 0
                ;;
            -p | --path-to-generate) shift
                [[ ! -d $1 ]] && echo "path $1 does not exist or not a dir" && exit 1
                GENERATED=$1
                ;;
            --package) shift
                [[ -z "$1" ]] && echo "package name not specified" && exit 1
                PACKAGE_NAME=$1
                ;;
            -i | --input) shift
                [[ ! -f "$1/openapi.yaml" ]] && echo "file openapi.yaml does not exist in $1 or not a file" && exit 1
                OPENAPI_DIR=$1
                ;;
            --image) shift
                [[ "$(docker images -q "$1")" == "" ]] && echo "image $1 does not exist locally" && exit 1
                GENERATOR_IMAGE=$1
                ;;
        esac
        shift
    done
}

function generate {
    (cd docs && npm run build)
    rm -rf "${GENERATED}"
    docker run --rm \
          -v $(pwd)/.openapi-generator-ignore:/generator/.openapi-generator-ignore \
          -v "${OPENAPI_DIR}":/openapi \
          -v "${GENERATED}":/generated \
          "${GENERATOR_IMAGE}" sh -c "
            /usr/local/bin/docker-entrypoint.sh \
              generate \
                  --input-spec /openapi/openapi.yaml \
                  --skip-validate-spec \
                  --generator-name go \
                  --output /generated \
                  --ignore-file-override /generator/.openapi-generator-ignore \
                  --package-name ${PACKAGE_NAME}; \
            chown -R $(id -u):$(id -g) /generated"
    goimports -w "${GENERATED}"
}

parseArgs "$@"
generate
