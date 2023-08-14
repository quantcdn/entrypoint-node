package main

import (
	"os"
	"os/exec"

	"github.com/alecthomas/kingpin"
	"github.com/quantcdn/entrypoint-node/internal/backend"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	app   = kingpin.New("entrypoint-node", "Docker entrypoint to start Node applications.")
	cmds  = kingpin.Arg("commands", "Node JS commands to execute.").Strings()
	dir   = kingpin.Flag("dir", "Directory to execute node commands in.").String()
	url   = kingpin.Flag("url", "The backend url.").Envar("NEXT_PUBLIC_DRUPAL_BASE_URL").String()
	retry = kingpin.Flag("retry", "Times to retry the backend connection.").Default("10").Envar("BACKEND_RETRY").Int()
	delay = kingpin.Flag("delay", "Delay between backend requests.").Default("5").Envar("BACKEND_DELAY").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Str("directory", *dir).Strs("commands", *cmds).Msg("attempting to start application")

	if *url != "" {
		log.Info().Str("backend", *url).Int("retry", *retry).Int("delay", *delay).Msg("validating connectivity")
		// We need to determine connectivity to the datasource.
		if !backend.Connect(*url, *delay, *retry) {
			log.Error().Msg("unable to connect to backend")
			os.Exit(1)
		}
		log.Info().Msg("backend connection ok")
	}

	for _, cmd := range *cmds {
		args := []string{"npm"}

		if *dir != "" {
			args = append(args, []string{"--prefix", *dir}...)
		}

		args = append(args, []string{"run", cmd}...)

		c := exec.Command(args[0], args[1:]...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			os.Exit(1)
		}
	}
}
