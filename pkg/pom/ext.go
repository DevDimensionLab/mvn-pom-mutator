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
		var found = false
		if nil != model.Dependencies {
			found = SetDependencyVersionElement(model.Dependencies, dep, newVersion)
		}
		if nil != model.DependencyManagement && nil != model.DependencyManagement.Dependencies {
			found = found || SetDependencyVersionElement(model.DependencyManagement.Dependencies, dep, newVersion)
		}
		if found {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("error setting new version [%s] for %s:%s", newVersion, dep.GroupId, dep.ArtifactId))
}

func SetDependencyVersionElement(dependencies *Dependencies, dep Dependency, newVersion string) bool {
	for i, d := range dependencies.Dependency {
		if cmp.Equal(dep, d) {
			dependencies.Dependency[i].Version = newVersion
			dep.Version = newVersion
			return true
		}
	}
	return false
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

func (model *Model) ReplaceVersionTagForDependency(dep Dependency) error {
	if strings.HasPrefix(dep.Version, "${") {
		return errors.New("version tag already contains a variable")
	}

	if model.Properties == nil {
		model.Properties = &Any{}
	}

	if model.Dependencies != nil {
		for i, d := range model.Dependencies.Dependency {
			if cmp.Equal(dep, d) {
				versionKey := dep.ArtifactId
				versionTag := fmt.Sprintf("%s.version", versionKey)
				versionVariable := fmt.Sprintf("${%s}", versionTag)
				model.Dependencies.Dependency[i].Version = versionVariable
				return model.Properties.AddKey(versionTag, dep.Version)
			}
		}
	}

	return errors.New(fmt.Sprintf("could not find dependency: %s:%s in model.Dependencies", dep.GroupId, dep.ArtifactId))
}

func (model *Model) ReplaceVersionTagForDependencyManagement(dep Dependency) error {
	if strings.HasPrefix(dep.Version, "${") {
		return errors.New("version tag already contains a variable")
	}

	if model.Properties == nil {
		model.Properties = &Any{}
	}

	if model.DependencyManagement != nil && model.DependencyManagement.Dependencies != nil {
		for i, d := range model.DependencyManagement.Dependencies.Dependency {
			if cmp.Equal(dep, d) {
				versionKey := dep.ArtifactId
				versionTag := fmt.Sprintf("%s.version", versionKey)
				versionVariable := fmt.Sprintf("${%s}", versionTag)
				model.DependencyManagement.Dependencies.Dependency[i].Version = versionVariable
				return model.Properties.AddKey(versionTag, dep.Version)
			}
		}
	}

	return errors.New(fmt.Sprintf("could not find dependency: %s:%s in model.dependencyManagement", dep.GroupId, dep.ArtifactId))
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

func (any *Any) AddKey(key string, value string) error {
	for _, a := range any.AnyElements {
		if a.XMLName.Local == key {
			if a.Value != value {
				return errors.New(fmt.Sprintf("found another key: %s with value %s", key, value))
			} else {
				return nil
			}
		}
	}

	any.AnyElements = append(any.AnyElements, Any{
		XMLName: struct{ Space, Local string }{Space: "", Local: key},
		Value:   value,
	})

	return nil
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

func (model *Model) RemoveDependency(dep Dependency) error {
	for i, d := range model.Dependencies.Dependency {
		if cmp.Equal(dep, d) {
			model.Dependencies.Dependency = append(model.Dependencies.Dependency[:i], model.Dependencies.Dependency[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("could not find dependency: %s:%s in model", dep.GroupId, dep.ArtifactId))
}

func (model *Model) FindDependency(groupId string, artifactId string) (Dependency, error) {
	for _, dep := range model.Dependencies.Dependency {
		if dep.GroupId == groupId && dep.ArtifactId == artifactId {
			return dep, nil
		}
	}

	errMsg := fmt.Sprintf("could not find dependency: %s:%s in model", groupId, artifactId)
	return Dependency{}, errors.New(errMsg)
}

func (model *Model) InsertDependency(dep Dependency) {
	_, err := model.FindDependency(dep.GroupId, dep.GroupId)
	if err == nil {
		return
	}

	model.Dependencies.Dependency = append(model.Dependencies.Dependency, dep)
}

func (model *Model) GetGroupId() (groupId string) {
	groupId = model.GroupId
	if groupId == "" {
		groupId = model.Parent.GroupId
	}

	return
}
func (model *Model) GetSecondPartyGroupId() (string, error) {
	if model.GetGroupId() != "" {
		return GetFirstTwoPartsOfGroupId(model.GetGroupId())
	}

	return "", errors.New("could not extract 2party groupId")
}

func GetFirstTwoPartsOfGroupId(groupId string) (string, error) {
	parts := strings.Split(groupId, ".")

	if len(parts) <= 1 {
		return "", errors.New("groupId must at least contain two punctuations")
	} else {
		return strings.Join(parts[:2], "."), nil
	}
}
