// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	multimapv1 "github.com/atomix/atomix/api/pkg/multimap/v1"
	"github.com/atomix/atomix/driver/pkg/driver"
)

type MultiMapProxy multimapv1.MultiMapServer

type MultiMapProvider interface {
	NewMultiMap(spec driver.PrimitiveSpec) (MultiMapProxy, error)
}
