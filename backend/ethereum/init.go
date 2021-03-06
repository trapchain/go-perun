// Copyright 2019 - See NOTICE file for copyright holders.
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

package ethereum

import (
	_ "perun.network/go-perun/backend/ethereum/channel"      // backend init
	_ "perun.network/go-perun/backend/ethereum/channel/test" // backend init
	_ "perun.network/go-perun/backend/ethereum/wallet"       // backend init
	_ "perun.network/go-perun/backend/ethereum/wallet/test"  // backend init
)
