package storage

import (
	"net/url"

	"github.com/99designs/keyring"
	"github.com/PuerkitoBio/purell"
)

const purellFlags = purell.FlagsSafe | purell.FlagsUsuallySafeGreedy | purell.FlagRemoveDuplicateSlashes

type Storage struct {
	config  keyring.Config
	keyring keyring.Keyring
}

var Backends = []keyring.BackendType{
	// Linux
	keyring.KWalletBackend,
	keyring.SecretServiceBackend,
	// MacOS
	keyring.KeychainBackend,
	// Windows
	keyring.WinCredBackend,
	// General
	keyring.PassBackend,
	keyring.FileBackend,
}

func New(cfg keyring.Config) (*Storage, error) {
	if kr, err := keyring.Open(cfg); err != nil {
		return nil, err
	} else {
		storage := &Storage{
			config:  cfg,
			keyring: kr,
		}
		return storage, nil
	}
}

// Store saves the token in the token storage, returning any errors that occur while trying to
// persist the token.
func (s *Storage) Store(token Token) error {
	vaultAddr := encodeVaultAddr(token.VaultAddr)

	item := keyring.Item{
		Key:         vaultAddr,
		Data:        []byte(token.Token),
		Label:       "Vault-token: " + decodeVaultAddr(vaultAddr),
		Description: "Vault-token: " + decodeVaultAddr(vaultAddr),
	}
	return s.keyring.Set(item)
}

// Get retrieves a token for the vaultAddr if one is available in the token store. A missing
// token is not an error. Errors are returned if there are errors communicating with the token store.
func (s *Storage) Get(vaultAddr string) (Token, error) {
	vaultAddr = encodeVaultAddr(vaultAddr)
	if i, err := s.keyring.Get(vaultAddr); err != nil {
		if err == keyring.ErrKeyNotFound {
			return Token{}, nil
		}
		return Token{}, err
	} else {
		t := Token{
			VaultAddr: decodeVaultAddr(vaultAddr),
			Token:     string(i.Data),
		}
		return t, nil
	}
}

// List retrieves all tokens available in the token storage.
// An empty store is not an error. Errors are returned if there are errors communicating
// with the token store.
func (s *Storage) List() ([]Token, error) {
	tokens := make([]Token, 0)

	list, err := s.keyring.Keys()
	if err != nil {
		return []Token{}, err
	}

	for _, i := range list {
		if t, err := s.Get(decodeVaultAddr(i)); err == nil {
			tokens = append(tokens, t)
		}
	}
	return tokens, nil
}

// Error erases the token for the vaultAddr from the token store. A missing token is not an error.
// Errors are returned if there are errors communicating with the token store.
func (s *Storage) Erase(vaultAddr string) error {
	vaultAddr = encodeVaultAddr(vaultAddr)
	return s.keyring.Remove(vaultAddr)
}

// AvailableBackends returns the available backends on this platform
func AvailableBackends() []string {
	backends := make([]string, 0)
	for _, i := range keyring.AvailableBackends() {
		for _, x := range Backends {
			if i == x {
				backends = append(backends, string(i))
			}
		}
	}
	return backends
}

func encodeVaultAddr(addr string) string {
	encoded, err := purell.NormalizeURLString(addr, purellFlags)
	if err != nil {
		return addr
	}
	encoded = url.PathEscape(encoded)

	return encoded
}

func decodeVaultAddr(encoded string) string {
	decoded, _ := url.PathUnescape(encoded)
	return decoded
}
