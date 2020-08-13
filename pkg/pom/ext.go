package pom

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
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

func (model *Model) GetDependencyVersion(dep Dependency) (string, error) {
	if strings.HasPrefix(dep.Version, "${") {
		versionKey := strings.Trim(dep.Version, "${}")
		return model.Properties.FindKey(versionKey)
	} else {
		return dep.Version, nil
	}
}

func (model *Model) GetPluginVersion(plugin Plugin) (string, error) {
	if strings.HasPrefix(plugin.Version, "${") {
		versionKey := strings.Trim(plugin.Version, "${}")
		return model.Properties.FindKey(versionKey)
	} else {
		return plugin.Version, nil
	}
}

func (model *Model) SetDependencyVersion(dep Dependency, newVersion string) error {
	if strings.HasPrefix(dep.Version, "${") {
		versionKey := strings.Trim(dep.Version, "${}")
		return model.Properties.SetKey(versionKey, newVersion)
	} else {
		for i, d := range model.Dependencies.Dependency {
			if cmp.Equal(dep, d) {
				model.Dependencies.Dependency[i].Version = newVersion
				dep.Version = newVersion
				return nil
			}
		}
	}

	return errors.New(fmt.Sprintf("error setting new version [%s] for %s:%s", newVersion, dep.GroupId, dep.ArtifactId))
}

func (model *Model) SetPluginVersion(plugin Plugin, newVersion string) error {
	if strings.HasPrefix(plugin.Version, "${") {
		versionKey := strings.Trim(plugin.Version, "${}")
		return model.Properties.SetKey(versionKey, newVersion)
	} else {
		for i, d := range model.Build.Plugins.Plugin {
			if cmp.Equal(plugin, d) {
				model.Build.Plugins.Plugin[i].Version = newVersion
				plugin.Version = newVersion
				return nil
			}
		}
	}

	return errors.New(fmt.Sprintf("error setting new version [%s] for %s:%s", newVersion, plugin.GroupId, plugin.ArtifactId))
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
