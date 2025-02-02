package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fred1268/go-clap/clap"
)

type Config struct {
	Help      bool     `clap:"--help,-h"`
	Directory string   `clap:"--directory,-d"`
	MinScore  float64  `clap:"--min-score,-s"`
	Images    []string `clap:"trailing"`
}

func main() {
	args := []string{}
	if len(os.Args) >= 2 {
		args = os.Args[1:]
	}

	config := Config{
		MinScore: 85,
		Help:     false,
	}
	if _, err := clap.Parse(args, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Failed parse cli arguments: %+v", err)
		os.Exit(1)
	}

	if config.Help {
		fmt.Print(`Rename your images based on their true identity (as seen by AI)
Usage: meme-tag [options] [...image_path]

Options:
  -h, --help              Print help
  -s, --min-score <score> Specify minimum score required in order for a word to be selected
  -d, --directory <path>  Specify directory to tag image
`)
		os.Exit(0)
	}

	if config.Directory == "" && len(config.Images) == 0 {
		fmt.Fprintf(os.Stderr, `No image or directory specified.
  $ meme-tag /path/to/image1 /path/to/image2 # for tagging specific images.
  $ meme-tag -d /path/to/image/dir' # for tagging all images in a directory.`)
		os.Exit(1)
	}

	for _, i := range config.Images {
		processImage(i, config.MinScore, true)
	}

	if config.Directory == "" {
		os.Exit(0)
	}

	fmt.Printf("All images in '%s' will be renamed.\n   Do you want to continue? (y/n): ", config.Directory)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input != "y" {
		os.Exit(0)
	}

	walkDirectory(config.Directory, config.MinScore)
}
