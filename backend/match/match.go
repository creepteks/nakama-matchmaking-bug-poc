package match

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

// Match holds an authoritative match instance
type Match struct{}

// Register the collection of functions with Nakama
func Register(initializer runtime.Initializer) error {
	if err := initializer.RegisterMatchmakerMatched(doMatchmaking); err != nil {
		return err
	}

	if err := initializer.RegisterMatch("testMatch", createMatch); err != nil {
		return err
	}

	return nil
}

func doMatchmaking(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	presences := make([]runtime.Presence, 0)
	for _, e := range entries {
		for k, v := range e.GetProperties() {
			logger.Info("Matched on '%s' value '%v'", k, v)
		}
		presences = append(presences, e.GetPresence())
	}

	matchId, err := nk.MatchCreate(ctx, "testMatch", map[string]interface{}{"joins": presences, "debug": true, "debug-verbose": false})
	if err != nil {
		return "", err
	}

	return matchId, nil
}

// createMatch creates an actual pvp match using the match module and returns its id to the match lobby
// the match lobby, then, notifies the players of the actual match id, so they can join to the match
func createMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	return &Match{}, nil
}

// MatchInit is called once when the match is created
func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	entries := params["joins"].([]runtime.Presence)
	for _, p := range entries {
		logger.Info(p.GetUsername())
	}

	state := &Match{}

	label := "skill=100-150"
	return state, 1, label
}

// MatchSignal implements runtime.Match
func (*Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return state, ""
}

// MatchJoinAttempt is called when a user tries to connect to match.
func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	return state, true, ""
}

// MatchJoin is called after some users have successfully attempted to join the match
func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	return state
}

// MatchLeave is called when a presence has left the match, whether intentionally or by DC
func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	return state
}

// MatchLoop is called on every server tick
func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	return state
}

// MatchTerminate is called just before the match is about to be destroyed
func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	return nil
}

// func beforeChannelJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, envelope *rtapi.Envelope) (*rtapi.Envelope, error) {
// 	logger.Info("Intercepted request to join channel '%v'", envelope.GetChannelJoin().Target)
// 	return envelope, nil
// }

// func afterGetAccount(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.Account) error {
// 	logger.Info("Intercepted response to get account '%v'", in)
// 	return nil
// }

// func eventSessionStart(ctx context.Context, logger runtime.Logger, evt *api.Event) {
// 	logger.Info("session start %v %v", ctx, evt)
// }

// func eventSessionEnd(ctx context.Context, logger runtime.Logger, evt *api.Event) {
// 	logger.Info("session end %v %v", ctx, evt)
// }
