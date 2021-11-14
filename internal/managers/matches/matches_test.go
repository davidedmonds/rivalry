package matches_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	kvMock "om-stream/internal/db/kv/mock"
	streamMock "om-stream/internal/db/stream/mock"
	"om-stream/internal/managers/matches"
	ticketsMock "om-stream/internal/managers/tickets/mock"
	"om-stream/pkg/pb"
)

var ctx = context.Background()

type MatchManagerTestSuite struct {
	suite.Suite
	manager        matches.Manager
	ticketsManager *ticketsMock.MockManager
	store          *kvMock.MockStore
	streamClient   *streamMock.MockClient
}

func (s *MatchManagerTestSuite) SetupTest() {
	s.ticketsManager = ticketsMock.NewMockManager(gomock.NewController(s.T()))
	s.store = kvMock.NewMockStore(gomock.NewController(s.T()))
	s.streamClient = streamMock.NewMockClient(gomock.NewController(s.T()))
	s.manager = matches.NewManager(s.ticketsManager, s.store, s.streamClient)
}

func (s *MatchManagerTestSuite) TearDownTest() {
	s.store.EXPECT().Close()
	s.streamClient.EXPECT().Close()
	s.manager.Close()
}

func (s *MatchManagerTestSuite) TestCreateMatch() {
	match := &pb.Match{AllocateGameserver: true}
	s.ticketsManager.EXPECT().AssignTicketsToMatch(ctx, match).DoAndReturn(
		func(_ context.Context, match *pb.Match) (bool, error) {
			assert.NotEmpty(s.T(), match.MatchId)
			return true, nil
		})
	s.store.EXPECT().Set(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	s.streamClient.EXPECT().SendMessage(matches.GetUnassignedMatchesTopic(), gomock.Any())
	outcome, err := s.manager.CreateMatch(ctx, match)
	assert.NoError(s.T(), err)
	assert.True(s.T(), outcome)
}

func (s *MatchManagerTestSuite) TestCreateMatchWithoutAllocation() {
	match := &pb.Match{AllocateGameserver: false}
	s.ticketsManager.EXPECT().AssignTicketsToMatch(ctx, match).DoAndReturn(
		func(_ context.Context, match *pb.Match) (bool, error) {
			assert.NotEmpty(s.T(), match.MatchId)
			return true, nil
		})
	s.store.EXPECT().Set(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	outcome, err := s.manager.CreateMatch(ctx, match)
	assert.NoError(s.T(), err)
	assert.True(s.T(), outcome)
}

func (s *MatchManagerTestSuite) TestCreateMatchesFailSetReleasesTickets() {
	match := &pb.Match{AllocateGameserver: true}
	s.ticketsManager.EXPECT().AssignTicketsToMatch(ctx, match).DoAndReturn(
		func(_ context.Context, match *pb.Match) (bool, error) {
			assert.NotEmpty(s.T(), match.MatchId)
			return true, nil
		})
	s.store.EXPECT().Set(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(fmt.Errorf("fail"))
	s.ticketsManager.EXPECT().ReleaseTicketsFromMatch(ctx, match)
	outcome, err := s.manager.CreateMatch(ctx, match)
	assert.Error(s.T(), err)
	assert.False(s.T(), outcome)
}

func (s *MatchManagerTestSuite) TestCreateMatchesFailSendMessageReleasesTicketsAndDeletesMatch() {
	match := &pb.Match{AllocateGameserver: true}
	s.ticketsManager.EXPECT().AssignTicketsToMatch(ctx, match).DoAndReturn(
		func(_ context.Context, match *pb.Match) (bool, error) {
			assert.NotEmpty(s.T(), match.MatchId)
			return true, nil
		})
	s.store.EXPECT().Set(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
	s.streamClient.EXPECT().SendMessage(matches.GetUnassignedMatchesTopic(), gomock.Any()).Return(fmt.Errorf("fail"))
	s.ticketsManager.EXPECT().ReleaseTicketsFromMatch(ctx, match)
	s.store.EXPECT().Del(ctx, gomock.Any(), gomock.Any())
	outcome, err := s.manager.CreateMatch(ctx, match)
	assert.Error(s.T(), err)
	assert.False(s.T(), outcome)
}

func (s *MatchManagerTestSuite) TestStreamMatches() {
	match := &pb.Match{MatchId: xid.New().String()}
	matchBytes, err := proto.Marshal(match)
	require.NoError(s.T(), err)
	s.streamClient.EXPECT().Subscribe(matches.GetUnassignedMatchesTopic(), gomock.Any()).Do(
		func(topic string, f func([]byte)) error {
			f([]byte(match.MatchId))
			return nil
		})
	s.store.EXPECT().Get(ctx, gomock.Any(), match.MatchId).Return(matchBytes, nil)
	assert.NoError(s.T(), s.manager.StreamMatches(ctx,
		func(ctx context.Context, m *pb.Match) {
			assert.Equal(s.T(), match.MatchId, m.MatchId)
		}))
}

func (s *MatchManagerTestSuite) TestRequeue() {
	s.streamClient.EXPECT().SendMessage(matches.GetUnassignedMatchesTopic(), gomock.Any())
	s.manager.Requeue(xid.New().String())
}

func TestMatchManagerTestSuite(t *testing.T) {
	suite.Run(t, new(MatchManagerTestSuite))
}
