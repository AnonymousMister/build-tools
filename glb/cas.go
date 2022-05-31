package glb

var IsDebug bool = false

var Con = &Context{}

func init() {
	if IsDebug {
		Con.Docker = &DockerContext{
			Tags:           []string{"11b606e7-dslyjava", "aaa-dslyjava"},
			DockerRegistry: "registry.cn-hangzhou.aliyuncs.com/winjoin/scygkj",
		}
	}
}

type Context struct {
	Docker *DockerContext
}

type DockerContext struct {
	DockerRegistry string
	Tags           []string
}
