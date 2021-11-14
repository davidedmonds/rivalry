package backend_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/rs/xid"
	"github.com/stretchr/testify/suite"
	"om-stream/internal/app/backend"
	"om-stream/internal/backoff"
	customlogic "om-stream/internal/managers/customlogic/mock"
	matches "om-stream/internal/managers/matches/mock"
	ticketsMock "om-stream/internal/managers/tickets/mock"
	"om-stream/pkg/pb"
)

type DispenserTestSuite struct {
	suite.Suite
	ctx                context.Context
	cancel             func()
	matchesManager     *matches.MockManager
	ticketsManager     *ticketsMock.MockManager
	customlogicManager *customlogic.MockAssignmentManager
	dispenser          backend.Dispenser
}

func (s *DispenserTestSuite) SetupTest() {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.ticketsManager = ticketsMock.NewMockManager(gomock.NewController(s.T()))
	s.matchesManager = matches.NewMockManager(gomock.NewController(s.T()))
	s.customlogicManager = customlogic.NewMockAssignmentManager(gomock.NewController(s.T()))
	s.dispenser = backend.NewDispenser(s.matchesManager, s.ticketsManager, s.customlogicManager)
}

func (s *DispenserTestSuite) TestDispenser() {
	assignment := &pb.Assignment{Connection: "foo"}
	match := &pb.Match{
		MatchId: xid.New().String(),
		Tickets: []*pb.Ticket{
			{Id: xid.New().String()},
		},
	}
	s.matchesManager.EXPECT().StreamMatches(s.ctx, gomock.Any()).Do(
		func(ctx context.Context, f func(ctx context.Context, match *pb.Match)) error {
			f(ctx, match)
			return nil
		})
	s.customlogicManager.EXPECT().MakeAssignment(s.ctx, gomock.Any()).Return(assignment, nil)
	s.ticketsManager.EXPECT().AddAssignmentToTickets(s.ctx, assignment, match.Tickets)
	go func() {
		time.Sleep(time.Second / 2)
		s.cancel()
	}()
	s.dispenser.Run(s.ctx)
}

func (s *DispenserTestSuite) TestDispenserRequeueFailedAssignment() {
	match := &pb.Match{
		MatchId: xid.New().String(),
		Tickets: []*pb.Ticket{
			{Id: xid.New().String()},
		},
	}
	s.matchesManager.EXPECT().StreamMatches(s.ctx, gomock.Any()).Do(
		func(ctx context.Context, f func(ctx context.Context, match *pb.Match)) error {
			f(ctx, match)
			return nil
		})
	s.customlogicManager.EXPECT().MakeAssignment(s.ctx, gomock.Any()).
		Return(nil, backoff.Permanent(fmt.Errorf("fail")))
	s.matchesManager.EXPECT().Requeue(match.MatchId)
	go func() {
		time.Sleep(time.Second / 2)
		s.cancel()
	}()
	s.dispenser.Run(s.ctx)
}

func TestDispenserTestSuite(t *testing.T) {
	suite.Run(t, new(DispenserTestSuite))
}
