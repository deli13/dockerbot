package docker

import (
	"context"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/rs/zerolog/log"
)

type DockerList struct {
	Id   string
	Name string
}

type DockerInfo struct {
	Id     string
	Name   string
	Image  string
	Status string
	Stats  string
}

func getController() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось запустить docker контроллер")
	}
	return cli
}

func ListContainer() []DockerList {
	cli := getController()
	list, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось получить список контейнеров")
	}
	var names_idx []DockerList
	var str_container string
	for _, container := range list {
		log.Info().Str("найден контейнер", str_container)
		names_idx = append(names_idx, DockerList{Name: strings.Join(container.Names, "-"), Id: container.ID[:10]})

	}
	return names_idx
}

func ContainerInfo(id string) DockerInfo {
	cli := getController()
	container, err := GetContainer(id)
	if err != nil {
		log.Err(err).Msg("Не удалось получить контейнер")
		return DockerInfo{}
	}
	if container.ID == "" {
		return DockerInfo{}
	}
	var stats_string string
	stats, err := cli.ContainerStatsOneShot(context.Background(), container.ID)
	defer stats.Body.Close()
	if err == nil {

		str_byte, err := ioutil.ReadAll(stats.Body)
		if err != nil {
			log.Err(err).Msg("Не удалось прочитать статистику")
		}
		stats_string = string(str_byte)
	}

	return DockerInfo{
		Id:     container.ID[:10],
		Name:   strings.Join(container.Names, "-"),
		Status: container.Status,
		Image:  container.Image,
		Stats:  stats_string,
	}
}

func GetContainer(containerId string) (types.Container, error) {
	cli := getController()
	list, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось получить список контейнеров")
		return types.Container{}, errors.New("Не удалось получить список контейнеров")
	}
	for _, cont := range list {
		if cont.ID[:10] == containerId {
			return cont, nil
		}
	}
	return types.Container{}, errors.New("Контейнер не найден")
}

func StartContainer(container string) error {
	cli := getController()
	return cli.ContainerStart(context.Background(), container, types.ContainerStartOptions{})
}

func StopContainer(container string) error {
	cli := getController()
	return cli.ContainerStop(context.Background(), container, nil)
}

func RestartContainer(container string) error {
	cli := getController()
	return cli.ContainerRestart(context.Background(), container, nil)
}

// func ContainerLogs(id string) {
// 	cli := getController()
// 	logs, err := cli.ContainerLogs(context.WithTimeout(context.Background(), 3*time.Second), container.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("Не удалось прочитать логи")
// 	}
// }
