package pom

import (
	"errors"
	"io/ioutil"
	"strings"
)

func (deps Dependencies) FindArtifact(artifactId string) (Dependency, error) {

	for _, dep := range deps.Dependency {
		if dep.ArtifactId == artifactId {
			return dep, nil
		}
	}

	return Dependency{}, errors.New("could not find artifact " + artifactId + " in dependencies")
}

func (model *Model) GetVersion(dep Dependency) (string, error) {
	if strings.HasPrefix(dep.Version, "${") {
		versionKey := strings.Trim(dep.Version, "${}")
		return model.Properties.FindKey(versionKey)
	} else {
		return dep.Version, nil
	}
}

func (model *Model) SetVersion(dep *Dependency, newVersion string) error {
	if strings.HasPrefix(dep.Version, "${") {
		versionKey := strings.Trim(dep.Version, "${}")
		return model.Properties.SetKey(versionKey, newVersion)
	} else {
		dep.Version = newVersion
		return nil
	}
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
	for i, a := range any.AnyElements {
		if a.XMLName.Local == key {
			any.AnyElements[i].Value = value
			return nil
		}
	}

	return errors.New("could not find key " + key + " in any structure")
}

func (model *Model) WriteToFile(outputFile string) error {
	bytes, err := Marshall(model)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFile, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
