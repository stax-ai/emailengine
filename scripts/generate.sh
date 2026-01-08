#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
# Require swagger.json (OpenAPI 3 JSON)
SPEC="${ROOT_DIR}/swagger.json"
if [[ ! -f "${SPEC}" ]]; then
  echo "swagger.json not found at ${SPEC}. Please provide a valid OpenAPI v3 JSON spec."
  exit 1
fi
API_DIR="${ROOT_DIR}/api"
BIN_DIR="${ROOT_DIR}/.bin"
OAPI_BIN="${BIN_DIR}/oapi-codegen"
FIX_JSON_BIN="${BIN_DIR}/fix-swagger"
FIXED_JSON_SPEC="${ROOT_DIR}/swagger.fixed.json"

mkdir -p "${API_DIR}" "${BIN_DIR}"

# Pin a known-good oapi-codegen version
VERSION="v2.4.1"

if [[ ! -x "${OAPI_BIN}" ]]; then
  echo "Installing oapi-codegen ${VERSION} into ${BIN_DIR}..."
  # Force public Go proxy to avoid corporate proxies which may be unavailable in CI/sandbox
  GOPROXY_ENV="https://proxy.golang.org,direct"
  GOSUMDB_ENV="sum.golang.org"
  # Ensure no corporate proxy/private overrides are used
  GONOPROXY_ENV="none"
  GOPRIVATE_ENV=""
  GOBIN="${BIN_DIR}" GOPROXY="${GOPROXY_ENV}" GOSUMDB="${GOSUMDB_ENV}" GONOPROXY="${GONOPROXY_ENV}" GOPRIVATE="${GOPRIVATE_ENV}" \
    go install "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@${VERSION}"
fi

# Sanitize swagger.json -> swagger.fixed.json (strip problematic fields)
if [[ ! -x "${FIX_JSON_BIN}" ]]; then
  echo "Building Swagger JSON sanitizer..."
  GOBIN="${BIN_DIR}" GOPROXY="https://proxy.golang.org,direct" GOSUMDB="sum.golang.org" \
    go build -o "${FIX_JSON_BIN}" "${ROOT_DIR}/scripts/fix_swagger_json.go"
fi
echo "Sanitizing ${SPEC} -> ${FIXED_JSON_SPEC} (strip examples and enums with backslashes)..."
"${FIX_JSON_BIN}" "${SPEC}" "${FIXED_JSON_SPEC}"
SPEC="${FIXED_JSON_SPEC}"

echo "Generating Go client with openapi-generator-cli (Docker)..."
rm -rf "${API_DIR}"/*
docker run --rm -v "${ROOT_DIR}:/local" openapitools/openapi-generator-cli:v7.9.0 generate \
  -i "/local/${SPEC##${ROOT_DIR}/}" \
  -g go \
  -o /local/api \
  --skip-validate-spec \
  --package-name api \
  --additional-properties=packageName=api,enumClassPrefix=true,withGoMod=false,allowUnicodeIdentifiers=true,preloadEnvVars=true
echo "OpenAPI Generator completed."


