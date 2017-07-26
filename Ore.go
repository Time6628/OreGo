package OreGo

import (
	"time"
	"encoding/json"
	"bytes"
	"github.com/valyala/fasthttp"
)

func New() (ore *Ore){
	ore = NewWithClient(fasthttp.Client{ReadTimeout: time.Second * 10})
	return
}

func NewWithClient(client fasthttp.Client) (ore *Ore) {
	ore = &Ore{
		Client:client,
	}
	return
}

func (ore *Ore) Request(url string) (body []byte) {
	stat, body, err := ore.Client.Get(nil, url)
	if err != nil || stat != 200 {
		return
	}
	return
}

func (ore *Ore) GetProjects() (projects *[]Project, err error) {
	body := ore.Request(projectsUrl)
	projects = &[]Project{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(projects)
	if err != nil {
		panic(err)
	}
	return
}

func (ore *Ore) GetProject(name string) (project *Project, err error) {
	project = &Project{}
	body := ore.Request(projectsUrl + name)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(project)
	if err != nil {
		panic(err)
	}
	return
}

func (ore *Ore) GetProjectVersions(name string) (versions *[]Version, err error) {
	body := ore.Request(projectsUrl + name + "/versions")
	versions = &[]Version{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(versions)
	if err != nil {
		panic(err)
	}
	return
}

func (ore *Ore) GetProjectVersion(name string, ver string) (version *Version, err error) {
	version = &Version{}
	body := ore.Request(projectsUrl + name + "/versions/" + ver)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(version)
	if err != nil {
		panic(err)
	}
	return
}

func (ore *Ore) GetRecommendedDownload(name string) string {
	project, err := ore.GetProject(name)
	if err != nil {
		return "An error occurred, could not obtain the url."
	}
	return project.RecommendedVersion.HRef + "/download"
}

func (ore *Ore) SearchProjects(query string) (projects *[]Project, err error) {
	body := ore.Request(projectsUrl + "?q=" + query)
	projects = &[]Project{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(projects)
	if err != nil {
		panic(err)
	}
	return
}

func (ore *Ore) GetUsers() (users *[]User, err error) {
	body := ore.Request(usersUrl)
	users = &[]User{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(users)
	if err != nil {
		panic(err)
	}
	return
}

func (ore *Ore) GetUser(username string) (user *User, err error) {
	user = &User{}
	body := ore.Request(usersUrl + username)
	err = json.NewDecoder(bytes.NewReader(body)).Decode(user)
	if err != nil {
		panic(err)
	}
	return
}
