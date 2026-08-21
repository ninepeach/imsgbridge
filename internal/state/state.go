package state

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type State struct {
	LastMessageID int64 `json:"last_message_id"`
}

type Store struct {
	path string

	mu sync.Mutex
}

func NewStore() (*Store, error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	dir := filepath.Join(
		home,
		".imsgbridge",
	)

	err = os.MkdirAll(
		dir,
		0755,
	)

	if err != nil {
		return nil, err
	}

	return &Store{

		path: filepath.Join(
			dir,
			"state.json",
		),
	}, nil
}

func (s *Store) Load() (int64, error) {

	s.mu.Lock()

	defer s.mu.Unlock()

	data, err := os.ReadFile(
		s.path,
	)

	if os.IsNotExist(err) {

		return 0, nil
	}

	if err != nil {

		return 0, err
	}

	var state State

	err = json.Unmarshal(
		data,
		&state,
	)

	if err != nil {

		return 0, err
	}

	return state.LastMessageID, nil
}

func (s *Store) Save(
	id int64,
) error {

	s.mu.Lock()

	defer s.mu.Unlock()

	state := State{

		LastMessageID: id,
	}

	data, err := json.MarshalIndent(
		state,
		"",
		"  ",
	)

	if err != nil {

		return err
	}

	tmp := s.path + ".tmp"

	err = os.WriteFile(
		tmp,
		data,
		0644,
	)

	if err != nil {

		return err
	}

	return os.Rename(
		tmp,
		s.path,
	)
}
