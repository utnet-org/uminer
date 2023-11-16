// Copyright 2015 The go-utility Authors
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

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/yanhuangpai/go-utility/metrics"
)

var (
	headerInMeter      = metrics.NewRegisteredMeter("unc/downloader/headers/in", nil)
	headerReqTimer     = metrics.NewRegisteredTimer("unc/downloader/headers/req", nil)
	headerDropMeter    = metrics.NewRegisteredMeter("unc/downloader/headers/drop", nil)
	headerTimeoutMeter = metrics.NewRegisteredMeter("unc/downloader/headers/timeout", nil)

	bodyInMeter      = metrics.NewRegisteredMeter("unc/downloader/bodies/in", nil)
	bodyReqTimer     = metrics.NewRegisteredTimer("unc/downloader/bodies/req", nil)
	bodyDropMeter    = metrics.NewRegisteredMeter("unc/downloader/bodies/drop", nil)
	bodyTimeoutMeter = metrics.NewRegisteredMeter("unc/downloader/bodies/timeout", nil)

	receiptInMeter      = metrics.NewRegisteredMeter("unc/downloader/receipts/in", nil)
	receiptReqTimer     = metrics.NewRegisteredTimer("unc/downloader/receipts/req", nil)
	receiptDropMeter    = metrics.NewRegisteredMeter("unc/downloader/receipts/drop", nil)
	receiptTimeoutMeter = metrics.NewRegisteredMeter("unc/downloader/receipts/timeout", nil)

	throttleCounter = metrics.NewRegisteredCounter("unc/downloader/throttle", nil)
)
