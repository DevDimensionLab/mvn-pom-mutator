package pom

import "encoding/xml"

type Model struct {
	Comment                string                  `xml:",comment"`
	XMLName                xml.Name                `xml:"project"`
	Xmlns                  string                  `xml:"xmlns,attr"`
	SchemaLocation         string                  `xml:"xsi,attr"`
	Xsi                    string                  `xml:"schemaLocation,attr"`
	ModelVersion           string                  `xml:"modelVersion,omitempty"`
	Parent                 *Parent                 `xml:"parent,omitempty"`
	GroupId                string                  `xml:"groupId,omitempty"`
	ArtifactId             string                  `xml:"artifactId,omitempty"`
	Version                string                  `xml:"version,omitempty"`
	Packaging              string                  `xml:"packaging,omitempty"`
	Name                   string                  `xml:"name,omitempty"`
	Description            string                  `xml:"description,omitempty"`
	Url                    string                  `xml:"url,omitempty"`
	InceptionYear          string                  `xml:"inceptionYear,omitempty"`
	Organization           *Organization           `xml:"organization,omitempty"`
	License                []License               `xml:"license,omitempty"`
	Developer              []Developer             `xml:"developer,omitempty"`
	Contributor            []Contributor           `xml:"contributor,omitempty"`
	MailingList            []MailingList           `xml:"mailingList,omitempty"`
	Prerequisites          *Prerequisites          `xml:"prerequisites,omitempty"`
	Module                 []string                `xml:"module,omitempty"`
	Scm                    *Scm                    `xml:"scm,omitempty"`
	IssueManagement        *IssueManagement        `xml:"issueManagement,omitempty"`
	CiManagement           *CiManagement           `xml:"ciManagement,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	Properties             *Any                    `xml:"properties,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	Dependency             []Dependency            `xml:"dependency,omitempty"`
	Repository             []Repository            `xml:"repository,omitempty"`
	PluginRepository       []Repository            `xml:"pluginRepository,omitempty"`
	Build                  *Build                  `xml:"build,omitempty"`
	Reports                *Any                    `xml:"reports,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
	Profile                []Profile               `xml:"profile,omitempty"`
}

type License struct {
	Comment      string `xml:",comment"`
	Name         string `xml:"name,omitempty"`
	Url          string `xml:"url,omitempty"`
	Distribution string `xml:"distribution,omitempty"`
	Comments     string `xml:"comments,omitempty"`
}

type CiManagement struct {
	Comment  string     `xml:",comment"`
	System   string     `xml:"system,omitempty"`
	Url      string     `xml:"url,omitempty"`
	Notifier []Notifier `xml:"notifier,omitempty"`
}

type Notifier struct {
	Comment       string `xml:",comment"`
	Type_         string `xml:"type,omitempty"`
	SendOnError   bool   `xml:"sendOnError,omitempty"`
	SendOnFailure bool   `xml:"sendOnFailure,omitempty"`
	SendOnSuccess bool   `xml:"sendOnSuccess,omitempty"`
	SendOnWarning bool   `xml:"sendOnWarning,omitempty"`
	Address       string `xml:"address,omitempty"`
	Configuration *Any   `xml:"configuration,omitempty"`
}

type Scm struct {
	Comment             string `xml:",comment"`
	Connection          string `xml:"connection,omitempty"`
	DeveloperConnection string `xml:"developerConnection,omitempty"`
	Tag                 string `xml:"tag,omitempty"`
	Url                 string `xml:"url,omitempty"`
}

type IssueManagement struct {
	Comment string `xml:",comment"`
	System  string `xml:"system,omitempty"`
	Url     string `xml:"url,omitempty"`
}

type DependencyManagement struct {
	Comment    string       `xml:",comment"`
	Dependency []Dependency `xml:"dependency,omitempty"`
}

type Dependency struct {
	Comment    string      `xml:",comment"`
	GroupId    string      `xml:"groupId,omitempty"`
	ArtifactId string      `xml:"artifactId,omitempty"`
	Version    string      `xml:"version,omitempty"`
	Type_      string      `xml:"type,omitempty"`
	Classifier string      `xml:"classifier,omitempty"`
	Scope      string      `xml:"scope,omitempty"`
	SystemPath string      `xml:"systemPath,omitempty"`
	Exclusion  []Exclusion `xml:"exclusion,omitempty"`
	Optional   string      `xml:"optional,omitempty"`
}

type Exclusion struct {
	Comment    string `xml:",comment"`
	ArtifactId string `xml:"artifactId,omitempty"`
	GroupId    string `xml:"groupId,omitempty"`
}

type Parent struct {
	Comment      string `xml:",comment"`
	GroupId      string `xml:"groupId,omitempty"`
	ArtifactId   string `xml:"artifactId,omitempty"`
	Version      string `xml:"version,omitempty"`
	RelativePath string `xml:"relativePath,omitempty"`
}

type Developer struct {
	Comment         string   `xml:",comment"`
	Id              string   `xml:"id,omitempty"`
	Name            string   `xml:"name,omitempty"`
	Email           string   `xml:"email,omitempty"`
	Url             string   `xml:"url,omitempty"`
	Organization    string   `xml:"organization,omitempty"`
	OrganizationUrl string   `xml:"organizationUrl,omitempty"`
	Role            []string `xml:"role,omitempty"`
	Timezone        string   `xml:"timezone,omitempty"`
	Properties      *Any     `xml:"properties,omitempty"`
}

type MailingList struct {
	Comment      string   `xml:",comment"`
	Name         string   `xml:"name,omitempty"`
	Subscribe    string   `xml:"subscribe,omitempty"`
	Unsubscribe  string   `xml:"unsubscribe,omitempty"`
	Post         string   `xml:"post,omitempty"`
	Archive      string   `xml:"archive,omitempty"`
	OtherArchive []string `xml:"otherArchive,omitempty"`
}

type Contributor struct {
	Comment         string   `xml:",comment"`
	Name            string   `xml:"name,omitempty"`
	Email           string   `xml:"email,omitempty"`
	Url             string   `xml:"url,omitempty"`
	Organization    string   `xml:"organization,omitempty"`
	OrganizationUrl string   `xml:"organizationUrl,omitempty"`
	Role            []string `xml:"role,omitempty"`
	Timezone        string   `xml:"timezone,omitempty"`
	Properties      *Any     `xml:"properties,omitempty"`
}

type Organization struct {
	Comment string `xml:",comment"`
	Name    string `xml:"name,omitempty"`
	Url     string `xml:"url,omitempty"`
}

type DistributionManagement struct {
	Comment            string                `xml:",comment"`
	Repository         *DeploymentRepository `xml:"repository,omitempty"`
	SnapshotRepository *DeploymentRepository `xml:"snapshotRepository,omitempty"`
	Site               *Site                 `xml:"site,omitempty"`
	DownloadUrl        string                `xml:"downloadUrl,omitempty"`
	Relocation         *Relocation           `xml:"relocation,omitempty"`
	Status             string                `xml:"status,omitempty"`
}

type DeploymentRepository struct {
	Comment       string            `xml:",comment"`
	UniqueVersion bool              `xml:"uniqueVersion,omitempty"`
	Releases      *RepositoryPolicy `xml:"releases,omitempty"`
	Snapshots     *RepositoryPolicy `xml:"snapshots,omitempty"`
	Id            string            `xml:"id,omitempty"`
	Name          string            `xml:"name,omitempty"`
	Url           string            `xml:"url,omitempty"`
	Layout        string            `xml:"layout,omitempty"`
}

type RepositoryPolicy struct {
	Comment        string `xml:",comment"`
	Enabled        string `xml:"enabled,omitempty"`
	UpdatePolicy   string `xml:"updatePolicy,omitempty"`
	ChecksumPolicy string `xml:"checksumPolicy,omitempty"`
}

type Relocation struct {
	Comment    string `xml:",comment"`
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	Message    string `xml:"message,omitempty"`
}

type Site struct {
	Comment string `xml:",comment"`
	Id      string `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	Url     string `xml:"url,omitempty"`
}

type Reporting struct {
	Comment         string         `xml:",comment"`
	ExcludeDefaults string         `xml:"excludeDefaults,omitempty"`
	OutputDirectory string         `xml:"outputDirectory,omitempty"`
	Plugin          []ReportPlugin `xml:"plugin,omitempty"`
}

type ReportPlugin struct {
	Comment       string      `xml:",comment"`
	GroupId       string      `xml:"groupId,omitempty"`
	ArtifactId    string      `xml:"artifactId,omitempty"`
	Version       string      `xml:"version,omitempty"`
	ReportSet     []ReportSet `xml:"reportSet,omitempty"`
	Inherited     string      `xml:"inherited,omitempty"`
	Configuration *Any        `xml:"configuration,omitempty"`
}

type ReportSet struct {
	Comment       string   `xml:",comment"`
	Id            string   `xml:"id,omitempty"`
	Report        []string `xml:"report,omitempty"`
	Inherited     string   `xml:"inherited,omitempty"`
	Configuration *Any     `xml:"configuration,omitempty"`
}

type Profile struct {
	Comment                string                  `xml:",comment"`
	Id                     string                  `xml:"id,omitempty"`
	Activation             *Activation             `xml:"activation,omitempty"`
	Build                  *BuildBase              `xml:"build,omitempty"`
	Module                 []string                `xml:"module,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	Properties             *Any                    `xml:"properties,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	Dependency             []Dependency            `xml:"dependency,omitempty"`
	Repository             []Repository            `xml:"repository,omitempty"`
	PluginRepository       []Repository            `xml:"pluginRepository,omitempty"`
	Reports                *Any                    `xml:"reports,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
}

type Activation struct {
	Comment         string              `xml:",comment"`
	ActiveByDefault bool                `xml:"activeByDefault,omitempty"`
	Jdk             string              `xml:"jdk,omitempty"`
	Os              *ActivationOS       `xml:"os,omitempty"`
	Property        *ActivationProperty `xml:"property,omitempty"`
	File            *ActivationFile     `xml:"file,omitempty"`
}

type ActivationProperty struct {
	Comment string `xml:",comment"`
	Name    string `xml:"name,omitempty"`
	Value   string `xml:"value,omitempty"`
}

type ActivationFile struct {
	Comment string `xml:",comment"`
	Missing string `xml:"missing,omitempty"`
	Exists  string `xml:"exists,omitempty"`
}

type ActivationOS struct {
	Comment string `xml:",comment"`
	Name    string `xml:"name,omitempty"`
	Family  string `xml:"family,omitempty"`
	Arch    string `xml:"arch,omitempty"`
	Version string `xml:"version,omitempty"`
}

type Repository struct {
	Comment   string            `xml:",comment"`
	Releases  *RepositoryPolicy `xml:"releases,omitempty"`
	Snapshots *RepositoryPolicy `xml:"snapshots,omitempty"`
	Id        string            `xml:"id,omitempty"`
	Name      string            `xml:"name,omitempty"`
	Url       string            `xml:"url,omitempty"`
	Layout    string            `xml:"layout,omitempty"`
}

type BuildBase struct {
	Comment          string            `xml:",comment"`
	DefaultGoal      string            `xml:"defaultGoal,omitempty"`
	Resource         []Resource        `xml:"resource,omitempty"`
	TestResource     []Resource        `xml:"testResource,omitempty"`
	Directory        string            `xml:"directory,omitempty"`
	FinalName        string            `xml:"finalName,omitempty"`
	Filter           []string          `xml:"filter,omitempty"`
	PluginManagement *PluginManagement `xml:"pluginManagement,omitempty"`
	Plugin           []Plugin          `xml:"plugin,omitempty"`
}

type Plugin struct {
	Comment       string            `xml:",comment"`
	GroupId       string            `xml:"groupId,omitempty"`
	ArtifactId    string            `xml:"artifactId,omitempty"`
	Version       string            `xml:"version,omitempty"`
	Extensions    string            `xml:"extensions,omitempty"`
	Execution     []PluginExecution `xml:"execution,omitempty"`
	Dependency    []Dependency      `xml:"dependency,omitempty"`
	Goals         *Any              `xml:"goals,omitempty"`
	Inherited     string            `xml:"inherited,omitempty"`
	Configuration *Any              `xml:"configuration,omitempty"`
}

type PluginExecution struct {
	Comment       string   `xml:",comment"`
	Id            string   `xml:"id,omitempty"`
	Phase         string   `xml:"phase,omitempty"`
	Goal          []string `xml:"goal,omitempty"`
	Inherited     string   `xml:"inherited,omitempty"`
	Configuration *Any     `xml:"configuration,omitempty"`
}

type Resource struct {
	Comment    string   `xml:",comment"`
	TargetPath string   `xml:"targetPath,omitempty"`
	Filtering  string   `xml:"filtering,omitempty"`
	Directory  string   `xml:"directory,omitempty"`
	Include    []string `xml:"include,omitempty"`
	Exclude    []string `xml:"exclude,omitempty"`
}

type PluginManagement struct {
	Comment string   `xml:",comment"`
	Plugin  []Plugin `xml:"plugin,omitempty"`
}

type Prerequisites struct {
	Comment string `xml:",comment"`
	Maven   string `xml:"maven,omitempty"`
}

type Build struct {
	Comment               string            `xml:",comment"`
	SourceDirectory       string            `xml:"sourceDirectory,omitempty"`
	ScriptSourceDirectory string            `xml:"scriptSourceDirectory,omitempty"`
	TestSourceDirectory   string            `xml:"testSourceDirectory,omitempty"`
	OutputDirectory       string            `xml:"outputDirectory,omitempty"`
	TestOutputDirectory   string            `xml:"testOutputDirectory,omitempty"`
	Extension             []Extension       `xml:"extension,omitempty"`
	DefaultGoal           string            `xml:"defaultGoal,omitempty"`
	Resource              []Resource        `xml:"resource,omitempty"`
	TestResource          []Resource        `xml:"testResource,omitempty"`
	Directory             string            `xml:"directory,omitempty"`
	FinalName             string            `xml:"finalName,omitempty"`
	Filter                []string          `xml:"filter,omitempty"`
	PluginManagement      *PluginManagement `xml:"pluginManagement,omitempty"`
	Plugin                []Plugin          `xml:"plugin,omitempty"`
}

type Extension struct {
	Comment    string `xml:",comment"`
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
}

type Any struct {
	XMLName     xml.Name
	Value       string `xml:",chardata"`
	AnyElements []Any  `xml:",any"`
}
