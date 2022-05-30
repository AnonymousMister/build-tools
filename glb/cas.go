package glb

var IsDebug bool = false

var Con = &Context{}

type Context struct {
	Docker *DockerContext
}

type DockerContext struct {
	DockerRegistry string
	Tags           []string
}
