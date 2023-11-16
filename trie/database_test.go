// Copyright 2019 The go-utility Authors
// This file is part of the go-utility library.
//
// The go-utility library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-utility library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-utility library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"github.com/yanhuangpai/go-utility/core/rawdb"
	"github.com/yanhuangpai/go-utility/trie/triedb/hashdb"
	"github.com/yanhuangpai/go-utility/trie/triedb/pathdb"
	"github.com/yanhuangpai/go-utility/uncdb"
)

// newTestDatabase initializes the trie database with specified scheme.
func newTestDatabase(diskdb uncdb.Database, scheme string) *Database {
	config := &Config{Preimages: false}
	if scheme == rawdb.HashScheme {
		config.HashDB = &hashdb.Config{
			CleanCacheSize: 0,
		} // disable clean cache
	} else {
		config.PathDB = &pathdb.Config{
			CleanCacheSize: 0,
			DirtyCacheSize: 0,
		} // disable clean/dirty cache
	}
	return NewDatabase(diskdb, config)
}
