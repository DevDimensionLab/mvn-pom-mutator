# DOCs
FISH: add the go bin path to be able to execute our programs
set -x PATH $PATH /usr/local/go/bin $GOPATH/bin

#https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
go install $GOPATH/src/github.com/spf13/cobra/cobra 

Oppstart av prosjekt: cobra init --pkg-name "mvn-cli"

* resources/maven-4.0.0.xsd => https://www.onlinetool.io/xmltogo => pkg/xsd_model/schema.go
* pkg/pom_gen/pom_gen_test.go => target/pom.go (structs for reading pom-model)
* target/pom.go => pkg/pom/pom.go
* mvn_crud/mvn_crud_test.go demos read/write pom  
