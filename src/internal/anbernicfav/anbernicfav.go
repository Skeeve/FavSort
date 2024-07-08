package anbernicfav

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	cli "github.com/urfave/cli/v3"
	yaml "gopkg.in/yaml.v2"
)

type FavEntry [6]string

type Config struct {
	internal struct {
		Favorites struct {
			Entry   string `yaml:"entry"`   // entry: ++ sort Favorites ++
			File    string `yaml:"file"`    // file: /mnt/data/misc/.favorite
			Version string `yaml:"version"` // version: Version=1
		} `yaml:"favorites"`
		Arcade struct {
			CSV string   `yaml:"csv"` // /mnt/vendor/bin/arcade-plus.csv
			Emu []string `yaml:"emu"` // List of emulator directories on the device
		} `yaml:"arcade"`
		isArcade      map[string]bool
		arcadeNames   map[string]string
		splittedEntry FavEntry
	}
}

func NewFavEntry(entryString string) (entry FavEntry, isEntry bool) {
	splitted := strings.Split(entryString, ":")
	if len(splitted) != 6 {
		return
	}
	return FavEntry(splitted), true
}

func (e FavEntry) String() string {
	return strings.Join(e[:], ":")
}

func (e *FavEntry) SetLine(lineNumber int) {
	e[4] = strconv.Itoa(lineNumber)
}

func (e FavEntry) GetName() string {
	return e[0]
}

func (e FavEntry) GetEmulator() string {
	return e[1]
}

func (e FavEntry) GetLocation() string {
	return e[3]
}

// GetConfig - reads the configuration
// and fills all require data structures
func GetConfig(filename string) (config *Config, err error) {
	// Get the absolute path to the configuration
	filename, err = filepath.Abs(filename)
	if err != nil {
		err = cli.Exit(fmt.Sprintf("Can't access path %s: %s", filename, err.Error()), 72 /* critical OS file missing */)
		return
	}
	// read it as one data block
	data, err := os.ReadFile(filename)
	if err != nil {
		err = cli.Exit(fmt.Sprintf("Can't read configuration %s: %s", filename, err.Error()), 78 /* configuration error */)
		return
	}

	// parse it as YAML
	var cfg Config
	err = yaml.Unmarshal(data, &cfg.internal)
	if err != nil {
		err = cli.Exit(fmt.Sprintf("Error parsing %s: %s", filename, err.Error()), 78 /* configuration error */)
		return
	}

	// initialize the isArcade slice
	cfg.internal.isArcade = make(map[string]bool, len(cfg.internal.Arcade.Emu))
	for _, emu := range cfg.internal.Arcade.Emu {
		cfg.internal.isArcade[emu] = true
	}

	// get the full names of all known arcade files
	cfg.internal.arcadeNames, err = getArcadeNames(cfg.internal.Arcade.CSV)

	// Try to find the starter.
	// It is an ".sh" file and has to be either in the same directory as the configuration
	// or in one of its parent directories.

	path := filepath.Dir(filename)
	for {
		_, err = os.Stat(filepath.Join(path, cfg.internal.Favorites.Entry+".sh"))
		if err == nil {
			break
		}
		newPath := filepath.Dir(path)
		if newPath == path {
			err = cli.Exit(
				fmt.Sprintf(
					"Couldn't find starter '%s.sh' anywhere in the path '%s'",
					cfg.internal.Favorites.Entry,
					filepath.Dir(filename),
				),
				70 /* internal software error */)
			return
		}
		path = newPath
	}

	splittedPath := strings.Split(filepath.ToSlash(path), "/")
	if len(splittedPath) < 4 {
		err = cli.Exit(
			fmt.Sprintf(
				"Installation error. Illegal path: '%s'",
				path,
			),
			78 /* configuration error */)
		return
	}
	location := "3"
	emuDir := splittedPath[4]
	if splittedPath[2] == "mmc" {
		location = "2"
	}
	cfg.internal.splittedEntry = FavEntry{
		cfg.internal.Favorites.Entry + ".sh",
		emuDir,
		"nul",
		location,
		"-", // will be fixed later
		"0",
	}

	config = &cfg
	return
}

// IsArcard will return whether an emulator is considered an arcade emulator
func (c *Config) IsArcade(emulator string) bool {
	return c.internal.isArcade[emulator]
}

// FullName will give the game's full name
func (c *Config) FullName(game string) string {
	return c.internal.arcadeNames[game]
}

// Favorites will give the filenam of the favorites file as defined in the configuration
func (c *Config) Starter() string {
	return c.internal.Favorites.Entry
}

// Favorites will give the filenam of the favorites file as defined in the configuration
func (c *Config) Favorites() string {
	return c.internal.Favorites.File
}

// FavVersion will return the default versio as defined in the configuration
func (c *Config) FavVersion() string {
	return c.internal.Favorites.Version
}

func (c *Config) FavEntry() FavEntry {
	return c.internal.splittedEntry
}

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, cli.Exit(fmt.Sprintf("Unable to read input file %s: %s", filePath, err.Error()), 72 /* critical OS file missing */)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, cli.Exit(fmt.Sprintf("Unable to parse CSV file %s: %s", filePath, err.Error()), 78 /* configuration error */)
	}

	return records, nil
}

func getArcadeNames(arcadeCsv string) (arcadeDirs map[string]string, err error) {
	records, err := readCsvFile(arcadeCsv)
	if err != nil {
		return
	}
	arcadeDirs = make(map[string]string)
	for _, rec := range records {
		if rec[2] != "" {
			arcadeDirs[rec[0]] = rec[2]
			/* //Debug code
				fmt.Printf("%s: %s\n", rec[0], rec[2])
			} else {
				fmt.Printf("%s: n/a\n", rec[0])
				//*/
		}
	}
	return
}
