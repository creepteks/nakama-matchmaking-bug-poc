// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"

	match "github.com/creepteks/davaa/backend/match"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {

	if err := match.Register(initializer); err != nil {
		return err
	}

	// if err := initializer.RegisterRpc("go_echo_sample", rpcEcho); err != nil {
	// 	return err
	// }
	// if err := initializer.RegisterBeforeRt("ChannelJoin", beforeChannelJoin); err != nil {
	// 	return err
	// }
	// if err := initializer.RegisterAfterGetAccount(afterGetAccount); err != nil {
	// 	return err
	// }

	// if err := initializer.RegisterEventSessionStart(eventSessionStart); err != nil {
	// 	return err
	// }
	// if err := initializer.RegisterEventSessionEnd(eventSessionEnd); err != nil {
	// 	return err
	// }
	// if err := initializer.RegisterEvent(func(ctx context.Context, logger runtime.Logger, evt *api.Event) {
	// 	logger.Info("Received event: %+v", evt)
	// }); err != nil {
	// 	return err
	// }
	return nil
}
