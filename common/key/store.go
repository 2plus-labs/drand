package key

import (
	"fmt"
	"os"
	"path"
	"reflect"

	"github.com/BurntSushi/toml"

	"github.com/drand/drand/common"
	"github.com/drand/drand/internal/fs"
)

// Store abstracts the loading and saving of any private/public cryptographic
// material to be used by drand. For the moment, only a file based store is
// implemented.
type Store interface {
	// SaveKeyPair saves the private key generated by drand as well as the
	// public identity key associated
	SaveKeyPair(p *Pair) error
	// LoadKeyPair loads the private/public key pair associated with the drand
	// operator
	LoadKeyPair() (*Pair, error)
	SaveShare(share *Share) error
	LoadShare() (*Share, error)
	SaveGroup(*Group) error
	LoadGroup() (*Group, error)
	Reset(...ResetOption) error
}

// FolderName is the name of the folder where drand keeps its keys
const FolderName = "key"

// GroupFolderName is the name of the folder where drand keeps its group files
const GroupFolderName = "groups"
const keyFileName = "drand_id"
const privateExtension = ".private"
const publicExtension = ".public"
const groupFileName = "drand_group.toml"
const shareFileName = "dist_key.private"
const distKeyFileName = "dist_key.public"

// Tomler represents any struct that can be (un)marshaled into/from toml format
type Tomler interface {
	TOML() interface{}
	FromTOML(i interface{}) error
	TOMLValue() interface{}
}

// fileStore is a Store using filesystem to store information
type fileStore struct {
	baseFolder     string
	beaconID       string
	privateKeyFile string
	publicKeyFile  string
	shareFile      string
	distKeyFile    string
	groupFile      string
}

// GetFirstStore will return the first store from the stores map
func GetFirstStore(stores map[string]Store) (string, Store) {
	for k, v := range stores {
		return k, v
	}
	return "", nil
}

// NewFileStores will list all folder on base path and load every file store it can find. It will
// return a map with a beacon id as key and a file store as value.
func NewFileStores(baseFolder string) (map[string]Store, error) {
	fileStores := make(map[string]Store)
	fi, err := os.ReadDir(path.Join(baseFolder))
	if err != nil {
		return nil, err
	}

	for _, f := range fi {
		if f.IsDir() {
			fileStores[f.Name()] = NewFileStore(baseFolder, f.Name())
		}
	}

	if len(fileStores) == 0 {
		fileStores[common.DefaultBeaconID] = NewFileStore(baseFolder, common.DefaultBeaconID)
	}

	return fileStores, nil
}

// NewFileStore is used to create the config folder and all the subfolders.
// If a folder already exists, we simply check the rights
func NewFileStore(baseFolder, beaconID string) Store {
	beaconID = common.GetCanonicalBeaconID(beaconID)

	store := &fileStore{baseFolder: baseFolder, beaconID: beaconID}

	keyFolder := fs.CreateSecureFolder(path.Join(baseFolder, beaconID, FolderName))
	groupFolder := fs.CreateSecureFolder(path.Join(baseFolder, beaconID, GroupFolderName))

	store.privateKeyFile = path.Join(keyFolder, keyFileName) + privateExtension
	store.publicKeyFile = path.Join(keyFolder, keyFileName) + publicExtension
	store.groupFile = path.Join(groupFolder, groupFileName)
	store.shareFile = path.Join(groupFolder, shareFileName)
	store.distKeyFile = path.Join(groupFolder, distKeyFileName)

	return store
}

// SaveKeyPair first saves the private key in a file with tight permissions and then
// saves the public part in another file.
func (f *fileStore) SaveKeyPair(p *Pair) error {
	if err := Save(f.privateKeyFile, p, true); err != nil {
		return err
	}
	fmt.Printf("Saved the key : %s at %s\n", p.Public.Addr, f.publicKeyFile) //nolint
	return Save(f.publicKeyFile, p.Public, false)
}

// LoadKeyPair decode private key first then public
func (f *fileStore) LoadKeyPair() (*Pair, error) {
	p := new(Pair)
	if err := Load(f.privateKeyFile, p); err != nil {
		return nil, err
	}
	return p, Load(f.publicKeyFile, p.Public)
}

func (f *fileStore) LoadGroup() (*Group, error) {
	g := new(Group)
	err := Load(f.groupFile, g)
	if err != nil {
		return nil, err
	}
	// we don't want to return a pointer to an empty `Group` struct if
	// there isn't a group in the file system
	//nolint:nilnil
	if g == new(Group) {
		return nil, nil
	}
	return g, nil
}

func (f *fileStore) SaveGroup(g *Group) error {
	return Save(f.groupFile, g, false)
}

func (f *fileStore) SaveShare(share *Share) error {
	fmt.Printf("crypto store: saving private share in %s\n", f.shareFile) //nolint
	return Save(f.shareFile, share, true)
}

func (f *fileStore) LoadShare() (*Share, error) {
	s := new(Share)
	return s, Load(f.shareFile, s)
}

func (f *fileStore) Reset(_ ...ResetOption) error {
	if err := Delete(f.distKeyFile); err != nil {
		return fmt.Errorf("drand: err deleting dist. key file: %w", err)
	}
	if err := Delete(f.shareFile); err != nil {
		return fmt.Errorf("drand: err deleting share file: %w", err)
	}

	if err := Delete(f.groupFile); err != nil {
		return fmt.Errorf("drand: err deleting group file: %w", err)
	}
	return nil
}

// Save the given Tomler interface to the given path. If secure is true, the
// file will have a 0700 security.
// TODO: move that to fs/
func Save(filePath string, t Tomler, secure bool) error {
	var fd *os.File
	var err error
	if secure {
		fd, err = fs.CreateSecureFile(filePath)
	} else {
		fd, err = os.Create(filePath)
	}
	if err != nil {
		return fmt.Errorf("config: can't save %s to %s: %w", reflect.TypeOf(t).String(), filePath, err)
	}
	defer fd.Close()
	return toml.NewEncoder(fd).Encode(t.TOML())
}

// Load the given Tomler from the given file path.
func Load(filePath string, t Tomler) error {
	tomlValue := t.TOMLValue()
	var err error
	if _, err = toml.DecodeFile(filePath, tomlValue); err != nil {
		return err
	}
	return t.FromTOML(tomlValue)
}

// Delete the resource denoted by the given path. If it is a file, it deletes
// the file; if it is a folder it delete the folder and all its content.
func Delete(filePath string) error {
	return os.RemoveAll(filePath)
}

// ResetOption is an option to allow for fine-grained reset
// operations
// TODO - unclear what needs to be done here.
type ResetOption int
