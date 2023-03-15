// Copyright Mia srl
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httphandler

import (
	"testing"

	"github.com/mia-platform/miactl/internal/cmd/login"
	"github.com/stretchr/testify/require"
)

func TestClientBuilding(t *testing.T) {
	mExpected := MiaClient{
		request: Request{
			url: "url",
		},
		auth: &Auth{
			browser:    login.Browser{},
			providerID: "id",
			url:        "url",
		},
	}

	m2Expected := MiaClient{
		request: Request{
			url: "url2",
		},
	}

	b := login.Browser{}
	r := Request{
		url: "url",
	}
	r2 := Request{
		url: "url2",
	}

	m := NewMiaClientBuilder().
		WithAuthentication("id", "url", b).
		WithRequest(r)

	m2 := NewMiaClientBuilder().
		WithRequest(r2)

	require.Equal(t, *m, mExpected)
	require.Equal(t, *m2, m2Expected)

}
