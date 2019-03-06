package beater

import (
	"strconv"
	"strings"
)

// ZFS zpool states, which can indicate if a pool is online, offline, degraded, etc
const (
	ZpoolOnline   = "ONLINE"
	ZpoolDegraded = "DEGRADED"
	ZpoolFaulted  = "FAULTED"
	ZpoolOffline  = "OFFLINE"
	ZpoolUnavail  = "UNAVAIL"
	ZpoolRemoved  = "REMOVED"
)

// Zpool is a ZFS zpool.  A pool is a top-level structure in ZFS, and can
// contain many descendent datasets.
type Zpool struct {
	Name                     string
	Size                     uint64
	Capacity                 uint64
	Altroot                  string
	Health                   string
	GUID                     string
	Version                  string
	Bootfs                   string
	Delegation               string
	Autoreplace              string
	Cachefile                string
	Failmode                 string
	Listsnapshots            string
	Autoexpand               string
	Dedupditto               string
	Dedupratio               float64
	Free                     uint64
	Allocated                uint64
	Readonly                 string
	Ashift                   uint64
	Comment                  string
	Expandsize               uint64
	Freeing                  uint64
	Fragmentation            uint64
	Leaked                   uint64
	FeatureAsyncDestroy      string
	FeatureEmptyBpobj        string
	FeatureLz4Compress       string
	FeatureSpacemapHistogram string
	FeatureEnabledTxg        string
	FeatureHoleBirth         string
	FeatureExtensibleDataset string
	FeatureEmbeddedData      string
	FeatureBookmarks         string
	FeatureFilesystemLimits  string
	FeatureLargeBlocks       string
}

// List of Zpool properties to retrieve from zpool list command
var zpoolPropList = []string{"name", "size", "capacity", "altroot", "health", "guid", "version", "bootfs", "delegation", "autoreplace", "cachefile", "failmode", "listsnapshots", "autoexpand", "dedupditto", "dedupratio", "free", "allocated", "readonly", "ashift", "comment", "expandsize", "freeing", "fragmentation", "leaked", "feature@async_destroy", "feature@empty_bpobj", "feature@lz4_compress", "feature@spacemap_histogram", "feature@enabled_txg", "feature@hole_birth", "feature@extensible_dataset", "feature@embedded_data", "feature@bookmarks", "feature@filesystem_limits", "feature@large_blocks"}
var zpoolPropListOptions = strings.Join(zpoolPropList, ",")
var zpoolArgs = []string{"get", "-p", zpoolPropListOptions}

//var zpoolArgs = []string{"get", zpoolPropListOptions}

// zpool is a helper function to wrap typical calls to zpool.
func zpool(arg ...string) ([][]string, error) {
	c := command{Command: "zpool"}
	return c.Run(arg...)
}

// ListZpools list all ZFS zpools accessible on the current system.
func ListZpools() ([]*Zpool, error) {
	args := []string{"list", "-Ho", "name"}
	out, err := zpool(args...)
	if err != nil {
		return nil, err
	}

	var pools []*Zpool

	for _, line := range out {
		z, err := GetZpool(line[0])
		if err != nil {
			return nil, err
		}
		pools = append(pools, z)
	}
	return pools, nil
}

// GetZpool retrieves a single ZFS zpool by name.
func GetZpool(name string) (*Zpool, error) {
	args := zpoolArgs
	args = append(args, name)
	out, err := zpool(args...)
	if err != nil {
		return nil, err
	}

	// there is no -H
	out = out[1:]

	z := &Zpool{Name: name}
	for _, line := range out {
		if err := z.parseLine(line); err != nil {
			return nil, err
		}
	}

	return z, nil
}

func (z *Zpool) parseLine(line []string) error {
	prop := line[1]
	val := line[2]

	var err error

	switch prop {
	case "name":
		setString(&z.Name, val)
	case "size":
		err = setUint(&z.Size, val)
	case "capacity":
		i := strings.Index(val, "%")
		if i < 0 {
			i = len(val)
		}
		err = setUint(&z.Capacity, val[:i])
	case "altroot":
		setString(&z.Altroot, val)
	case "health":
		setString(&z.Health, val)
	case "guid":
		setString(&z.GUID, val)
	case "version":
		setString(&z.Version, val)
	case "bootfs":
		setString(&z.Bootfs, val)
	case "delegation":
		setString(&z.Delegation, val)
	case "autoreplace":
		setString(&z.Autoreplace, val)
	case "cachefile":
		setString(&z.Cachefile, val)
	case "failmode":
		setString(&z.Failmode, val)
	case "listsnapshots":
		setString(&z.Listsnapshots, val)
	case "autoexpand":
		setString(&z.Autoexpand, val)
	case "dedupditto":
		setString(&z.Dedupditto, val)
	case "dedupratio":
		z.Dedupratio, err = strconv.ParseFloat(val[:len(val)-1], 64)
	case "free":
		err = setUint(&z.Free, val)
	case "allocated":
		err = setUint(&z.Allocated, val)
	case "readonly":
		setString(&z.Readonly, val)
	case "ashift":
		err = setUint(&z.Ashift, val)
	case "comment":
		setString(&z.Comment, val)
	case "expandsize":
		err = setUint(&z.Expandsize, val)
	case "freeing":
		err = setUint(&z.Freeing, val)
	case "fragmentation":
		i := strings.Index(val, "%")
		if i < 0 {
			i = len(val)
		}
		err = setUint(&z.Fragmentation, val[:i])
	case "leaked":
		err = setUint(&z.Leaked, val)
	case "feature@async_destroy":
		setString(&z.FeatureAsyncDestroy, val)
	case "feature@empty_bpobj":
		setString(&z.FeatureEmptyBpobj, val)
	case "feature@lz4_compress":
		setString(&z.FeatureLz4Compress, val)
	case "feature@spacemap_histogram":
		setString(&z.FeatureSpacemapHistogram, val)
	case "feature@enabled_txg":
		setString(&z.FeatureEnabledTxg, val)
	case "feature@hole_birth":
		setString(&z.FeatureHoleBirth, val)
	case "feature@extensible_dataset":
		setString(&z.FeatureExtensibleDataset, val)
	case "feature@embedded_data":
		setString(&z.FeatureEmbeddedData, val)
	case "feature@bookmarks":
		setString(&z.FeatureBookmarks, val)
	case "feature@filesystem_limits":
		setString(&z.FeatureFilesystemLimits, val)
	case "feature@large_blocks":
		setString(&z.FeatureLargeBlocks, val)
	}
	return err
}
