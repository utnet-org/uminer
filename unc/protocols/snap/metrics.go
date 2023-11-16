// Copyright 2023 The go-utility Authors
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

package snap

import (
	metrics "github.com/yanhuangpai/go-utility/metrics"
)

var (
	ingressRegistrationErrorName = "unc/protocols/snap/ingress/registration/error"
	egressRegistrationErrorName  = "unc/protocols/snap/egress/registration/error"

	IngressRegistrationErrorMeter = metrics.NewRegisteredMeter(ingressRegistrationErrorName, nil)
	EgressRegistrationErrorMeter  = metrics.NewRegisteredMeter(egressRegistrationErrorName, nil)

	// deletionGauge is the metric to track how many trie node deletions
	// are performed in total during the sync process.
	deletionGauge = metrics.NewRegisteredGauge("unc/protocols/snap/sync/delete", nil)

	// lookupGauge is the metric to track how many trie node lookups are
	// performed to determine if node needs to be deleted.
	lookupGauge = metrics.NewRegisteredGauge("unc/protocols/snap/sync/lookup", nil)

	// boundaryAccountNodesGauge is the metric to track how many boundary trie
	// nodes in account trie are met.
	boundaryAccountNodesGauge = metrics.NewRegisteredGauge("unc/protocols/snap/sync/boundary/account", nil)

	// boundaryAccountNodesGauge is the metric to track how many boundary trie
	// nodes in storage tries are met.
	boundaryStorageNodesGauge = metrics.NewRegisteredGauge("unc/protocols/snap/sync/boundary/storage", nil)

	// smallStorageGauge is the metric to track how many storages are small enough
	// to retrieved in one or two request.
	smallStorageGauge = metrics.NewRegisteredGauge("unc/protocols/snap/sync/storage/small", nil)

	// largeStorageGauge is the metric to track how many storages are large enough
	// to retrieved concurrently.
	largeStorageGauge = metrics.NewRegisteredGauge("unc/protocols/snap/sync/storage/large", nil)

	// skipStorageHealingGauge is the metric to track how many storages are retrieved
	// in multiple requests but healing is not necessary.
	skipStorageHealingGauge = metrics.NewRegisteredGauge("unc/protocols/snap/sync/storage/noheal", nil)
)
