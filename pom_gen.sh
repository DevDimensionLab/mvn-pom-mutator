#!/usr/bin/env bash
mkdir -p ./pkg/pom_model
xsdgen -o ./pkg/pom_model/pom_model.go -pkg pom_model ./resources/maven-4.0.0.xsd
