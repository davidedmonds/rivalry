package backend

import (
	"context"

	"github.com/amg84/om-stream/internal/backoff"
	"github.com/amg84/om-stream/internal/managers/customlogic"
	"github.com/amg84/om-stream/internal/managers/matches"
	"github.com/amg84/om-stream/internal/managers/tickets"
	"github.com/amg84/om-stream/pkg/pb"
	"github.com/rs/zerolog/log"
)

// Dispenser defines the interface for the dispenser backend
type Dispenser interface {
	Run(ctx context.Context) error
}

type dispenser struct {
	matchesManager     matches.Manager
	ticketsManager     tickets.Manager
	customlogicManager customlogic.AssignmentManager
}

// NewDispenser returns a new dispenser
func NewDispenser(matchesManager matches.Manager, ticketsManager tickets.Manager, customlogicManager customlogic.AssignmentManager) Dispenser {
	return &dispenser{
		matchesManager:     matchesManager,
		ticketsManager:     ticketsManager,
		customlogicManager: customlogicManager,
	}
}

// Run is the main method for the service
func (d *dispenser) Run(ctx context.Context) error {
	err := d.matchesManager.StreamMatches(ctx, func(ctx context.Context, match *pb.Match) {
		var (
			assignment *pb.Assignment
			err        error
		)

		err = backoff.Retry(ctx, func() error {
			var err error
			assignment, err = d.customlogicManager.MakeAssignment(ctx, match)
			return err
		}, backoff.Exponential())

		if err != nil {
			log.Err(err).Str("match_id", match.MatchId).Msg("failed to make assignment")

			err = backoff.Retry(ctx, func() error {
				return d.matchesManager.Requeue(match.MatchId)
			}, backoff.Exponential())
			if err != nil {
				log.Err(err).Str("match_id", match.MatchId).Msg("failed to requeue match")
			}
			return
		}

		err = backoff.Retry(ctx, func() error {
			return d.ticketsManager.AddAssignmentToTickets(ctx, assignment, match.Tickets)
		}, backoff.Exponential())
		if err != nil {
			log.Err(err).Str("match_id", match.MatchId).Msg("failed to add assignment to ticket")
		}
	})
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}
