package networks

import (
	"github.com/containers/podman-tui/pdcs/registry"
	"github.com/containers/podman/v5/pkg/bindings/network"
	"github.com/containers/podman/v5/pkg/errorhandling"
	"github.com/rs/zerolog/log"
)

// Prune removes all unused network.
func Prune() error {
	var errorReport []error

	log.Debug().Msgf("pdcs: podman network remove")

	conn, err := registry.GetConnection()
	if err != nil {
		return err
	}

	response, err := network.Prune(conn, new(network.PruneOptions))
	if err != nil {
		return err
	}

	for _, r := range response {
		if r.Error != nil {
			errorReport = append(errorReport, r.Error)
		}
	}

	return errorhandling.JoinErrors(errorReport)
}
