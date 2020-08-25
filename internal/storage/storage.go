package storage

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/99designs/keyring"
	"github.com/PuerkitoBio/purell"
)

const purellFlags = purell.FlagsSafe | purell.FlagsUsuallySafeGreedy | purell.FlagRemoveDuplicateSlashes

type Keyring struct {
	Backend string          `mapstructure:"backend"`
	Config  keyring.Config  `mapstructure:"config"`
	Ring    keyring.Keyring `yaml:"-"`
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

func newKeyring(cfg keyring.Config) (keyring.Keyring, error) {
	if k, err := keyring.Open(cfg); err != nil {
		return nil, err
	} else {
		return k, nil
	}
}

// InitKeyring initializes the token storage into the main package var 'backendStorage'.
// It is not called during the early init phase to avoid errors with commands
// that do not need access to a backend. Instead, commands that interact with a backend
// should call InitKeyring and propagate errors back to the rootCmd.
func (k *Keyring) InitKeyring() error {
	var err error
	switch k.Backend {
	case "automatic", "":
		k.Config.AllowedBackends = Backends
	case "keychain":
		k.Config.AllowedBackends = []keyring.BackendType{keyring.KeychainBackend}
	case "kdewallet":
		k.Config.AllowedBackends = []keyring.BackendType{keyring.KWalletBackend}
	case "secret-service":
		k.Config.AllowedBackends = []keyring.BackendType{keyring.SecretServiceBackend}
	case "wincred":
		k.Config.AllowedBackends = []keyring.BackendType{keyring.WinCredBackend}
	case "keepass":
		k.Config.AllowedBackends = []keyring.BackendType{keyring.PassBackend}
	case "file":
		k.Config.AllowedBackends = []keyring.BackendType{keyring.FileBackend}
	default:
		return errors.New(fmt.Sprintf("Unknown backend '%s'", k.Backend))
	}

	if k.Ring, err = newKeyring(k.Config); err == nil {
		return nil
	} else {
		return err
	}
}

// Store saves the token in the token storage, returning any errors that occur while trying to
// persist the token.
func (k *Keyring) Store(token Token) error {
	vaultAddr := encodeVaultAddr(token.VaultAddr)

	item := keyring.Item{
		Key:         vaultAddr,
		Data:        []byte(token.Token),
		Label:       "Vault-token: " + decodeVaultAddr(vaultAddr),
		Description: "Vault-token: " + decodeVaultAddr(vaultAddr),
	}
	return k.Ring.Set(item)
}

// Get retrieves a token for the vaultAddr if one is available in the token store. A missing
// token is not an error. Errors are returned if there are errors communicating with the token store.
func (k *Keyring) Get(vaultAddr string) (Token, error) {
	vaultAddr = encodeVaultAddr(vaultAddr)
	if i, err := k.Ring.Get(vaultAddr); err != nil {
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
func (k *Keyring) List() ([]Token, error) {
	tokens := make([]Token, 0)

	list, err := k.Ring.Keys()
	if err != nil {
		return []Token{}, err
	}

	for _, i := range list {
		if t, err := k.Get(decodeVaultAddr(i)); err == nil {
			tokens = append(tokens, t)
		}
	}
	return tokens, nil
}

// Error erases the token for the vaultAddr from the token store. A missing token is not an error.
// Errors are returned if there are errors communicating with the token store.
func (k *Keyring) Erase(vaultAddr string) error {
	vaultAddr = encodeVaultAddr(vaultAddr)
	return k.Ring.Remove(vaultAddr)
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
