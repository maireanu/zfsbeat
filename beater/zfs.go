package beater

import (
	"strings"
)

// ZFS dataset types, which can indicate if a dataset is a filesystem,
// snapshot, or volume.
const (
	DatasetFilesystem = "filesystem"
	DatasetSnapshot   = "snapshot"
	DatasetVolume     = "volume"
)

// Dataset is a ZFS dataset.  A dataset could be a clone, filesystem, snapshot,
// or volume.  The Type struct member can be used to determine a dataset's type.
//
// The field definitions can be found in the ZFS manual:
// http://www.freebsd.org/cgi/man.cgi?zfs(8).
type Dataset struct {
	Name                 string
	Available            string
	Clones               string
	Compressratio        string
	Creation             string
	DeferDestroy         string
	Logicalreferenced    string
	Logicalused          string
	Mounted              string
	Origin               string
	Refcompressratio     string
	Referenced           string
	Type                 string
	Used                 string
	Usedbychildren       string
	Usedbydataset        string
	Usedbyrefreservation string
	Usedbysnapshots      string
	Userrefs             string
	Written              string
	Aclinherit           string
	Acltype              string
	Atime                string
	Canmount             string
	Casesensitivity      string
	Checksum             string
	Compression          string
	Context              string
	Copies               string
	Dedup                string
	Defcontext           string
	Devices              string
	Exec                 string
	FilesystemCount      string
	FilesystemLimit      string
	Fscontext            string
	Logbias              string
	Mlslabel             string
	Mountpoint           string
	Nbmand               string
	Normalization        string
	Overlay              string
	Primarycache         string
	Quota                string
	Readonly             string
	Recordsize           string
	RedundantMetadata    string
	Refquota             string
	Refreservation       string
	Relatime             string
	Reservation          string
	Rootcontext          string
	Secondarycache       string
	Setuid               string
	Sharenfs             string
	Sharesmb             string
	Snapdev              string
	Snapdir              string
	SnapshotCount        string
	SnapshotLimit        string
	Sync                 string
	Utf8only             string
	Version              string
	Volblocksize         string
	Volsize              string
	Vscan                string
	Xattr                string
	Zoned                string
}

var dsPropList = []string{"name", "available", "clones", "compressratio", "creation", "defer_destroy", "logicalreferenced", "logicalused", "mounted", "origin", "refcompressratio", "referenced", "type", "used", "usedbychildren", "usedbydataset", "usedbyrefreservation", "usedbysnapshots", "userrefs", "written", "aclinherit", "acltype", "atime", "canmount", "casesensitivity", "checksum", "compression", "context", "copies", "dedup", "defcontext", "devices", "exec", "filesystem_count", "filesystem_limit", "fscontext", "logbias", "mlslabel", "mountpoint", "nbmand", "normalization", "overlay", "primarycache", "quota", "readonly", "recordsize", "redundant_metadata", "refquota", "refreservation", "relatime", "reservation", "rootcontext", "secondarycache", "setuid", "sharenfs", "sharesmb", "snapdev", "snapdir", "snapshot_count", "snapshot_limit", "sync", "utf8only", "version", "volblocksize", "volsize", "vscan", "xattr", "zoned"}
var dsPropListOptions = strings.Join(dsPropList, ",")

func zfs(arg ...string) ([][]string, error) {
	c := command{Command: "zfs"}
	return c.Run(arg...)
}

// Datasets returns a slice of ZFS datasets, regardless of type.
// A filter argument may be passed to select a dataset with the matching name,
// or empty string ("") may be used to select all datasets.
func Datasets(filter string) ([]*Dataset, error) {
	return listByType("all", filter)
}

// Filesystems returns a slice of ZFS filesystems.
// A filter argument may be passed to select a filesystem with the matching name,
// or empty string ("") may be used to select all filesystems.
func Filesystems(filter string) ([]*Dataset, error) {
	return listByType(DatasetFilesystem, filter)
}

// Volumes returns a slice of ZFS volumes.
// A filter argument may be passed to select a volume with the matching name,
// or empty string ("") may be used to select all volumes.
func Volumes(filter string) ([]*Dataset, error) {
	return listByType(DatasetVolume, filter)
}

// Snapshots returns a slice of ZFS snapshots.
// A filter argument may be passed to select a snapshot with the matching name,
// or empty string ("") may be used to select all snapshots.
func Snapshots(filter string) ([]*Dataset, error) {
	return listByType(DatasetSnapshot, filter)
}

// Snapshots returns a slice of all ZFS snapshots of a given dataset.
func (d *Dataset) Snapshots() ([]*Dataset, error) {
	return Snapshots(d.Name)
}

// GetProperty returns the current value of a ZFS property from the

// GetDataset retrieves a single ZFS dataset by name.  This dataset could be
// any valid ZFS dataset type, such as a clone, filesystem, snapshot, or volume.
func GetDataset(name string) (*Dataset, error) {
	out, err := zfs("list", "-Hp", "-o", dsPropListOptions, name)
	if err != nil {
		return nil, err
	}

	ds := &Dataset{Name: name}
	for _, line := range out {
		if err := ds.parseLine(line); err != nil {
			return nil, err
		}
	}

	return ds, nil
}

func listByType(t, filter string) ([]*Dataset, error) {
	args := []string{"list", "-rpH", "-t", t, "-o", dsPropListOptions}

	if filter != "" {
		args = append(args, filter)
	}
	out, err := zfs(args...)
	if err != nil {
		return nil, err
	}

	var datasets []*Dataset

	name := ""
	var ds *Dataset
	for _, line := range out {
		if name != line[0] {
			name = line[0]
			ds = &Dataset{Name: name}
			datasets = append(datasets, ds)
		}
		if err := ds.parseLine(line); err != nil {
			return nil, err
		}
	}

	return datasets, nil
}

func (d *Dataset) parseLine(line []string) error {

	setString(&d.Name, line[0])
	setString(&d.Available, line[1])
	setString(&d.Clones, line[2])
	setString(&d.Compressratio, line[3])
	setString(&d.Creation, line[4])
	setString(&d.DeferDestroy, line[5])
	setString(&d.Logicalreferenced, line[6])
	setString(&d.Logicalused, line[7])
	setString(&d.Mounted, line[8])
	setString(&d.Origin, line[9])
	setString(&d.Refcompressratio, line[10])
	setString(&d.Referenced, line[11])
	setString(&d.Type, line[12])
	setString(&d.Used, line[13])
	setString(&d.Usedbychildren, line[14])
	setString(&d.Usedbydataset, line[15])
	setString(&d.Usedbyrefreservation, line[16])
	setString(&d.Usedbysnapshots, line[17])
	setString(&d.Userrefs, line[18])
	setString(&d.Written, line[19])
	setString(&d.Aclinherit, line[20])
	setString(&d.Acltype, line[21])
	setString(&d.Atime, line[22])
	setString(&d.Canmount, line[23])
	setString(&d.Casesensitivity, line[24])
	setString(&d.Checksum, line[25])
	setString(&d.Compression, line[26])
	setString(&d.Context, line[27])
	setString(&d.Copies, line[28])
	setString(&d.Dedup, line[29])
	setString(&d.Defcontext, line[30])
	setString(&d.Devices, line[31])
	setString(&d.Exec, line[32])
	setString(&d.FilesystemCount, line[33])
	setString(&d.FilesystemLimit, line[34])
	setString(&d.Fscontext, line[35])
	setString(&d.Logbias, line[36])
	setString(&d.Mlslabel, line[37])
	setString(&d.Mountpoint, line[38])
	setString(&d.Nbmand, line[39])
	setString(&d.Normalization, line[40])
	setString(&d.Overlay, line[41])
	setString(&d.Primarycache, line[42])
	setString(&d.Quota, line[43])
	setString(&d.Readonly, line[44])
	setString(&d.Recordsize, line[45])
	setString(&d.RedundantMetadata, line[46])
	setString(&d.Refquota, line[47])
	setString(&d.Refreservation, line[48])
	setString(&d.Relatime, line[49])
	setString(&d.Reservation, line[50])
	setString(&d.Rootcontext, line[51])
	setString(&d.Secondarycache, line[52])
	setString(&d.Setuid, line[53])
	setString(&d.Sharenfs, line[54])
	setString(&d.Sharesmb, line[55])
	setString(&d.Snapdev, line[56])
	setString(&d.Snapdir, line[57])
	setString(&d.SnapshotCount, line[58])
	setString(&d.SnapshotLimit, line[59])
	setString(&d.Sync, line[60])
	setString(&d.Utf8only, line[61])
	setString(&d.Version, line[62])
	setString(&d.Volblocksize, line[63])
	setString(&d.Volsize, line[64])
	setString(&d.Vscan, line[65])
	setString(&d.Xattr, line[66])

	return nil
}
