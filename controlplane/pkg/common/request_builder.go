// Copyright (c) 2019 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/registry"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/model"
)

type RequestBuilder interface {
	Build(string, *registry.NSERegistration, *model.Forwarder, *connection.Connection) *networkservice.NetworkServiceRequest
}

type LocalRequestBuilder struct {
	nsmName     string
	idGenerator func() string
}

type RemoteRequestBuilder struct {
	nsmName     string
	idGenerator func() string
}

func createLocalNSERequest(connectionId string, idGenerator func() string, nsmName string, localMechanisms []*connection.Mechanism, requestConn *connection.Connection) *networkservice.NetworkServiceRequest {
	if connectionId == "" {
		connectionId = idGenerator()
	}
	return &networkservice.NetworkServiceRequest{
		Connection: &connection.Connection{
			Id:                     connectionId, // ID for NSE is managed by NSMgr
			NetworkService:         requestConn.GetNetworkService(),
			Context:                requestConn.GetContext(),
			Labels:                 requestConn.GetLabels(),
			NetworkServiceManagers: []string{nsmName},
		},
		MechanismPreferences: localMechanisms,
	}
}

func createRemoteNSMRequest(connectionId string, srcNsmName string, endpoint *registry.NSERegistration, remoteMechanisms []*connection.Mechanism, requestConn *connection.Connection) *networkservice.NetworkServiceRequest {
	return &networkservice.NetworkServiceRequest{
		Connection: &connection.Connection{
			Id:                         connectionId,
			NetworkService:             requestConn.GetNetworkService(),
			Context:                    requestConn.GetContext(),
			Labels:                     requestConn.GetLabels(),
			NetworkServiceEndpointName: endpoint.GetNetworkServiceEndpoint().GetName(),
			NetworkServiceManagers: []string{
				srcNsmName, // src
				endpoint.GetNetworkServiceManager().GetName(), // dst
			},
		},
		MechanismPreferences: remoteMechanisms,
	}
}

func (builder *LocalRequestBuilder) Build(connectionId string, endpoint *registry.NSERegistration, fwd *model.Forwarder, requestConn *connection.Connection) *networkservice.NetworkServiceRequest {
	if builder.nsmName == endpoint.GetNetworkServiceEndpoint().GetNetworkServiceManagerName() {
		return createLocalNSERequest(connectionId, builder.idGenerator, builder.nsmName, fwd.LocalMechanisms, requestConn)
	}
	return createRemoteNSMRequest(connectionId, builder.nsmName, endpoint, fwd.RemoteMechanisms, requestConn)
}

func (builder *RemoteRequestBuilder) Build(connectionId string, endpoint *registry.NSERegistration, fwd *model.Forwarder,
	requestConn *connection.Connection) *networkservice.NetworkServiceRequest {
	return createLocalNSERequest(connectionId, builder.idGenerator, builder.nsmName, fwd.LocalMechanisms, requestConn)
}

func NewLocalRequestBuilder(m model.Model) *LocalRequestBuilder {
	return &LocalRequestBuilder{
		nsmName: m.GetNsm().GetName(),
		idGenerator: func() string {
			return m.ConnectionID()
		},
	}
}

func NewRemoteRequestBuilder(m model.Model) *RemoteRequestBuilder {
	return &RemoteRequestBuilder{
		nsmName: m.GetNsm().GetName(),
		idGenerator: func() string {
			return m.ConnectionID()
		},
	}
}
