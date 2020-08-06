package pom

import (
	"errors"
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
