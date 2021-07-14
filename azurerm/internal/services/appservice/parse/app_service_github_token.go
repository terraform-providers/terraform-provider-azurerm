package parse

type AppServiceGitHubTokenId struct{}

const gitHubTokenId = "/providers/Microsoft.Web/sourcecontrols/GitHub"

func (id AppServiceGitHubTokenId) String() string {
	return ""
}

func (id AppServiceGitHubTokenId) ID() string {
	return gitHubTokenId
}
