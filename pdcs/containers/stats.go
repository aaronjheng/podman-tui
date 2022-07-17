package containers

import (
	"github.com/containers/podman-tui/pdcs/registry"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/rs/zerolog/log"
)

// Stats returns live stream of containers stats result.
func Stats(id string, opts *containers.StatsOptions) (chan entities.ContainerStatsReport, error) {
	log.Debug().Msgf("pdcs: podman container stats %s", id)

	conn, err := registry.GetConnection()
	if err != nil {
		return nil, err
	}

	statReportChan, err := containers.Stats(conn, []string{id}, opts)
	if err != nil {
		return nil, err
	}

	return statReportChan, nil
}
