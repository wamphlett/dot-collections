package collections

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Service struct {
	installPath  string
	configurator Configurator
	collections  []*Collection
}

type Configurator interface {
	LoadCollectionFromConfig(slug string) (*Collection, error)
	WriteVariableFile(collection *Collection) error
}

func New(installPath string, configurator Configurator) *Service {
	return &Service{
		installPath:  installPath,
		collections:  []*Collection{},
		configurator: configurator,
	}
}

func (s *Service) Get(slug string) (*Collection, error) {
	// check that the given collection exists
	if !s.collectionExists(slug) {
		return nil, errors.New(fmt.Sprintf("unknown collection: %s", slug))
	}
	return s.configurator.LoadCollectionFromConfig(slug)
}

func (s *Service) GetAll() ([]*Collection, error) {
	if len(s.collections) != 0 {
		return s.collections, nil
	}

	slugs, err := s.getCollectionSlugs()
	if err != nil {
		return nil, err
	}

	collections := make([]*Collection, len(slugs))
	for i, slug := range slugs {
		collection, err := s.Get(slug)
		if err != nil {
			// TODO print a warning to the terminal
			continue
		}
		collections[i] = collection
	}

	return collections, nil
}

func (s *Service) UpdateVariables(collection *Collection) error {
	return s.configurator.WriteVariableFile(collection)
}

func (s *Service) getCollectionSlugs() ([]string, error) {
	var slugs []string
	err := filepath.Walk(s.installPath, func(p string, info fs.FileInfo, err error) error {
		if info.IsDir() && s.installPath != p {
			if err != nil {
				return err
			}
			slugs = append(slugs, info.Name())
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return slugs, nil
}

func (s *Service) collectionExists(slug string) bool {
	_, err := os.Stat(filepath.Join(s.installPath, slug))
	return !os.IsNotExist(err)
}

func IsValidSlug(slug string) bool {
	return true
}
