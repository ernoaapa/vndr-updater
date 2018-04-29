package vndr

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// ReadConfig read deps from the config file
func ReadConfig(configFile string) ([]DepEntry, error) {
	cfg, err := getConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to open config file: %v", err)
	}
	deps, err := ParseDeps(cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse config: %v", err)
	}
	return deps, nil
}

func getConfig(configFile string) (io.Reader, error) {
	if strings.HasPrefix(configFile, "http") {
		resp, err := http.Get(configFile)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		return buf, nil
	}

	cfg, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to open config file: %v", err)
	}
	return cfg, nil
}

// ParseDeps parses vndr.conf file
func ParseDeps(r io.Reader) ([]DepEntry, error) {
	var deps []DepEntry
	s := bufio.NewScanner(r)
	for s.Scan() {
		ln := strings.TrimSpace(s.Text())
		if strings.HasPrefix(ln, "#") || ln == "" {
			continue
		}
		cidx := strings.Index(ln, "#")
		if cidx > 0 {
			ln = ln[:cidx]
		}
		ln = strings.TrimSpace(ln)
		parts := strings.Fields(ln)
		if len(parts) != 2 && len(parts) != 3 {
			return nil, fmt.Errorf("invalid config format: %s", ln)
		}
		d := DepEntry{
			ImportPath: parts[0],
			Rev:        parts[1],
		}
		if len(parts) == 3 {
			d.RepoPath = parts[2]
		}
		deps = append(deps, d)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return deps, nil
}

// WriteConfig writes list of dependencies to config file
func WriteConfig(deps []DepEntry, cfgFile string) error {
	var lines []string
	for _, d := range deps {
		lines = append(lines, d.String())
	}
	return ioutil.WriteFile(cfgFile, []byte(strings.Join(lines, "")), os.FileMode(0666))
}
