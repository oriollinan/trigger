package trigger

import (
	"trigger.com/trigger/pkg/action"
)

type GithubWorkspaceCtx string

const GithubCommitCtxKey GithubWorkspaceCtx = GithubWorkspaceCtx("GithubCommitCtxKey")

const userIdCtxKey GithubWorkspaceCtx = GithubWorkspaceCtx("userIdCtxKey")

type Service interface {
	action.Trigger
}

type Handler struct {
	Service
}

type Model struct {
}
