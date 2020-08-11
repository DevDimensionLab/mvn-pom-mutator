package pom

import "encoding/xml"

type Licenses struct {
	License []License `xml:"license,omitempty"`
}

type Developers struct {
	Developer []Developer `xml:"developer,omitempty"`
}

type Contributors struct {
	Contributor []Contributor `xml:"contributor,omitempty"`
}

type MailingLists struct {
	MailingList []MailingList `xml:"mailingList,omitempty"`
}

type Modules struct {
	Module []string `xml:"module,omitempty"`
}

type Dependencies struct {
	Comment    string       `xml:",comment"`
	Dependency []Dependency `xml:"dependency,omitempty"`
}

type Repositories struct {
	Repository []Repository `xml:"repository,omitempty"`
}

type PluginRepositories struct {
	PluginRepository []Repository `xml:"pluginRepository,omitempty"`
}

type Profiles struct {
	Profile []Profile `xml:"profile,omitempty"`
}

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
	Licenses               *Licenses               `xml:"licenses,omitempty"`
	Developers             *Developers             `xml:"developers,omitempty"`
	Contributors           *Contributors           `xml:"contributors,omitempty"`
	MailingLists           *MailingLists           `xml:"mailingLists,omitempty"`
	Prerequisites          *Prerequisites          `xml:"prerequisites,omitempty"`
	Modules                *Modules                `xml:"modules,omitempty"`
	Scm                    *Scm                    `xml:"scm,omitempty"`
	IssueManagement        *IssueManagement        `xml:"issueManagement,omitempty"`
	CiManagement           *CiManagement           `xml:"ciManagement,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	Properties             *Any                    `xml:"properties,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	Dependencies           *Dependencies           `xml:"dependencies,omitempty"`
	Repositories           *Repositories           `xml:"repositories,omitempty"`
	PluginRepositories     *PluginRepositories     `xml:"pluginRepositories,omitempty"`
	Build                  *Build                  `xml:"build,omitempty"`
	Reports                *[]Any                  `xml:"reports,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
	Profiles               *Profiles               `xml:"profiles,omitempty"`
}

type License struct {
	Name         string `xml:"name,omitempty"`
	Url          string `xml:"url,omitempty"`
	Distribution string `xml:"distribution,omitempty"`
	Comments     string `xml:"comments,omitempty"`
}

type Notifiers struct {
	Notifier []Notifier `xml:"notifier,omitempty"`
}

type CiManagement struct {
	System    string     `xml:"system,omitempty"`
	Url       string     `xml:"url,omitempty"`
	Notifiers *Notifiers `xml:"notifiers,omitempty"`
}

type Notifier struct {
	Type_         string `xml:"type,omitempty"`
	SendOnError   bool   `xml:"sendOnError,omitempty"`
	SendOnFailure bool   `xml:"sendOnFailure,omitempty"`
	SendOnSuccess bool   `xml:"sendOnSuccess,omitempty"`
	SendOnWarning bool   `xml:"sendOnWarning,omitempty"`
	Address       string `xml:"address,omitempty"`
	Configuration *[]Any `xml:"configuration,omitempty"`
}

type Scm struct {
	Comment             string `xml:",comment"`
	Connection          string `xml:"connection,omitempty"`
	DeveloperConnection string `xml:"developerConnection,omitempty"`
	Tag                 string `xml:"tag,omitempty"`
	Url                 string `xml:"url,omitempty"`
}

type IssueManagement struct {
	System string `xml:"system,omitempty"`
	Url    string `xml:"url,omitempty"`
}

type DependencyManagement struct {
	Dependencies *Dependencies `xml:"dependencies,omitempty"`
}

type Exclusions struct {
	Comment   string      `xml:",comment"`
	Exclusion []Exclusion `xml:"exclusion,omitempty"`
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
	Exclusions *Exclusions `xml:"exclusions,omitempty"`
	Optional   string      `xml:"optional,omitempty"`
}

type Exclusion struct {
	ArtifactId string `xml:"artifactId,omitempty"`
	GroupId    string `xml:"groupId,omitempty"`
}

type Parent struct {
	GroupId      string `xml:"groupId,omitempty"`
	ArtifactId   string `xml:"artifactId,omitempty"`
	Version      string `xml:"version,omitempty"`
	RelativePath string `xml:"relativePath,omitempty"`
}

type Roles struct {
	Role []string `xml:"role,omitempty"`
}

type Developer struct {
	Id              string `xml:"id,omitempty"`
	Name            string `xml:"name,omitempty"`
	Email           string `xml:"email,omitempty"`
	Url             string `xml:"url,omitempty"`
	Organization    string `xml:"organization,omitempty"`
	OrganizationUrl string `xml:"organizationUrl,omitempty"`
	Roles           *Roles `xml:"roles,omitempty"`
	Timezone        string `xml:"timezone,omitempty"`
	Properties      *[]Any `xml:"properties,omitempty"`
}

type OtherArchives struct {
	OtherArchive []string `xml:"otherArchive,omitempty"`
}

type MailingList struct {
	Name          string         `xml:"name,omitempty"`
	Subscribe     string         `xml:"subscribe,omitempty"`
	Unsubscribe   string         `xml:"unsubscribe,omitempty"`
	Post          string         `xml:"post,omitempty"`
	Archive       string         `xml:"archive,omitempty"`
	OtherArchives *OtherArchives `xml:"otherArchives,omitempty"`
}

type Contributor struct {
	Name            string `xml:"name,omitempty"`
	Email           string `xml:"email,omitempty"`
	Url             string `xml:"url,omitempty"`
	Organization    string `xml:"organization,omitempty"`
	OrganizationUrl string `xml:"organizationUrl,omitempty"`
	Roles           *Roles `xml:"roles,omitempty"`
	Timezone        string `xml:"timezone,omitempty"`
	Properties      *[]Any `xml:"properties,omitempty"`
}

type Organization struct {
	Name string `xml:"name,omitempty"`
	Url  string `xml:"url,omitempty"`
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
	UniqueVersion bool              `xml:"uniqueVersion,omitempty"`
	Releases      *RepositoryPolicy `xml:"releases,omitempty"`
	Snapshots     *RepositoryPolicy `xml:"snapshots,omitempty"`
	Id            string            `xml:"id,omitempty"`
	Name          string            `xml:"name,omitempty"`
	Url           string            `xml:"url,omitempty"`
	Layout        string            `xml:"layout,omitempty"`
}

type RepositoryPolicy struct {
	Enabled        string `xml:"enabled,omitempty"`
	UpdatePolicy   string `xml:"updatePolicy,omitempty"`
	ChecksumPolicy string `xml:"checksumPolicy,omitempty"`
}

type Relocation struct {
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
	Message    string `xml:"message,omitempty"`
}

type Site struct {
	Id   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	Url  string `xml:"url,omitempty"`
}

type Plugins struct {
	Comment string   `xml:",comment"`
	Plugin  []Plugin `xml:"plugin,omitempty"`
}

type Reporting struct {
	ExcludeDefaults string   `xml:"excludeDefaults,omitempty"`
	OutputDirectory string   `xml:"outputDirectory,omitempty"`
	Plugins         *Plugins `xml:"plugins,omitempty"`
}

type ReportSets struct {
	ReportSet []ReportSet `xml:"reportSet,omitempty"`
}

type ReportPlugin struct {
	Comment       string      `xml:",comment"`
	GroupId       string      `xml:"groupId,omitempty"`
	ArtifactId    string      `xml:"artifactId,omitempty"`
	Version       string      `xml:"version,omitempty"`
	ReportSets    *ReportSets `xml:"reportSets,omitempty"`
	Inherited     string      `xml:"inherited,omitempty"`
	Configuration *[]Any      `xml:"configuration,omitempty"`
}

type Plugin struct {
	Comment       string      `xml:",comment"`
	GroupId       string      `xml:"groupId,omitempty"`
	ArtifactId    string      `xml:"artifactId,omitempty"`
	Version       string      `xml:"version,omitempty"`
	Extensions    string      `xml:"extensions,omitempty"`
	Executions    *Executions `xml:"executions,omitempty"`
	Inherited     string      `xml:"inherited,omitempty"`
	Configuration *[]Any      `xml:"configuration,omitempty"`
}

type Reports struct {
	Report []string `xml:"report,omitempty"`
}

type ReportSet struct {
	Id            string   `xml:"id,omitempty"`
	Reports       *Reports `xml:"reports,omitempty"`
	Inherited     string   `xml:"inherited,omitempty"`
	Configuration *[]Any   `xml:"configuration,omitempty"`
}

type Profile struct {
	Id                     string                  `xml:"id,omitempty"`
	Activation             *Activation             `xml:"activation,omitempty"`
	Build                  *BuildBase              `xml:"build,omitempty"`
	Modules                *Modules                `xml:"modules,omitempty"`
	DistributionManagement *DistributionManagement `xml:"distributionManagement,omitempty"`
	Properties             *[]Any                  `xml:"properties,omitempty"`
	DependencyManagement   *DependencyManagement   `xml:"dependencyManagement,omitempty"`
	Dependencies           *Dependencies           `xml:"dependencies,omitempty"`
	Repositories           *Repositories           `xml:"repositories,omitempty"`
	PluginRepositories     *PluginRepositories     `xml:"pluginRepositories,omitempty"`
	Reports                *[]Any                  `xml:"reports,omitempty"`
	Reporting              *Reporting              `xml:"reporting,omitempty"`
}

type Activation struct {
	ActiveByDefault bool                `xml:"activeByDefault,omitempty"`
	Jdk             string              `xml:"jdk,omitempty"`
	Os              *ActivationOS       `xml:"os,omitempty"`
	Property        *ActivationProperty `xml:"property,omitempty"`
	File            *ActivationFile     `xml:"file,omitempty"`
}

type ActivationProperty struct {
	Name  string `xml:"name,omitempty"`
	Value string `xml:"value,omitempty"`
}

type ActivationFile struct {
	Missing string `xml:"missing,omitempty"`
	Exists  string `xml:"exists,omitempty"`
}

type ActivationOS struct {
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

type Resources struct {
	Resource []Resource `xml:"resource,omitempty"`
}

type TestResources struct {
	TestResource []Resource `xml:"testResource,omitempty"`
}

type Filters struct {
	Filter []string `xml:"filter,omitempty"`
}

type BuildBase struct {
	DefaultGoal      string            `xml:"defaultGoal,omitempty"`
	Resources        *Resources        `xml:"resources,omitempty"`
	TestResources    *TestResources    `xml:"testResources,omitempty"`
	Directory        string            `xml:"directory,omitempty"`
	FinalName        string            `xml:"finalName,omitempty"`
	Filters          *Filters          `xml:"filters,omitempty"`
	PluginManagement *PluginManagement `xml:"pluginManagement,omitempty"`
	Plugins          *Plugins          `xml:"plugins,omitempty"`
}

type Executions struct {
	Comment   string      `xml:",comment"`
	Text      string      `xml:",chardata"`
	Execution []Execution `xml:"execution,omitempty"`
}

type Execution struct {
	Text          string `xml:",chardata"`
	Id            string `xml:"id,omitempty"`
	Phase         string `xml:"phase,omitempty"`
	Goals         *[]Any `xml:"goals,omitempty"`
	Inherited     string `xml:"inherited,omitempty"`
	Configuration *[]Any `xml:"configuration,omitempty"`
}

type Includes struct {
	Include []string `xml:"include,omitempty"`
}

type Excludes struct {
	Exclude []string `xml:"exclude,omitempty"`
}

type Resource struct {
	Comment    string    `xml:",comment"`
	TargetPath string    `xml:"targetPath,omitempty"`
	Filtering  string    `xml:"filtering,omitempty"`
	Directory  string    `xml:"directory,omitempty"`
	Includes   *Includes `xml:"includes,omitempty"`
	Excludes   *Excludes `xml:"excludes,omitempty"`
}

type PluginManagement struct {
	Plugins *Plugins `xml:"plugins,omitempty"`
}

type Prerequisites struct {
	Maven string `xml:"maven,omitempty"`
}

type Extensions struct {
	Extension []Extension `xml:"extension,omitempty"`
}

type Build struct {
	Comment               string            `xml:",comment"`
	SourceDirectory       string            `xml:"sourceDirectory,omitempty"`
	ScriptSourceDirectory string            `xml:"scriptSourceDirectory,omitempty"`
	TestSourceDirectory   string            `xml:"testSourceDirectory,omitempty"`
	OutputDirectory       string            `xml:"outputDirectory,omitempty"`
	TestOutputDirectory   string            `xml:"testOutputDirectory,omitempty"`
	Extensions            *Extensions       `xml:"extensions,omitempty"`
	DefaultGoal           string            `xml:"defaultGoal,omitempty"`
	Resources             *Resources        `xml:"resources,omitempty"`
	TestResources         *TestResources    `xml:"testResources,omitempty"`
	Directory             string            `xml:"directory,omitempty"`
	FinalName             string            `xml:"finalName,omitempty"`
	Filters               *Filters          `xml:"filters,omitempty"`
	PluginManagement      *PluginManagement `xml:"pluginManagement,omitempty"`
	Plugins               *Plugins          `xml:"plugins,omitempty"`
}

type Extension struct {
	GroupId    string `xml:"groupId,omitempty"`
	ArtifactId string `xml:"artifactId,omitempty"`
	Version    string `xml:"version,omitempty"`
}

type Any struct {
	Comment     string `xml:",comment"`
	XMLName     xml.Name
	Value       string `xml:",chardata"`
	AnyElements []Any  `xml:",any"`
}
