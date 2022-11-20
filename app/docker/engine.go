package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/rs/zerolog/log"
)

func getController() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить docker контроллер")
	}
	return cli
}

func ListContainer() []string {
	cli := getController()
	list, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось получить список контейнеров")
	}
	var names_idx []string
	var str_container string
	for _, container := range list {
		str_container = "docker_" + container.ID[:10] + "_" + "_" + container.Image
		log.Info().Str("найден контейнер", str_container)
		names_idx = append(names_idx, str_container)

	}
	return names_idx
}
