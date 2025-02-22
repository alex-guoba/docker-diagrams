package node

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/firebase"
	"github.com/blushft/go-diagrams/nodes/gcp"
	"github.com/blushft/go-diagrams/nodes/programming"

	"github.com/alex-guoba/docker-diagrams/pkg/image"
)

var nodeMapping = map[string]func(opts ...diagram.NodeOption) *diagram.Node{
	"mysql":     apps.Database.Mysql,
	"postgres":  apps.Database.Postgresql,
	"oracle":    apps.Database.Oracle,
	"mssql":     apps.Database.Mssql,
	"mongodb":   apps.Database.Mongodb,
	"cassandra": apps.Database.Cassandra,

	// in-memmory node
	"redis":      apps.Inmemory.Redis,
	"memcached":  apps.Inmemory.Memcached,
	"kong":       apps.Network.Kong,
	"envoy":      apps.Network.Envoy,
	"traefik":    apps.Network.Traefik,
	"nginx":      apps.Network.Nginx,
	"envoyproxy": apps.Network.Envoy,
	"apache":     apps.Network.Apache,
	"caddy":      apps.Network.Caddy,
	"haproxy":    apps.Network.Haproxy,
	"zookeeper":  apps.Network.Zookeeper,

	// firebase
	"realtime": firebase.Develop.RealtimeDatabase,
	"gotrue":   firebase.Develop.Authentication,

	"prometheus": apps.Monitoring.Prometheus,

	"go":     programming.Language.Go,
	"python": programming.Language.Python,
	"java":   programming.Language.Java,
	"js":     programming.Language.Javascript,
	"ruby":   programming.Language.Ruby,
	"nodejs": programming.Language.Nodejs,
	"php":    programming.Language.Php,
}

func nodeByName(name string) *diagram.Node {
	var node *diagram.Node
	if nodeFunc, ok := nodeMapping[name]; ok {
		node = nodeFunc(
			diagram.Width(2.0),
			diagram.Height(2.0),
		)
	} else {
		// default to gcp compute engine
		node = gcp.Compute.ComputeEngine(
			diagram.Width(2.0),
			diagram.Height(2.0),
		)
	}
	// node.Options.Font.Size = 13
	// node.Options.FixedSize = false
	return node
}

func ImageToNode(serviceName string, imageName string, iconName string) *diagram.Node {
	if iconName != "" {
		return nodeByName(iconName)
	}

	name, err := image.ExtractImageName(imageName)
	if err != nil {
		name = serviceName
	}

	return nodeByName(name)
}
