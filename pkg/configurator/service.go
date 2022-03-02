package configurator

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/wamphlett/dot-collections/pkg/core/collections"
)

const (
	configFile       = "collection.yaml"
	varFile          = ".vars"
	varFileDelimiter = '|'
)

type Service struct {
	installPath string
}

func New(installPath string) *Service {
	return &Service{
		installPath: installPath,
	}
}

func (s *Service) LoadCollectionFromConfig(slug string) (*collections.Collection, error) {
	// Load the collection data if it is present
	collectionPath := s.getCollectionPath(slug)
	configPath := filepath.Join(collectionPath, configFile)

	c := &collection{}
	// If the collection has a collection file, load the data
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		data, err := os.ReadFile(configPath)
		if err != nil {
			return nil, err
		}
		yaml.Unmarshal(data, c)
	}

	collection := &collections.Collection{
		Slug:         slug,
		Path:         collectionPath,
		Description:  c.Description,
		Bootstrapper: c.Bootstrapper,
		Variables:    make([]*collections.Variable, len(c.Variables)),
	}

	for i, variable := range c.Variables {
		collection.Variables[i] = &collections.Variable{
			Key:     variable.Key,
			Default: variable.Default,
			IsEnv:   variable.IsEnv,
		}
	}

	if err := s.setVarsFromFile(collection); err != nil {
		return nil, err
	}

	return collection, nil
}

func (s *Service) WriteVariableFile(collection *collections.Collection) error {
	vars := make([][]string, len(collection.Variables))
	for i, variable := range collection.Variables {
		export := "false"
		if variable.IsEnv {
			export = "true"
		}
		vars[i] = []string{variable.Key, variable.Value, export}
	}

	variableFile := filepath.Join(s.getCollectionPath(collection.Slug), varFile)
	f, err := os.OpenFile(variableFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer f.Close()
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	w.Comma = varFileDelimiter
	if err := w.WriteAll(vars); err != nil {
		return err
	}

	return nil
}

func (s *Service) setVarsFromFile(collection *collections.Collection) error {
	collectionPath := s.getCollectionPath(collection.Slug)
	variableFile := filepath.Join(collectionPath, varFile)
	if _, err := os.Stat(variableFile); os.IsNotExist(err) {
		return nil
	}

	f, err := os.Open(variableFile)
	defer f.Close()
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	r.Comma = varFileDelimiter
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, variable := range collection.Variables {
		for _, record := range records {
			if variable.Key == record[0] {
				variable.Value = record[1]
				break
			}
		}
	}

	return nil
}

func (s *Service) getCollectionPath(slug string) string {
	return filepath.Join(s.installPath, slug)
}
