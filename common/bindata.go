// Code generated by go-bindata. DO NOT EDIT.
// sources:
// common/data/migrations/1_accounts_create.sql
// common/data/migrations/1_active_probes_create.sql
// common/data/migrations/1_jobs_create.sql
// common/data/migrations/1_probe_updates_create.sql
// common/data/migrations/1_tasks_create.sql
// common/data/migrations/2_add_jobs_state.sql
// common/data/migrations/2_add_language_column.sql
// common/data/migrations/3_add_job_type_tables.sql
// common/data/migrations/4_rendezvous_tables.sql
// common/data/migrations/5_experiment_tables.sql

package common

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _commonDataMigrations1_accounts_createSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xc1\x8e\xda\x30\x10\xbd\xfb\x2b\xde\x01\x29\xa0\xee\x1e\x7a\x8e" +
	"\x7a\x30\xc9\x50\xac\x26\x0e\x75\x9c\xee\xd2\x4b\x64\x25\x16\x6b\x09\x4c\x84\x4d\x77\xf7\xef\x2b\x42\xa9\x36\x52" +
	"\xcb\x6d\xfc\xe6\xcd\x9b\x79\x4f\x7e\x7c\xc4\xa7\x83\xdb\x9d\x4c\xb4\xc8\x8f\xaf\x9e\x7d\x04\xea\x68\xa2\x3d\x58" +
	"\x1f\x97\x76\xe7\x3c\xcb\x55\xb5\x81\xe6\xcb\x82\x20\x56\xb0\x6f\x2e\xc4\x00\xd3\x75\xc7\xb3\x8f\x21\xfd\xf7\x24" +
	"\xf9\x9e\x4d\x3a\xcd\x70\x77\x45\x85\xd9\x8c\x01\xc0\x92\xbe\x0a\x39\x56\x62\x05\x59\x69\xd0\xb3\xa8\x75\x8d\x79" +
	"\x4d\x05\x65\x1a\x9f\xb1\x52\x55\x89\x61\xd7\xc6\xf7\xc1\xe2\x69\x4d\x8a\x10\xdf\x07\x6f\x0e\x16\x5f\x90\xfc\xb9" +
	"\xab\x3d\x1d\xf7\x36\x59\x40\xaf\xe9\xaa\x96\x29\xe2\x9a\xa0\xb7\x1b\x02\xcf\xb2\xaa\x91\xba\x55\x55\x41\xe0\x35" +
	"\x48\x36\x25\xe6\x49\x6f\x7f\xb9\xce\x26\x0f\x48\x4c\x7f\x70\xfe\x52\x9c\x83\x3d\x25\x8b\x74\x54\x20\x99\x43\xac" +
	"\x52\x46\x32\x9f\xcd\x52\xc6\x6e\x8a\xb7\x60\x3e\x1c\x7b\x0b\x87\xcd\xc7\x49\xd7\xa3\x26\x25\x78\x81\x8d\x12\x25" +
	"\x57\x5b\x7c\xa3\xed\xc8\x97\x4d\x51\x3c\x8c\x9c\xcb\xa6\xd1\xc3\x0f\xae\xb2\x35\x57\x57\x74\x30\x21\xbc\x1e\x4f" +
	"\x7d\xfb\x62\xc2\xcb\xb4\x15\xcc\x3e\x4e\x91\xbd\x09\xb1\x35\x5d\x67\x43\x80\x16\x25\xd5\x9a\x97\x1b\x3c\x09\xbd" +
	"\x1e\x9f\xf8\x59\x49\xba\x32\x2f\xe1\x4c\x52\x60\x8b\xf4\xe6\xa7\x91\xe2\x7b\x43\x10\x32\xa7\xe7\xff\xd8\x6a\x5d" +
	"\xdf\x9e\x9d\xef\xed\x1b\x2a\xf9\x17\xc5\xdc\xf5\x8b\x3b\xdf\xe1\x77\x00\x00\x00\xff\xff\x49\x92\xd5\x50\x73\x02" +
	"\x00\x00")

func commonDataMigrations1_accounts_createSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations1_accounts_createSql,
		"common/data/migrations/1_accounts_create.sql",
	)
}

func commonDataMigrations1_accounts_createSql() (*asset, error) {
	bytes, err := commonDataMigrations1_accounts_createSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/1_accounts_create.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations1_active_probes_createSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xcf\x8e\xd3\x30\x10\x87\xef\x7e\x8a\x39\xb6\x82\x95\x16\x38\xf6" +
	"\x94\xdd\x18\xad\x45\x9b\x94\xfc\x81\x2d\x08\x59\x6e\x3c\x2d\x56\x13\x3b\xb2\xa7\x0d\x7d\x7b\xa4\x84\x94\xa6\xea" +
	"\x1e\x67\xbe\xf9\xc6\xb2\x7f\x7e\x78\x80\x77\x8d\xd9\x7b\x45\x08\xb1\xeb\x2c\x8b\xb3\x74\x0d\x45\xf4\xb4\xe4\x20" +
	"\x3e\x03\x7f\x15\x79\x91\x83\xaa\xc8\x9c\x50\xb6\xde\x6d\x31\x2c\x18\xbb\xb6\xca\x76\x52\xe6\xa4\x08\x1b\xb4\xf4" +
	"\x84\x7b\x63\xd9\x73\xc6\xa3\x82\xff\x5f\x98\xa4\xc5\xdd\xa5\x6c\xc6\x00\x00\x8c\x86\xb2\x14\x31\xac\x33\xb1\x8a" +
	"\xb2\x0d\x7c\xe1\x9b\x5e\x49\xca\xe5\xf2\x7d\x3f\x51\x79\x54\x64\x9c\x95\x64\x1a\x84\x42\xac\x78\x5e\x44\xab\x35" +
	"\x7c\x17\xc5\x4b\x5f\xc2\x8f\x34\xe1\xc3\x6c\xbf\x5a\x56\x15\x7c\x8b\xb2\xe7\x97\x28\x9b\x7d\x9c\x5f\x03\x15\xec" +
	"\x85\x7c\x7a\x1c\x51\xad\x68\xe7\x7c\x33\x92\xa1\x1b\xdc\x8e\x3a\xe5\x51\x5a\xd5\xe0\x45\xfa\xf0\x38\x5a\x17\x7e" +
	"\x42\x1f\x8c\xb3\x37\xf6\xb1\x6d\x9d\x27\xd4\x92\x30\x50\x18\xe1\xcf\x5f\x03\xb6\x48\x9d\xf3\x07\x49\xe7\x16\xa7" +
	"\xa2\x3a\x29\x53\xab\x6d\x8d\x72\xab\xac\xee\x8c\xa6\xdf\xd3\x01\x72\x07\xbc\x39\x6c\xb8\xdb\x4e\x35\xa6\x3e\xdf" +
	"\x23\x46\x4f\xbb\xb5\x0a\x24\x8f\xad\x56\x84\xfa\xcd\xf7\x64\xf3\xc5\x18\x65\x99\x88\xaf\x25\x07\x91\xc4\xfc\xf5" +
	"\x26\x51\x8f\x7b\x13\x08\x3d\xea\x7f\xa9\x4a\xa3\xe5\xd1\x58\x8d\x7f\x20\x4d\xa6\x81\xc3\xcc\xe8\xf9\xe2\xfe\xd7" +
	"\xe1\x56\xb3\xbf\x01\x00\x00\xff\xff\x64\x48\x47\x43\x99\x02\x00\x00")

func commonDataMigrations1_active_probes_createSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations1_active_probes_createSql,
		"common/data/migrations/1_active_probes_create.sql",
	)
}

func commonDataMigrations1_active_probes_createSql() (*asset, error) {
	bytes, err := commonDataMigrations1_active_probes_createSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/1_active_probes_create.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations1_jobs_createSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x8f\x9b\x30\x10\x85\xef\xfe\x15\xef\x10\x89\x44\xdd\x95\xba" +
	"\xbd\xa2\x1e\x60\x71\x1a\x6f\xc1\x44\x60\xba\x4d\xab\xca\xf2\x86\x29\x65\x1b\x0c\x02\xa7\x6d\xfe\x7d\x05\x49\x94" +
	"\x44\xda\x9b\x9f\xe7\xe3\x31\x6f\xc6\xf7\xf7\x78\xd7\xd4\x55\x6f\x1c\x21\x6a\xff\x5a\x16\x65\xe9\x1a\x2a\x08\x63" +
	"\x0e\xb1\x04\xff\x2a\x72\x95\xe3\xb5\x7d\x19\x7c\xc6\xae\xe1\xa2\xbb\x91\xb9\x33\x8e\x1a\xb2\x2e\xa4\xaa\xb6\x2c" +
	"\x4a\x31\x9b\x31\x00\x08\xf9\x27\x21\xa7\x93\x58\x42\xa6\xea\x6c\x39\xcf\x79\xcc\x1f\x15\x1e\xb0\xcc\xd2\x04\x5d" +
	"\xa5\xdd\xa1\x23\x3c\xaf\x78\xc6\xe1\x0e\x9d\x35\x0d\xe1\x23\xbc\xd7\xf6\x45\x0f\xa3\xb9\xb7\x80\x5a\xf1\xa3\xd5" +
	"\x63\xc6\x03\xc5\xa1\x36\x6b\x8e\xa7\x34\xd4\xb9\x1a\x65\x90\x83\xcb\x22\xc1\xdc\x33\x5b\x57\xff\x21\xef\x0e\x5e" +
	"\x49\x3b\x72\x54\x4e\xc7\xd6\x92\xb7\xf0\x27\x03\x2e\x23\x88\xa5\xcf\xb8\x8c\x66\x33\x9f\xb1\xb3\xe1\x39\xf7\x55" +
	"\xa3\x63\x76\x36\x9f\xbe\xaa\x4b\x14\x85\x88\xb0\xce\x44\x12\x64\x1b\x7c\xe6\x9b\x89\x94\x45\x1c\xdf\x4d\xc4\xb6" +
	"\x6d\xc6\x21\xe0\x4b\x90\x3d\xae\x82\xec\x74\xd9\x93\x71\x75\x6b\xb5\xab\x1b\x82\x12\x09\xcf\x55\x90\xac\xf1\x2c" +
	"\xd4\x6a\x92\xf8\x96\x4a\x7e\x64\x87\xed\x2f\x2a\xf7\x3b\xba\x75\x28\x69\x67\x0e\x10\x52\x1d\xa5\x33\x7d\x45\x4e" +
	"\x6f\xdb\xbd\x75\x7d\x4d\xc3\x19\x9e\x7f\x58\xe0\xfb\x8f\x1b\xa6\xdb\x19\xf7\xb3\xed\x9b\x0b\xf3\xf0\xfe\x1a\x1a" +
	"\x7e\x6b\x47\x83\xd3\xd3\xb8\x6f\xfe\x39\xd5\x4c\x5f\xed\xc7\x40\x03\x9e\xf2\x54\x86\xa7\x4a\xdd\xd0\xa0\xfb\xbd" +
	"\xbd\x74\x64\xe9\x9f\x1b\x6f\xb4\x71\xc7\x44\x6f\x65\xab\x07\x3d\xee\x00\x61\x9a\xc6\x3c\x90\xa7\xc0\xe3\x6e\x2f" +
	"\x4b\x64\x0b\xff\xed\x77\xc5\x6d\xc9\xfe\x07\x00\x00\xff\xff\x0d\xbe\x24\xa3\xad\x02\x00\x00")

func commonDataMigrations1_jobs_createSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations1_jobs_createSql,
		"common/data/migrations/1_jobs_create.sql",
	)
}

func commonDataMigrations1_jobs_createSql() (*asset, error) {
	bytes, err := commonDataMigrations1_jobs_createSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/1_jobs_create.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations1_probe_updates_createSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xcf\x6e\xc2\x30\x0c\x87\xef\x79\x0a\x1f\x41\x1b\x12\xdb\x8e\x9c" +
	"\x0a\x64\xa2\x1a\xff\xd4\xa6\xdb\xd8\x34\x45\x81\x18\x16\xd1\x26\x55\x6a\xa8\x78\xfb\x49\x74\x65\x94\xf5\x68\x7f" +
	"\xbf\xcf\x56\xe2\x5e\x0f\xee\x32\xb3\xf3\x8a\x10\xc6\xae\xb4\x6c\x1c\x2d\x96\x20\x82\xe1\x94\x43\xf8\x0c\xfc\x3d" +
	"\x8c\x45\x0c\xb9\x77\x6b\x94\x87\x5c\x2b\xc2\x62\xc0\xd8\xb5\x95\xe4\x8d\x32\x26\x45\x98\xa1\xa5\x21\xee\x8c\x65" +
	"\xa3\x88\x07\x82\xff\x0d\x9c\x2f\x44\xeb\x50\xd6\x61\x00\x00\x46\x43\x92\x84\x63\x58\x46\xe1\x2c\x88\x56\xf0\xc2" +
	"\x57\x67\x65\x9e\x4c\xa7\xf7\xe7\x44\x95\x97\x64\x32\x04\x11\xce\x78\x2c\x82\xd9\x12\xde\x42\x31\x39\x97\xf0\xb1" +
	"\x98\xf3\x2a\x59\xcd\xdf\x6c\xe0\x35\x88\x46\x93\x20\xea\x3c\x76\xaf\x81\x2a\xec\x85\x3c\xf5\x6b\x94\x2a\xda\x3a" +
	"\x9f\xd5\xa4\xea\x16\x6e\x4b\xa5\xf2\x28\xad\xca\xf0\x22\x3d\xf4\x6b\xeb\xc2\x8f\xe8\x0b\xe3\xec\x8d\x7d\xc8\x73" +
	"\xe7\x09\xb5\x24\x2c\xa8\xa8\xe1\xe7\x57\x85\x2d\x52\xe9\xfc\x5e\xd2\x29\xc7\xa6\xa8\x8e\xca\xa4\x6a\x9d\xa2\x5c" +
	"\x2b\xab\x4b\xa3\xe9\xbb\x19\x20\xb7\xc7\x9b\x65\xd5\xdb\xb6\x2a\x33\xe9\xa9\x8d\x18\xdd\xec\xd6\xdf\xf9\x6f\xf9" +
	"\x26\x35\x68\x49\xfe\x1e\x84\x75\x07\xed\x57\xe6\x56\xb3\x9f\x00\x00\x00\xff\xff\xaf\x10\xdb\xee\x44\x02\x00\x00" +
	"")

func commonDataMigrations1_probe_updates_createSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations1_probe_updates_createSql,
		"common/data/migrations/1_probe_updates_create.sql",
	)
}

func commonDataMigrations1_probe_updates_createSql() (*asset, error) {
	bytes, err := commonDataMigrations1_probe_updates_createSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/1_probe_updates_create.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations1_tasks_createSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x8f\xd3\x30\x10\x85\xef\xfe\x15\xef\x50\x29\xad\xd8\x3d\x70" +
	"\x8e\x38\xa4\x8d\x4b\x0d\xad\x53\xc5\x0e\x0b\x5c\x22\x6f\x32\x44\x59\x68\x12\xc5\xb3\x82\xfe\x7b\x14\xa7\x8b\xba" +
	"\x12\x2b\xf5\xf6\xe6\xf9\xcd\x8c\xe6\x93\xef\xef\xf1\xee\xd4\x36\xa3\x63\x42\xda\xff\xee\x44\x9a\x67\x47\xd8\x64" +
	"\xbd\x97\x50\x5b\xd0\x9f\xd6\xb3\x07\x3b\xff\xd3\xc7\x42\x5c\xa7\x8b\xe1\x55\x69\xd8\x31\x9d\xa8\xe3\x35\x35\x6d" +
	"\x27\xd2\x0c\x8b\x85\x00\x80\xb5\xfc\xa8\x74\x50\x6a\x0b\x9d\x59\xc8\xaf\xca\x58\x83\xa5\x91\x7b\xb9\xb1\x78\x8f" +
	"\x6d\x9e\x1d\x30\x34\x25\x9f\x07\xc2\xc3\x4e\xe6\x12\x7c\x1e\x3a\x77\x22\x7c\x40\x34\xed\x2e\xfd\x34\x3d\x5a\xc1" +
	"\xee\xe4\x3c\x6b\x93\xcb\xc4\x4a\xd8\x6f\x47\x09\x9b\x98\xcf\xa5\xb1\x53\x9d\x18\x48\x5d\x1c\xb0\x8c\x46\x72\xf5" +
	"\x39\xba\x43\xd4\xf5\xdc\xfe\x68\xa9\x9e\xb4\xab\x2a\x1a\x78\xd6\x23\x3d\x51\x75\xd1\x75\xdf\x51\xb4\x8a\xc3\x64" +
	"\xa9\x53\xa8\x6d\x2c\xa4\x4e\x17\x8b\x58\x88\x97\x4d\x2f\x48\xae\x4e\x08\x58\xc4\x32\xb4\xb5\x35\x8a\x42\xa5\xe1" +
	"\x59\x17\xfb\xfd\x5d\x70\x87\xb1\x7f\xa4\xf2\xf2\x36\x5b\x4f\xfd\xe3\x6b\x83\xc9\x73\x19\xae\xfd\x92\xe4\x9b\x5d" +
	"\x92\xcf\xb6\x1b\x9b\xe7\x89\xa7\xc7\x27\x93\xe9\xf5\x6c\x06\x0e\x57\x07\xff\xdb\xd2\x8c\xe4\x3d\x94\xb6\xb3\x53" +
	"\x8d\xe4\xb8\xed\xbb\x92\xdb\x13\xc1\xaa\x83\x34\x36\x39\x1c\xf1\xa0\xec\x2e\x94\xf8\x9e\xe9\x4b\xf7\x0c\xa8\xba" +
	"\x39\x3f\x43\xbc\x25\x39\x61\xbd\x25\xf7\xcb\x79\x2e\x9f\x87\xda\x31\xd5\x6f\x46\xc5\x2a\xfe\xff\x87\x93\x5d\x2d" +
	"\xfe\x06\x00\x00\xff\xff\x3a\x61\xc2\x6a\xc7\x02\x00\x00")

func commonDataMigrations1_tasks_createSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations1_tasks_createSql,
		"common/data/migrations/1_tasks_create.sql",
	)
}

func commonDataMigrations1_tasks_createSql() (*asset, error) {
	bytes, err := commonDataMigrations1_tasks_createSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/1_tasks_create.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations2_add_jobs_stateSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4f\x8f\x9b\x30\x10\xc5\xef\xfe\x14\xef\x80\xe4\x5d\xb5\x5b\xa9" +
	"\x67\xd4\x03\x7f\x86\xc6\x15\x31\x11\x38\xda\xed\x09\xd8\x60\x45\xac\xc0\xa0\xe0\xb4\xcd\xb7\xaf\x70\x9b\x7f\x4d" +
	"\xba\x9c\x06\xcf\xf8\xcd\xef\x3d\x3f\x3d\xe1\x43\xdf\x6e\x77\xb5\xd5\x88\x87\x9f\x86\x5d\x1e\x14\xb6\xb6\xba\xd7" +
	"\xc6\x86\x7a\xdb\x1a\x16\xe7\xd9\x0a\xea\xfb\x8a\xf0\x2d\x0b\xcb\x42\x05\x8a\x20\x12\xd0\x8b\x28\x54\xe1\xb3\x20" +
	"\x55\x94\x43\x05\x61\x4a\x78\x1b\x5e\x27\xb8\xf9\x28\x4b\xd7\x4b\x79\x9e\xc3\x34\x8b\xfa\xf7\xf7\x90\x69\xd8\x55" +
	"\x67\x3d\xbe\x0b\x94\xc1\xf3\x18\x00\x84\xf4\x55\x48\x57\x89\x04\x32\x53\xc7\x65\x0f\x05\xa5\x14\x29\x7c\x46\x92" +
	"\x67\x4b\x8c\xdb\xd2\x1e\x46\x8d\xe7\x05\xe5\x04\x7b\x18\x4d\xdd\x6b\x7c\x01\x7f\x1b\x5e\x4b\x07\xc6\x1f\xa1\x16" +
	"\xf4\x47\x2a\xca\x69\xb6\xf8\x8f\xe3\xa0\x00\xc9\xf5\x12\x0f\xbc\xde\xd8\xf6\x87\xe6\x1f\xc1\x1b\xdd\x69\xab\x1b" +
	"\x57\x0e\x46\xf3\x47\xdf\x09\x90\x8c\x21\x12\x9f\x91\x8c\x3d\xcf\x67\x77\x79\x6f\xff\xe6\xef\x26\xcb\x20\x8e\x8f" +
	"\x51\x3a\xce\x33\x90\x7f\xba\x48\x2f\x11\xad\x94\xc8\xae\xa5\x9e\x17\x24\xd1\xec\xc7\xae\xdd\xd4\x56\x97\x9b\xa1" +
	"\xdb\xf7\xc6\x99\x44\x1e\x88\x82\xe6\xb8\x44\x44\xe0\x7f\x3b\x95\xd3\xaf\x50\x77\x3b\x5d\x37\x07\xe8\x5f\xed\x64" +
	"\x27\xb4\x06\xd5\x4c\x52\x7d\xe2\x17\x1b\x65\x7c\x72\xea\x33\xcf\xfb\xff\xab\xfe\x0e\x00\x00\xff\xff\x64\x65\x02" +
	"\xda\x68\x02\x00\x00")

func commonDataMigrations2_add_jobs_stateSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations2_add_jobs_stateSql,
		"common/data/migrations/2_add_jobs_state.sql",
	)
}

func commonDataMigrations2_add_jobs_stateSql() (*asset, error) {
	bytes, err := commonDataMigrations2_add_jobs_stateSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/2_add_jobs_state.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations2_add_language_columnSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x92\xcd\x4e\x84\x30\x14\x85\xf7\x3c\xc5\x59\x90\x8c\xc6\x8c\x3b\x57" +
	"\xb3\xea\xd0\xea\x34\x41\x98\x40\x47\x67\x07\x15\x1a\xd2\x04\x0a\x81\xe2\xcf\xdb\x9b\x41\x25\xd3\xe8\xa8\x71\x65" +
	"\x57\xf7\x27\xe7\xde\xef\xa4\x77\xb9\xc4\x45\xa3\xab\x5e\x5a\x05\xda\x3e\x19\xef\xb8\x90\x5a\x69\x55\xa3\x8c\x5d" +
	"\xab\x4a\x1b\x8f\x84\x82\x25\x10\x64\x1d\x32\xc8\xc2\xea\x47\x95\x75\x7d\xfb\xa0\x06\xd0\x24\xde\x22\x88\xc3\xdd" +
	"\x6d\x04\x7e\x0d\xb6\xe7\xa9\x48\x51\x4b\x53\x65\x45\x5b\xaa\x95\x23\x9d\x34\xd9\xd8\x95\xd2\xfe\x46\xfa\x25\x10" +
	"\x33\xa5\xe7\x74\x76\xdd\x77\xe4\x1e\x8d\xe1\xfb\x1e\x00\xac\xd9\x0d\x8f\xa6\xe8\x73\x76\x78\xa7\x4d\x12\x4a\x3f" +
	"\x40\x67\x3c\xdc\x91\x24\xd8\x90\xe4\xec\xea\x7c\x35\x8f\x61\xfb\x80\x6d\x05\x8f\xdd\xc1\xf7\x1b\x16\xa1\x1c\xbb" +
	"\x5a\x17\xd2\xaa\xac\x68\xeb\xb1\x31\x10\x87\x6a\x42\x78\xca\x10\xc5\x82\x07\x0c\x8b\xf7\x4e\x3e\x2f\xc9\x21\xeb" +
	"\x5e\xc9\xf2\x05\xea\x59\x0f\x76\x80\x36\xc8\x1d\xb6\xfc\x72\x71\xb4\x3e\xa2\x6f\xc9\x14\xf8\xfe\xea\xaf\xf6\xdd" +
	"\x8f\xfa\x5f\xf6\x1d\xb6\x9f\xec\x9f\x3c\xa1\xd7\x00\x00\x00\xff\xff\xa2\x48\x52\x38\xfe\x02\x00\x00")

func commonDataMigrations2_add_language_columnSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations2_add_language_columnSql,
		"common/data/migrations/2_add_language_column.sql",
	)
}

func commonDataMigrations2_add_language_columnSql() (*asset, error) {
	bytes, err := commonDataMigrations2_add_language_columnSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/2_add_language_column.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations3_add_job_type_tablesSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\x4f\x73\xda\x3c\x10\xc6\xef\xfe\x14\xba\x91\xcc\xfb\xa6\x1f\x20" +
	"\x9c\x0c\xa8\x6d\x5a\x0a\xd4\x98\xce\x70\xd2\xac\xed\xc5\x28\xc8\x92\x23\xad\x21\xf9\xf6\x1d\xdb\xc8\x31\x34\x10" +
	"\x8e\x7e\xf6\xd1\xea\xa7\xfd\xe3\x87\x07\xf6\x5f\x21\x73\x0b\x84\x6c\x62\x0e\x3a\xe8\x0b\x4b\x02\xc2\x02\x35\x8d" +
	"\x30\x97\x3a\x98\x44\xf3\x05\x8b\xc3\xd1\x94\xb3\x67\x93\x08\x02\xb7\x73\xc3\x73\x15\x14\x5a\x72\xc3\x20\x08\xa7" +
	"\x31\x8f\x8e\x81\xb2\x4a\x94\x4c\xbf\x3c\x9b\xc4\xb1\xc6\x5f\x9f\x15\xda\x0c\xaf\xbb\x9a\x5c\x8d\xed\xa2\x2f\x9c" +
	"\x4c\xda\x64\x84\x8e\x84\x86\x02\xd9\x9f\x30\x1a\x7f\x0f\x23\x36\x5b\x4d\xa7\x97\x2f\xe8\x0e\x82\xcd\xab\xfa\x8d" +
	"\x8e\xfd\x58\xce\x67\xa3\x61\x70\x7c\xd1\x7a\xc1\xd9\x7c\x1e\xf3\x65\x7c\x7c\xe3\x92\xff\x5e\xf1\xd9\x98\x7b\x78" +
	"\xe1\xf0\xe5\x3c\xe4\x89\xdb\xd8\x87\xb5\xe4\x3a\x0b\x4e\x22\xab\xf2\x5a\xd1\x53\x8b\xb5\x4a\x6f\x25\x1e\x71\x18" +
	"\x38\x86\xba\x2a\xd8\x5d\xc0\x18\x63\x83\x03\x26\x22\x35\x5a\x63\x4a\x72\x2f\xe9\x6d\xf0\x7f\xab\x6f\x89\x4a\x61" +
	"\xf1\xa5\x42\x47\xce\x8b\x99\x76\xb5\xd9\x49\x47\xa8\xd3\x53\xaf\xd4\x7b\x50\x32\xf3\x67\x84\x92\x1a\xbd\x21\xb1" +
	"\x32\xcb\x51\x58\x84\x74\x0b\x89\x54\xbd\x7b\x28\x2d\xfd\xfd\x27\xe9\xb6\x08\x19\x5a\xb1\x91\xa8\x32\x51\x80\x96" +
	"\x65\xa5\x80\xa4\xd1\xa7\x2e\xe3\xba\x63\x45\xa5\x48\x8a\xd2\x1a\x32\xa9\x51\x82\x2c\xa4\x68\x4d\x45\x1d\x45\x81" +
	"\xb8\x13\x1b\x6b\x34\x61\x87\xe9\x9a\xd6\x7b\xc7\x61\x0b\xe4\xa0\x2c\xfd\xf7\x1e\xb4\x54\x0a\x04\x19\xeb\xa5\x0d" +
	"\xa4\x98\x18\xb3\x13\x05\x3a\x87\x3a\xc7\x2e\xa2\x33\x1a\x04\xf7\xc3\x20\x18\x47\x3c\x8c\xf9\x85\x8e\xfb\xe8\xd9" +
	"\x26\x34\xed\x38\x1a\xd9\xd3\x2c\xe6\xdf\x78\xc4\x26\xfc\x6b\xb8\x9a\xc6\x4c\xe3\x2b\xed\x41\xdd\x0d\x7a\x99\x06" +
	"\x8f\x8f\x16\xf3\x54\x81\x73\xf7\x6c\x11\x3d\xfd\x0a\xa3\x35\xfb\xc9\xd7\x6c\x36\x8f\x9b\xe1\xad\xa9\xde\xc7\xba" +
	"\x6d\x7e\xad\x9d\x4d\xec\x87\xc4\xa7\x83\xf8\x2f\x72\xbb\xa6\x0d\xb3\xb7\x5e\x86\xee\x27\xbb\x89\xba\xae\x2c\xe4" +
	"\xdd\x2a\xd6\x12\xbe\x92\x85\x1e\xf0\xa7\x7b\xd9\x02\xdd\xb0\xc5\x7d\xfe\x4f\xec\xef\xff\x9e\xae\xae\xb7\x98\xbb" +
	"\x82\x5f\x59\xe9\xbf\x01\x00\x00\xff\xff\x93\xf2\x5e\x9a\x49\x05\x00\x00")

func commonDataMigrations3_add_job_type_tablesSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations3_add_job_type_tablesSql,
		"common/data/migrations/3_add_job_type_tables.sql",
	)
}

func commonDataMigrations3_add_job_type_tablesSql() (*asset, error) {
	bytes, err := commonDataMigrations3_add_job_type_tablesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/3_add_job_type_tables.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations4_rendezvous_tablesSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\xc1\x72\xa3\x38\x10\xbd\xf3\x15\x7d\xc3\xae\x4d\x2e\xc9\x2d\x3e" +
	"\x11\x47\xd9\x50\x8b\xc1\xc1\xb0\xbb\x99\xa9\x29\x4a\x81\xb6\xad\x2a\x59\xf2\x20\x39\x33\x99\xaf\x9f\x12\x06\x82" +
	"\x00\x33\x39\xaa\xfb\xa9\xbb\xdf\xd3\x93\x74\x7d\x0d\x7f\x1d\xd8\xae\xa4\x1a\xe1\x41\xfe\x10\x4e\x37\xb0\xd1\x54" +
	"\xe3\x01\x85\xbe\xc7\x1d\x13\x8e\xf3\x10\x47\x6b\x48\xbc\xfb\x80\x80\xff\x08\xe4\x7f\x7f\x93\x6c\x20\x97\x9c\x63" +
	"\xae\x65\xa9\x16\xe3\x80\x53\xc9\x2f\xa5\x34\x2a\x9d\xed\x91\x1f\x71\x6a\x77\x96\x53\x8d\x3b\x59\x32\x54\x8b\x7a" +
	"\x88\x0d\x79\x4e\x49\xb8\xb4\xe6\xa0\x3a\x13\x32\x53\xf8\x7d\x71\x11\x63\x8a\xfd\x09\x93\xcb\x93\xd0\xe5\xfb\x27" +
	"\x70\x35\xef\x16\x59\x13\x78\x59\x77\x61\xcb\x28\x08\xc8\x32\x89\xe2\xcc\x24\x16\xce\xb8\xc0\x44\x14\x76\x26\x3d" +
	"\x4e\x9e\xc4\x32\x26\x5e\x42\xce\xbd\xec\x0e\xe0\x6d\x80\x84\xe9\x0a\x66\x0e\x00\x80\xbb\xd7\xfa\xa8\xdc\xab\xf3" +
	"\x42\x0a\x26\x45\xb3\x28\xe4\x81\x32\x91\x6d\x4b\x29\x34\x16\xae\x33\x5f\xb4\x75\x5b\xba\x43\x92\x4d\xe7\xe6\x98" +
	"\xc2\x28\x19\x3a\xc1\x31\xcd\xbb\x7b\xc1\x0f\x13\xf2\x37\x89\xe1\x81\x3c\x7a\x69\x90\x80\xc0\x9f\xfa\x8d\xf2\x99" +
	"\xdb\xef\xe0\xde\xdd\x95\xb8\xcb\x39\x55\x6a\x0e\xeb\xd8\x5f\x79\xf1\x0b\xfc\x43\x5e\xaa\x46\x61\x1a\x04\x66\x7a" +
	"\xfd\x7e\xc4\x1e\x6f\x13\xa6\x45\x51\xa2\x52\xf0\xaf\x17\x2f\x9f\xbc\xd8\x84\x2a\x76\xd9\x99\x6a\x13\xaf\x98\x0e" +
	"\xa8\xda\x5c\xba\x4e\x99\xa0\x6c\xbc\xed\x9c\x95\x3e\xef\x30\x44\xdb\x51\xa1\xc0\x2d\x3d\x71\xfd\xc1\xf6\xa3\xac" +
	"\x6b\xb1\xbb\x6a\x4a\x34\x33\x5a\x74\xa1\x36\xb7\x55\xbc\x4e\xb4\x6e\x1d\x49\x16\x54\x63\x46\x8b\x02\x0b\x48\xfc" +
	"\x15\xd9\x24\xde\x6a\x0d\xff\xf9\xc9\x53\xb5\x84\x2f\x51\x48\x7a\x3b\x94\x3c\x95\x39\x76\x05\x04\x10\x52\xa3\xb2" +
	"\x43\x34\xd7\xec\x0d\xe1\x3e\x8a\x02\xe2\x85\xbd\x1a\x69\xe8\x3f\xa7\x04\x66\xa7\x92\x5f\x75\xe6\x9b\x1b\xd9\x73" +
	"\x79\x30\x2e\x06\x29\x40\xd3\x57\x8e\x95\x80\xc0\x14\xb8\x4b\x29\x34\x65\x42\x01\x13\x5b\x59\x1e\xa8\x66\x52\x18" +
	"\x58\x1a\x07\x26\x96\xf3\x93\xa1\xc1\x04\xe8\x3d\x42\xce\x34\xfb\x85\x82\xd3\x57\x93\x07\xce\x94\x76\x47\xdc\xdb" +
	"\x79\x5e\x3e\x73\x98\xdd\xd7\xa8\x3e\x54\xbb\xc2\x65\x17\x0f\x3b\x7d\xca\xc7\x00\x82\x1e\xf0\xc2\x99\xf7\xdc\xdc" +
	"\xcb\x56\x57\xa0\x31\xf4\xe8\xdd\xed\x5f\x4e\xfb\x59\x9b\xbc\xc7\x06\xca\xb0\x11\xe1\x82\xc7\x86\xee\xb6\x7b\x8c" +
	"\x39\x7c\x7b\xe2\x3c\xb3\x38\xd7\x6e\x99\x50\x65\x14\x41\xf9\x71\x4f\xb3\x1b\x30\x88\xd9\xcd\x7c\x0a\x74\x7b\x06" +
	"\xdd\x0e\x40\xa3\x86\x6c\xc9\xdb\xae\xac\xa9\x55\x93\x29\xa0\xa2\x00\x7f\x13\x41\x2e\x0b\x54\x63\xce\xeb\x09\xda" +
	"\xf9\x9a\xa6\x1f\x93\xce\x57\xd7\x88\x3f\xbc\xf9\x23\xc2\xb7\x0d\xc6\x44\x37\x59\x33\xe9\xb4\xa2\x06\x55\xa0\xca" +
	"\x27\x5e\x20\x2e\xc5\xce\xc2\x7c\xa4\x24\x2f\xaa\x1e\xad\x5d\xbf\x7e\xab\x4c\x79\xf1\xb7\xfb\x1d\x00\x00\xff\xff" +
	"\x2d\xa1\xf1\xcf\x79\x08\x00\x00")

func commonDataMigrations4_rendezvous_tablesSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations4_rendezvous_tablesSql,
		"common/data/migrations/4_rendezvous_tables.sql",
	)
}

func commonDataMigrations4_rendezvous_tablesSql() (*asset, error) {
	bytes, err := commonDataMigrations4_rendezvous_tablesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/4_rendezvous_tables.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _commonDataMigrations5_experiment_tablesSql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x96\xcd\x72\xdb\x36\x10\xc7\xcf\xc2\x53\xec\xf0\x12\x6b\xa2\x68\xda" +
	"\x1e\xed\xe9\x81\xb6\x90\x84\x89\x44\x3a\x22\xd4\xd6\xed\x74\x38\x10\xb9\xa6\x11\xf3\xab\x00\x54\xd9\x6f\xdf\x01" +
	"\x28\x8a\xa0\x24\xbb\xea\x8c\x2f\x3d\x09\xd8\x5d\x2c\x16\xc0\x7f\x7f\x22\xf9\xf0\x01\xde\x97\x22\x97\x5c\x23\xcc" +
	"\xea\x6d\x35\x30\xc4\x9a\x6b\x2c\xb1\xd2\xd7\x98\x8b\x8a\x10\x7f\xce\xe8\x12\x98\x7f\x3d\xa7\xd0\x6c\xd6\x85\x48" +
	"\xa7\xdf\xeb\xb5\x82\x25\x0d\xfd\x05\x85\x9b\x68\xbe\x5a\x84\xe0\xe1\x53\x83\x52\x98\x75\x49\x55\x7b\xc0\x22\xf0" +
	"\x34\x57\x8f\x66\x72\x45\xc8\xcd\x92\xfa\x8c\x42\x4c\xbf\xad\x68\x78\x43\x61\xe7\x4a\x14\xfe\x75\x45\x66\xcb\xe8" +
	"\xb6\x77\x0d\x12\x39\x01\x6d\x05\xdf\xeb\x75\xd2\x47\xa8\x3e\x75\xef\x36\xb9\x15\x5c\x10\xe8\x76\x81\x20\x64\xf4" +
	"\x13\x5d\xc2\x8c\x7e\xf4\x57\x73\x06\x15\x3e\xe9\xbf\x79\x71\xf1\xce\x29\xe3\xdd\xe5\xa5\xc4\x3c\x2d\xb8\x52\x63" +
	"\xb8\x5d\x06\x0b\x7f\x79\x07\x5f\xe9\x1d\x84\x11\x83\x70\x35\x9f\x4f\x4c\x42\x54\x3a\xa9\x78\x89\x10\x45\x8c\xc6" +
	"\xcc\xd8\xb8\xcc\x37\xb6\x16\xf8\x12\x47\xe1\x35\x19\x5f\x11\x12\xd3\x39\xbd\x61\x64\x64\xca\xe1\x05\x4a\xad\xa6" +
	"\xf6\x27\xa9\xea\xc9\xc0\x5a\xa2\x52\x3c\xc7\xa1\x11\x9f\xb4\xe4\x13\x12\x84\x2c\x82\xde\x9c\xe8\xb2\x21\x1f\x97" +
	"\xd1\xc2\xb1\x5d\x11\xf3\x74\xec\x01\x21\xab\xb7\x15\xb4\x4f\x28\xea\x0a\x54\x2a\x45\xa3\x41\x28\x28\xc4\x23\x16" +
	"\xcf\xb0\x96\xf5\x23\x56\x70\x2f\xeb\x12\x1e\x50\xb6\x0b\xb6\x5c\x66\x6a\xba\xaf\x56\x64\x13\x32\x4a\xeb\xd2\x9c" +
	"\xc6\x8c\x24\xda\x64\x89\x16\xa5\xa9\x50\xa5\x0f\x98\x6d\x0a\x33\xcc\xb0\xe0\xcf\x13\x32\xd2\x5c\xe6\xa8\x93\xb4" +
	"\xde\x54\x5a\x0a\x54\xbd\xa9\x29\xb8\xbe\xaf\x65\x69\x4d\xa2\x44\x95\xc8\x4d\x35\x21\x23\x73\xf5\x66\x98\x70\xbd" +
	"\x3f\xa0\x3a\x71\xac\x83\x27\xef\xec\xae\x18\xcd\x4a\xe7\x6e\x3a\x41\x0e\x2e\xed\xea\x25\xfd\xee\xfc\xd0\x7a\x77" +
	"\x22\x06\xcf\x29\xcf\x03\x76\x77\x4b\x81\x05\x0b\x0a\xbf\x06\xec\x73\x3b\xfa\x3d\x0a\x29\xac\xe2\x20\xfc\x04\x4e" +
	"\xec\xe5\xa5\x39\xe3\xbf\x6d\x46\x46\x83\xdd\xbc\x4e\x12\x1e\xc4\x94\x1d\x8b\xb3\x73\x1f\xaa\x73\xf2\x7a\x9e\x5e" +
	"\xaf\x23\x7f\x36\x1b\x48\xf9\xa2\x0f\x1d\xb7\xd2\x39\xee\x7a\x5a\x65\x43\xcf\xaa\x79\x0b\x3c\x74\x2c\xb0\x60\x18" +
	"\xb2\xe2\x04\x1e\x4e\x30\xe0\xb8\xcf\x1d\x0c\xd8\x6e\x1f\x2c\x7a\xb9\xe7\x8f\x72\x9f\xd1\xf9\x04\x60\xd7\x17\xf0" +
	"\x8b\xbf\xbc\xf9\xec\x2f\x4d\xe7\x0f\x1a\xc4\xca\x23\x66\xfe\xe2\xf6\x40\x2d\x26\xb2\x6b\x1d\x77\xb5\xed\x21\x53" +
	"\xa6\x05\xcb\x41\x23\x75\x81\x17\x3f\x8d\xe1\x8f\x3f\x9d\x88\x7d\x5f\xed\x23\x7e\xfc\x61\x1f\xd2\xf5\x59\x97\xd5" +
	"\x51\xe8\xab\xe5\x09\x95\x64\x75\x85\x70\x1d\x45\x73\xea\x87\xb6\x62\xf3\xcc\xf0\x25\xba\x4e\x62\xe6\x33\x6a\xaf" +
	"\xa0\xa7\x9f\x73\x0c\x25\xf2\x4a\x54\x79\xf2\x88\xcf\x89\xc8\x0e\x3d\x98\x39\xcf\xd4\x39\xc9\xb8\xe7\xd6\x7d\x5d" +
	"\x14\xf5\x56\x54\xb9\x03\xaf\x66\xa3\x15\xf0\xa2\x00\xfd\x80\x60\x15\x0b\x12\x0b\xae\x31\x03\xcd\xd7\x05\x2a\x10" +
	"\x95\xae\x81\x43\x85\x5b\x50\xd8\x70\x23\x4c\x93\xd0\x7a\xdf\x06\xbe\xc6\xa4\xa6\x96\x88\x76\xd4\x63\xb1\x9d\x1e" +
	"\xb0\xd1\x1a\x1d\x40\xda\x79\x47\x49\x3b\x39\x81\x4a\xd7\xee\xf2\xb2\xb5\x3b\xd0\xb4\x06\xe7\x31\xcd\xed\xb6\x3b" +
	"\x9a\x57\xda\xcf\x76\xcf\x78\xce\x1f\x07\xf9\x12\x05\xa1\x5d\x05\x51\xfb\xbb\xbf\x23\xf8\x19\x4e\xdc\xdc\x11\x92" +
	"\xd5\xd1\xdf\xf2\x0b\x8c\x3e\x03\xd1\x56\x0b\x42\x81\x42\xad\xec\xa3\xbb\xc2\xd5\x35\xac\x11\x78\xab\x6e\xcd\xcb" +
	"\x06\xb6\x42\x3f\xd8\xb0\x8c\x6b\x9c\xbe\x05\xe0\x4f\x35\xc6\x8e\xf2\xe9\x46\x4a\xc3\x0a\xb3\x17\xbc\x77\x2b\xfb" +
	"\xbf\xe1\x7e\x00\xd0\xe0\xa3\x5d\x49\x7f\x0b\x62\x16\x43\x5a\x08\x73\x46\x87\xa8\xc4\x10\x15\x40\x64\xb0\x5a\x05" +
	"\xb3\xc1\x47\x10\x40\x23\xeb\x35\x26\x3b\x5f\x6b\x3a\x09\xdf\xd6\xc5\x65\xae\x12\x91\x3d\x99\xae\xc5\x1c\x65\x0b" +
	"\xab\x0e\x31\xcc\x8f\xbf\x76\x8c\xd9\xe5\xce\x25\x2a\xd5\x31\xec\xbf\x60\x16\xa0\xaa\xb5\xb8\x17\xe9\xd9\xf1\x3c" +
	"\x4d\xb1\xd1\xe7\x44\x9a\xce\x3a\x27\xae\xe0\x4a\x27\x9b\x26\xb3\xb8\x7a\x29\xd4\x12\xf0\xa0\x7d\xec\x67\xeb\xa0" +
	"\xa9\x5c\xcb\xe9\x2f\xe7\x81\x00\xbd\x56\x81\xde\xd4\xe3\xa9\xc5\x8c\xf2\x08\x0c\xf5\xef\x6d\x14\x4a\x43\xf0\x43" +
	"\xed\x00\x18\xf1\xac\xc2\xe0\xdb\x8a\xc2\x45\x17\x35\xee\x1c\xdd\xf2\x47\x7c\x16\x99\xd7\x71\xfc\xb5\x4f\x88\x7f" +
	"\x02\x00\x00\xff\xff\xe4\x43\xf1\x28\x64\x0c\x00\x00")

func commonDataMigrations5_experiment_tablesSqlBytes() ([]byte, error) {
	return bindataRead(
		_commonDataMigrations5_experiment_tablesSql,
		"common/data/migrations/5_experiment_tables.sql",
	)
}

func commonDataMigrations5_experiment_tablesSql() (*asset, error) {
	bytes, err := commonDataMigrations5_experiment_tablesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "common/data/migrations/5_experiment_tables.sql", size: 0, md5checksum: "", mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// AssetNames returns the names of the assets.
// nolint: deadcode
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"common/data/migrations/1_accounts_create.sql":      commonDataMigrations1_accounts_createSql,
	"common/data/migrations/1_active_probes_create.sql": commonDataMigrations1_active_probes_createSql,
	"common/data/migrations/1_jobs_create.sql":          commonDataMigrations1_jobs_createSql,
	"common/data/migrations/1_probe_updates_create.sql": commonDataMigrations1_probe_updates_createSql,
	"common/data/migrations/1_tasks_create.sql":         commonDataMigrations1_tasks_createSql,
	"common/data/migrations/2_add_jobs_state.sql":       commonDataMigrations2_add_jobs_stateSql,
	"common/data/migrations/2_add_language_column.sql":  commonDataMigrations2_add_language_columnSql,
	"common/data/migrations/3_add_job_type_tables.sql":  commonDataMigrations3_add_job_type_tablesSql,
	"common/data/migrations/4_rendezvous_tables.sql":    commonDataMigrations4_rendezvous_tablesSql,
	"common/data/migrations/5_experiment_tables.sql":    commonDataMigrations5_experiment_tablesSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"common": {nil, map[string]*bintree{
		"data": {nil, map[string]*bintree{
			"migrations": {nil, map[string]*bintree{
				"1_accounts_create.sql": {commonDataMigrations1_accounts_createSql, map[string]*bintree{}},
				"1_active_probes_create.sql": {commonDataMigrations1_active_probes_createSql, map[string]*bintree{}},
				"1_jobs_create.sql": {commonDataMigrations1_jobs_createSql, map[string]*bintree{}},
				"1_probe_updates_create.sql": {commonDataMigrations1_probe_updates_createSql, map[string]*bintree{}},
				"1_tasks_create.sql": {commonDataMigrations1_tasks_createSql, map[string]*bintree{}},
				"2_add_jobs_state.sql": {commonDataMigrations2_add_jobs_stateSql, map[string]*bintree{}},
				"2_add_language_column.sql": {commonDataMigrations2_add_language_columnSql, map[string]*bintree{}},
				"3_add_job_type_tables.sql": {commonDataMigrations3_add_job_type_tablesSql, map[string]*bintree{}},
				"4_rendezvous_tables.sql": {commonDataMigrations4_rendezvous_tablesSql, map[string]*bintree{}},
				"5_experiment_tables.sql": {commonDataMigrations5_experiment_tablesSql, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
