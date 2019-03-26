package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/maireanu/zfsbeat/config"
)

// Zfsbeat configuration.
type Zfsbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of zfsbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Zfsbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts zfsbeat.
func (bt *Zfsbeat) Run(b *beat.Beat) error {
	logp.Info("zfsbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		var events = []beat.Event{}

		pools, err := ListZpools()
		if err != nil {
			panic(err)
		}

		snapshots, err := Snapshots("")
		if err != nil {
			panic(err)
		}

		filesystems, err := Filesystems("")
		if err != nil {
			panic(err)
		}

		if bt.config.SourceFilesystem == true {
			for _, filesystem := range filesystems {
				event := beat.Event{
					Timestamp: time.Now(),
					Fields: common.MapStr{
						"source":                "filesystem",
						"name":                  filesystem.Name,
						"available":             filesystem.Available,
						"clones":                filesystem.Clones,
						"compressratio":         filesystem.Compressratio,
						"creation":              filesystem.Creation,
						"defer.destroy":         filesystem.DeferDestroy,
						"logical.referenced":    filesystem.Logicalreferenced,
						"logical.used":          filesystem.Logicalused,
						"mounted":               filesystem.Mounted,
						"origin":                filesystem.Origin,
						"ref.compressratio":     filesystem.Refcompressratio,
						"referenced":            filesystem.Referenced,
						"type":                  filesystem.Type,
						"used":                  filesystem.Used,
						"usedby.children":       filesystem.Usedbychildren,
						"usedby.dataset":        filesystem.Usedbydataset,
						"usedby.refreservation": filesystem.Usedbyrefreservation,
						"usedby.snapshots":      filesystem.Usedbysnapshots,
						"userrefs":              filesystem.Userrefs,
						"written":               filesystem.Written,
						"acl.inherit":           filesystem.Aclinherit,
						"acl.type":              filesystem.Acltype,
						"atime":                 filesystem.Atime,
						"canmount":              filesystem.Canmount,
						"casesensitivity":       filesystem.Casesensitivity,
						"checksum":              filesystem.Checksum,
						"compression":           filesystem.Compression,
						"context":               filesystem.Context,
						"copies":                filesystem.Copies,
						"dedup":                 filesystem.Dedup,
						"defcontext":            filesystem.Defcontext,
						"devices":               filesystem.Devices,
						"exec":                  filesystem.Exec,
						"filesystem.count":      filesystem.FilesystemCount,
						"filesystem.limit":      filesystem.FilesystemLimit,
						"fscontext":             filesystem.Fscontext,
						"logbias":               filesystem.Logbias,
						"mlslabel":              filesystem.Mlslabel,
						"mountpoint":            filesystem.Mountpoint,
						"nbmand":                filesystem.Nbmand,
						"normalization":         filesystem.Normalization,
						"overlay":               filesystem.Overlay,
						"primarycache":          filesystem.Primarycache,
						"quota":                 filesystem.Quota,
						"readonly":              filesystem.Readonly,
						"recordsize":            filesystem.Recordsize,
						"redundant.metadata":    filesystem.RedundantMetadata,
						"ref.quota":             filesystem.Refquota,
						"ref.reservation":       filesystem.Refreservation,
						"relatime":              filesystem.Relatime,
						"reservation":           filesystem.Reservation,
						"rootcontext":           filesystem.Rootcontext,
						"secondarycache":        filesystem.Secondarycache,
						"setuid":                filesystem.Setuid,
						"share.nfs":             filesystem.Sharenfs,
						"share.smb":             filesystem.Sharesmb,
						"snap.dev":              filesystem.Snapdev,
						"snap.dir":              filesystem.Snapdir,
						"snapshot.count":        filesystem.SnapshotCount,
						"snapshot.limit":        filesystem.SnapshotLimit,
						"sync":                  filesystem.Sync,
						"utf8only":              filesystem.Utf8only,
						"version":               filesystem.Version,
						"vol.blocksize":         filesystem.Volblocksize,
						"vol.size":              filesystem.Volsize,
						"vscan":                 filesystem.Vscan,
						"xattr":                 filesystem.Xattr,
					},
				}
				events = append(events, event)
			}
		}

		if bt.config.SourceSnapshot == true {
			for _, snapshot := range snapshots {
				event := beat.Event{
					Timestamp: time.Now(),
					Fields: common.MapStr{
						"source":                "snapshot",
						"name":                  snapshot.Name,
						"available":             snapshot.Available,
						"clones":                snapshot.Clones,
						"compressratio":         snapshot.Compressratio,
						"creation":              snapshot.Creation,
						"defer.destroy":         snapshot.DeferDestroy,
						"logical.referenced":    snapshot.Logicalreferenced,
						"logical.used":          snapshot.Logicalused,
						"mounted":               snapshot.Mounted,
						"origin":                snapshot.Origin,
						"ref.compressratio":     snapshot.Refcompressratio,
						"referenced":            snapshot.Referenced,
						"type":                  snapshot.Type,
						"used":                  snapshot.Used,
						"usedby.children":       snapshot.Usedbychildren,
						"usedby.dataset":        snapshot.Usedbydataset,
						"usedby.refreservation": snapshot.Usedbyrefreservation,
						"usedby.snapshots":      snapshot.Usedbysnapshots,
						"userrefs":              snapshot.Userrefs,
						"written":               snapshot.Written,
						"acl.inherit":           snapshot.Aclinherit,
						"acl.type":              snapshot.Acltype,
						"atime":                 snapshot.Atime,
						"canmount":              snapshot.Canmount,
						"casesensitivity":       snapshot.Casesensitivity,
						"checksum":              snapshot.Checksum,
						"compression":           snapshot.Compression,
						"context":               snapshot.Context,
						"copies":                snapshot.Copies,
						"dedup":                 snapshot.Dedup,
						"defcontext":            snapshot.Defcontext,
						"devices":               snapshot.Devices,
						"exec":                  snapshot.Exec,
						"filesystem.count":      snapshot.FilesystemCount,
						"filesystem.limit":      snapshot.FilesystemLimit,
						"fscontext":             snapshot.Fscontext,
						"logbias":               snapshot.Logbias,
						"mlslabel":              snapshot.Mlslabel,
						"mountpoint":            snapshot.Mountpoint,
						"nbmand":                snapshot.Nbmand,
						"normalization":         snapshot.Normalization,
						"overlay":               snapshot.Overlay,
						"primarycache":          snapshot.Primarycache,
						"quota":                 snapshot.Quota,
						"readonly":              snapshot.Readonly,
						"recordsize":            snapshot.Recordsize,
						"redundant.metadata":    snapshot.RedundantMetadata,
						"ref.quota":             snapshot.Refquota,
						"ref.reservation":       snapshot.Refreservation,
						"relatime":              snapshot.Relatime,
						"reservation":           snapshot.Reservation,
						"rootcontext":           snapshot.Rootcontext,
						"secondarycache":        snapshot.Secondarycache,
						"setuid":                snapshot.Setuid,
						"share.nfs":             snapshot.Sharenfs,
						"share.smb":             snapshot.Sharesmb,
						"snap.dev":              snapshot.Snapdev,
						"snap.dir":              snapshot.Snapdir,
						"snapshot.count":        snapshot.SnapshotCount,
						"snapshot.limit":        snapshot.SnapshotLimit,
						"sync":                  snapshot.Sync,
						"utf8only":              snapshot.Utf8only,
						"version":               snapshot.Version,
						"vol.blocksize":         snapshot.Volblocksize,
						"vol.size":              snapshot.Volsize,
						"vscan":                 snapshot.Vscan,
						"xattr":                 snapshot.Xattr,
					},
				}
				events = append(events, event)
			}
		}

		if bt.config.SourceZpool == true {
			for _, pool := range pools {
				event := beat.Event{
					Timestamp: time.Now(),
					Fields: common.MapStr{
						"source":                    "zpool",
						"name":                      pool.Name,
						"size":                      pool.Size,
						"capacity":                  pool.Capacity,
						"altroot":                   pool.Altroot,
						"health":                    pool.Health,
						"guid":                      pool.GUID,
						"version":                   pool.Version,
						"bootfs":                    pool.Bootfs,
						"delegation":                pool.Delegation,
						"autoreplace":               pool.Autoreplace,
						"cachefile":                 pool.Cachefile,
						"failmode":                  pool.Failmode,
						"listsnapshots":             pool.Listsnapshots,
						"autoexpand":                pool.Autoexpand,
						"dedup_ditto":               pool.Dedupditto,
						"dedup_ratio":               pool.Dedupratio,
						"free":                      pool.Free,
						"allocated":                 pool.Allocated,
						"readonly":                  pool.Readonly,
						"ashift":                    pool.Ashift,
						"comment":                   pool.Comment,
						"expandsize":                pool.Expandsize,
						"freeing":                   pool.Freeing,
						"fragmentation":             pool.Fragmentation,
						"leaked":                    pool.Leaked,
						"feature.asyncdestroy":      pool.FeatureAsyncDestroy,
						"feature.emptybpobj":        pool.FeatureEmptyBpobj,
						"feature.lz4compress":       pool.FeatureLz4Compress,
						"feature.spacemaphistogram": pool.FeatureSpacemapHistogram,
						"feature.enabledtxg":        pool.FeatureEnabledTxg,
						"feature.holebirth":         pool.FeatureHoleBirth,
						"feature.extensibledataset": pool.FeatureExtensibleDataset,
						"feature.embeddeddata":      pool.FeatureEmbeddedData,
						"feature.bookmarks":         pool.FeatureBookmarks,
						"feature.filesystemlimits":  pool.FeatureFilesystemLimits,
						"feature.largeblocks":       pool.FeatureLargeBlocks,
					},
				}
				events = append(events, event)
			}
		}
		bt.client.PublishAll(events)
		logp.Info("Event sent")
	}
}

// Stop stops zfsbeat.
func (bt *Zfsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
