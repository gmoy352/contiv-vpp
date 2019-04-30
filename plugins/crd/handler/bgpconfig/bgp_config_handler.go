/*
 * // Copyright (c) 2018 Cisco and/or its affiliates.
 * //
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at:
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

//go:generate protoc -I ./model --gogo_out=plugins=grpc:./model ./model/bgpconfig.proto

package bgpconfig

import (
	"sync"

	informers "github.com/contiv/vpp/plugins/crd/pkg/client/informers/externalversions/bgpconfig/v1"
	"github.com/gogo/protobuf/proto"
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/datasync/kvdbsync"
	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/logging"
)

const (
	minResyncTimeout = 100  // minimum timeout between resync attempts, in ms
	maxResyncTimeout = 1000 // maximum timeout between resync attempts, in ms
)

// Handler handler implements Handler interface,
type Handler struct {
	Deps

	// Data store sync status and the mutex that protects access to it
	dsSynced bool
	dsMutex  sync.Mutex

	broker KeyProtoValBroker
	prefix string
	kpc    K8sToProtoConverter

	syncStopCh chan bool
}

// Deps defines dependencies for BgpConfig CRD Handler.
type Deps struct {
	Log                logging.Logger
	Publish            *kvdbsync.Plugin // KeyProtoValWriter does not define Delete
	ControllerInformer informers.BgpConfigInformer
}

// KeyProtoValBroker defines handler's interface to the key-value data store. It
// defines a subset of operations from a generic cn-infra broker interface
// (keyval.ProtoBroker in cn-infra).
type KeyProtoValBroker interface {
	// Put <data> to ETCD or to any other key-value based data source.
	Put(key string, data proto.Message, opts ...datasync.PutOption) error

	// Delete data under the <key> in ETCD or in any other key-value based data
	// source.
	Delete(key string, opts ...datasync.DelOption) (existed bool, err error)

	// GetValue reads a value from etcd stored under the given key.
	GetValue(key string, reqObj proto.Message) (found bool, revision int64, err error)

	// List values stored in etcd under the given prefix.
	ListValues(prefix string) (keyval.ProtoKeyValIterator, error)
}

// K8sToProtoConverter defines the signature for a function converting k8s
// objects to bgp config protobuf objects.
type K8sToProtoConverter func(interface{}) (interface{}, string, bool)

// DsItems defines the structure holding items listed from the data store.
type DsItems map[string]interface{}

// Init initializes handler configuration
// BgpConfig Handler will be taking action on resource CRUD
func (h *Handler) Init() error {
	return nil
}

// ObjectCreated is called when a CRD object is created
func (h *Handler) ObjectCreated(obj interface{}) {
	h.Log.Debugf("Object created with value: %v", obj)

}

// ObjectDeleted is called when a CRD object is deleted
func (h *Handler) ObjectDeleted(obj interface{}) {
	h.Log.Debugf("Object deleted with value: %v", obj)

}

// ObjectUpdated is called when a CRD object is updated
func (h *Handler) ObjectUpdated(oldObj, newObj interface{}) {
	h.Log.Debugf("Object updated with value: %v", newObj)

}
