#!/bin/sh

# Setup:
# Install antlr: pip install antlr4-tools

gen() {
  # antlr4-parse documentation/DesignLanguage.g4 design ./test/data/Test.1.nzsd.txt -tree

  outDir="model/designlanguage/parser"
  mkdir -p "$outDir"
  # Go runtime (once): go get github.com/antlr/antlr4/runtime/Go/antlr/v4@latest
  tempDir=$(mktemp -d)
  antlr4 -Dlanguage=Go -package parser -o "$tempDir" documentation/design/DesignLanguage.g4
  rm -rf "$outDir/" && mv "$tempDir/documentation"/* "$outDir/"
  rm -rf "$tempDir"
}

test() {
  DATA_PATH="$(realpath ./test/data)"

  TEST_DATA_PATH=$DATA_PATH go test -test.v ./test/unit/...
}

case $1 in
  # TODO Is this right?
  b  | build) go install ./cmd/designlanguage/designlanguage.go;;
  f  | format) go fmt ./*;;  
  g  | gen) gen;;
  h  | help)  echo "build|gen (antlr)|help|test|test_functional";;
  t  | test)  test;;
  tf | test_functional)  go test ./test/functional/...;;
  *) echo "Unknown argument";;
esac
