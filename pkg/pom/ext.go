package pom

import (
	"errors"
	"github.com/perottobc/mvn-pom-mutator/pkg/mvn_crud"
	"io/ioutil"
)

func (deps Dependencies) FindArtifact(artifactId string) (Dependency, error) {

	for _, dep := range deps.Dependency {
		if dep.ArtifactId == artifactId {
			return dep, nil
		}
	}

	return Dependency{}, errors.New("could not find artifact " + artifactId + " in dependencies")
}

func (any Any) FindKey(key string) (string, error) {
	for _, a := range any.AnyElements {
		if a.XMLName.Local == key {
			return a.Value, nil
		}
	}

	return "", errors.New("could not find key " + key + " in any structure")
}

func (any *Any) SetKey(key string, value string) error {
	for _, a := range any.AnyElements {
		if a.XMLName.Local == key {
			a.Value = value
			return nil
		}
	}

	return errors.New("could not find key " + key + " in any structure")
}

func (model *Model) WriteToFile(outputFile string) error {
	bytes, err := mvn_crud.Marshall(model)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFile, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
