package images

import (
	"github.com/containers/podman-tui/pdcs/registry"
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/errorhandling"
	"github.com/rs/zerolog/log"
)

// Remove removes the specified image ID.
func Remove(id string) ([]string, error) {
	log.Debug().Msgf("pdcs: podman image rm %s", id)

	var report []string

	ids := []string{id}

	conn, err := registry.GetConnection()
	if err != nil {
		return report, err
	}

	response, errs := images.Remove(conn, ids, new(images.RemoveOptions))
	if len(errs) > 0 {
		return report, errorhandling.JoinErrors(errs)
	}

	report = append(report, response.Deleted...)
	report = append(report, response.Untagged...)

	return report, nil
}
