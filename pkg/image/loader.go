package image

import (
	"context"

	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/compose-spec/compose-go/v2/types"
)

func LoadProject(ctx context.Context, composeFile string, envFiles []string) (*types.Project, error) {
	options, err := cli.NewProjectOptions(
		[]string{composeFile},
		cli.WithOsEnv,
		cli.WithEnvFiles(envFiles...),
		cli.WithDotEnv,
	)
	if err != nil {
		return nil, err
	}

	project, err := options.LoadProject(ctx)
	if err != nil {
		return nil, err
	}
	return project, nil
}
