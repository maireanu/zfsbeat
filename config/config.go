// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

//Config for period
type Config struct {
	Period           time.Duration `config:"period"`
	SourceZpool      bool          `config:"source_zpool"`
	SourceFilesystem bool          `config:"source_filesystem"`
	SourceSnapshot   bool          `config:"source_snapshot"`
}

//DefaultConfig configuration for period
var DefaultConfig = Config{
	Period:           1 * time.Second,
	SourceZpool:      true,
	SourceFilesystem: true,
	SourceSnapshot:   true,
}
