package image

import (
	"errors"
	"regexp"
	"strings"
)

const (
	TagIgnore = "docker-diagram.ignore"
	TagGroup  = "docker-diagram.group"
	TagIcon   = "docker-diagram.icon"
)

func ExtractImageName(image string) (string, error) {
	re := regexp.MustCompile(`^(?:([\w.-]+\/)?(?:[\w.-]+\/)?)?([\w.-]+)(?::([\w.-]+)|@([\w:-]+))?$`)
	match := re.FindStringSubmatch(image)
	if match == nil {
		return "", errors.New(`invalid image format: "${image}"`)
	}

	// extract service name, strip "library/" prefix if present
	serviceName := match[2]
	if strings.HasPrefix(serviceName, "library/") {
		serviceName = serviceName[len("library/"):]
	}
	return serviceName, nil
}
