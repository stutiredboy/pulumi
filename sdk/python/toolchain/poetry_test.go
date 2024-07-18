package toolchain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDependenciesFromRequirementsTxt(t *testing.T) {
	t.Parallel()

	b := `
pulumi>=3.0.0,<4.0.0
requests>1

# Comment
setuptools    # comment here

	spaces-before  ==   1.2.3
`
	r := strings.NewReader(b)
	deps, err := dependenciesFromRequirementsTxt(r)
	require.NoError(t, err)
	require.Equal(t, map[string]string{
		"pulumi":        ">=3.0.0,<4.0.0",
		"requests":      ">1",
		"python":        "^3.8",
		"setuptools":    "*",
		"spaces-before": "1.2.3",
	}, deps)
}

func TestGeneratePyProjectTOML(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	p, err := newPoetry(dir)
	require.NoError(t, err)
	deps := map[string]string{
		"pulumi":        ">=3.0.0,<4.0.0",
		"requests":      ">1",
		"setuptools":    "*",
		"spaces-before": "1.2.3",
	}
	s, err := p.generatePyProjectTOML(deps)
	require.NoError(t, err)
	require.Equal(t, `[build-system]
requires = ["poetry-core"]
build-backend = "poetry.core.masonry.api"

[tool]
[tool.poetry]
package-mode = false
[tool.poetry.dependencies]
pulumi = ">=3.0.0,<4.0.0"
requests = ">1"
setuptools = "*"
spaces-before = "1.2.3"
`, s)
}

func TestCheckVersion(t *testing.T) {
	t.Parallel()
	require.NoError(t, validateVersion("Poetry (version 1.8.3)"))
	require.NoError(t, validateVersion("Poetry (version 2.1.2)"))
	require.NoError(t, validateVersion("Poetry (version 3.0)"))
	require.NoError(t, validateVersion("Poetry (version 1.9.0.dev0)"))
	require.ErrorContains(t, validateVersion("Poetry (version 1.7.0)"), "is less than the minimum required version")
	require.ErrorContains(t, validateVersion("invalid version string"), "unexpected output from poetry --version")
	require.ErrorContains(t, validateVersion(""), "unexpected output from poetry --version")
}
