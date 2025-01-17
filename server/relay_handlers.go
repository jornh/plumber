package server

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/common"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/plumber/backends"
	"github.com/batchcorp/plumber/embed/etcd"
	"github.com/batchcorp/plumber/server/types"
	"github.com/batchcorp/plumber/validate"
)

func (s *Server) GetAllRelays(_ context.Context, req *protos.GetAllRelaysRequest) (*protos.GetAllRelaysResponse, error) {
	if err := s.validateAuth(req.Auth); err != nil {
		return nil, CustomError(common.Code_UNAUTHENTICATED, fmt.Sprintf("invalid auth: %s", err))
	}

	RelayOptions := make([]*opts.RelayOptions, 0)
	for _, v := range s.PersistentConfig.Relays {
		RelayOptions = append(RelayOptions, v.Options)
	}

	return &protos.GetAllRelaysResponse{
		Status: &common.Status{
			Code:      common.Code_OK,
			RequestId: uuid.NewV4().String(),
		},
		Opts: RelayOptions,
	}, nil
}

func (s *Server) CreateRelay(ctx context.Context, req *protos.CreateRelayRequest) (*protos.CreateRelayResponse, error) {
	if err := s.validateAuth(req.Auth); err != nil {
		return nil, CustomError(common.Code_UNAUTHENTICATED, fmt.Sprintf("invalid auth: %s", err))
	}

	if err := validate.RelayOptionsForServer(req.Opts); err != nil {
		return nil, CustomError(common.Code_ABORTED, fmt.Sprintf("unable to validate relay options: %s", err))
	}

	// Get stored connection information
	conn := s.PersistentConfig.GetConnection(req.Opts.ConnectionId)
	if conn == nil {
		return nil, CustomError(common.Code_NOT_FOUND, validate.ErrConnectionNotFound.Error())
	}

	// Try to create a backend from given connection options
	be, err := backends.New(conn)
	if err != nil {
		return nil, CustomError(common.Code_ABORTED, fmt.Sprintf("unable to create backend: %s", err))
	}

	// Save to etcd
	data, err := proto.Marshal(req.Opts)
	if err != nil {
		return nil, CustomError(common.Code_ABORTED, "could not marshal relay")
	}

	_, err = s.Etcd.Put(ctx, etcd.CacheRelaysPrefix+"/"+req.Opts.XRelayId, string(data))
	if err != nil {
		s.rollbackCreateRelay(ctx, req.Opts)
		return nil, CustomError(common.Code_ABORTED, err.Error())
	}

	// Used to shutdown relays on StopRelay() gRPC call
	shutdownCtx, shutdownFunc := context.WithCancel(context.Background())

	r := &types.Relay{
		Id:         uuid.NewV4().String(),
		Backend:    be,
		CancelFunc: shutdownFunc,
		CancelCtx:  shutdownCtx,
		Options:    req.Opts,
	}

	// Save to memory
	s.PersistentConfig.SetRelay(r.Id, r)

	// Publish CreateSchema event
	if err := s.Etcd.PublishCreateRelay(ctx, r.Options); err != nil {
		s.rollbackCreateRelay(ctx, req.Opts)
		s.Log.Error(err)
	}

	if err := r.StartRelay(); err != nil {
		s.rollbackCreateRelay(ctx, req.Opts)
		return nil, errors.Wrap(err, "unable to start relay")
	}

	r.Options.XRelayId = r.Id
	r.Active = true

	s.Log.Infof("Relay '%s' started", r.Id)

	return &protos.CreateRelayResponse{
		Status: &common.Status{
			Code:      common.Code_OK,
			Message:   "Relay started",
			RequestId: uuid.NewV4().String(),
		},
		RelayId: r.Id,
	}, nil
}

func (s *Server) rollbackCreateRelay(ctx context.Context, req *opts.RelayOptions) {
	// TODO: Implement
}

func (s *Server) UpdateRelay(_ context.Context, req *protos.UpdateRelayRequest) (*protos.UpdateRelayResponse, error) {
	if err := s.validateAuth(req.Auth); err != nil {
		return nil, CustomError(common.Code_UNAUTHENTICATED, fmt.Sprintf("invalid auth: %s", err))
	}

	// Stop existing relay
	relay := s.PersistentConfig.GetRelay(req.RelayId)
	if relay == nil {
		return nil, CustomError(common.Code_NOT_FOUND, "relay does not exist")
	}

	// Stop workers
	relay.CancelFunc()

	s.Log.Infof("Relay '%s' stopped", relay.Id)

	// TODO: need to get signal of when relay shutdown is complete
	time.Sleep(time.Second)

	// Update relay config
	relay.Options = req.Opts

	// Give relayer new ctx & cancel func
	ctx, cancelFunc := context.WithCancel(context.Background())

	relay.CancelCtx = ctx
	relay.CancelFunc = cancelFunc

	if err := relay.StartRelay(); err != nil {
		return nil, errors.Wrap(err, "unable to start relay")
	}

	data, err := proto.Marshal(relay.Options)
	if err != nil {
		s.rollbackUpdateRelay(ctx, req.Opts)
		fullErr := fmt.Sprintf("unable to marshal relay options to bytes for relay id '%s': %s", req.Opts.XRelayId, err)
		s.Log.Error(fullErr)

		return nil, CustomError(common.Code_ABORTED, fullErr)
	}

	// Save to etcd
	_, err = s.Etcd.Put(ctx, etcd.CacheRelaysPrefix+"/"+req.Opts.XRelayId, string(data))
	if err != nil {
		s.rollbackUpdateRelay(ctx, req.Opts)
		fullErr := fmt.Sprintf("unable to save new relay options to etcd for relay id '%s': %s", req.Opts.XRelayId, err)
		s.Log.Error(fullErr)

		return nil, CustomError(common.Code_ABORTED, fullErr)
	}

	// Save to memory
	s.PersistentConfig.SetRelay(relay.Id, relay)

	// Publish CreateSchema event
	if err := s.Etcd.PublishUpdateRelay(ctx, relay.Options); err != nil {
		s.rollbackUpdateRelay(ctx, req.Opts)
		fullErr := fmt.Sprintf("unable to publish update relay event for relay id '%s': %s", req.Opts.XRelayId, err)
		s.Log.Error(fullErr)

		return nil, CustomError(common.Code_ABORTED, fullErr)
	}

	s.Log.Infof("Relay '%s' started", relay.Id)

	relay.Active = true

	return &protos.UpdateRelayResponse{
		Status: &common.Status{
			Code:      common.Code_OK,
			Message:   "Relay updated",
			RequestId: uuid.NewV4().String(),
		},
	}, nil
}

// TODO: Implement
func (s *Server) rollbackUpdateRelay(ctx context.Context, relayOptions *opts.RelayOptions) {
	// Remove from etcd

	// Stop grpc relayers

	// Remove from persistent storage
}

func (s *Server) StopRelay(_ context.Context, req *protos.StopRelayRequest) (*protos.StopRelayResponse, error) {
	if err := s.validateAuth(req.Auth); err != nil {
		return nil, CustomError(common.Code_UNAUTHENTICATED, fmt.Sprintf("invalid auth: %s", err))
	}

	relay := s.PersistentConfig.GetRelay(req.RelayId)
	if relay == nil {
		return nil, CustomError(common.Code_NOT_FOUND, "relay does not exist")
	}

	if !relay.Active {
		return nil, CustomError(common.Code_FAILED_PRECONDITION, "Relay is not active")
	}

	// Stop workers
	relay.CancelFunc()
	relay.Active = false

	s.Log.Infof("Relay '%s' stopped", relay.Id)

	return &protos.StopRelayResponse{
		Status: &common.Status{
			Code:      common.Code_OK,
			Message:   "Relay stopped",
			RequestId: uuid.NewV4().String(),
		},
	}, nil
}

func (s *Server) ResumeRelay(ctx context.Context, req *protos.ResumeRelayRequest) (*protos.ResumeRelayResponse, error) {
	if err := s.validateAuth(req.Auth); err != nil {
		return nil, CustomError(common.Code_UNAUTHENTICATED, fmt.Sprintf("invalid auth: %s", err))
	}

	relay := s.PersistentConfig.GetRelay(req.RelayId)
	if relay == nil {
		return nil, CustomError(common.Code_NOT_FOUND, "relay does not exist")
	}

	if relay.Active {
		return nil, CustomError(common.Code_FAILED_PRECONDITION, "Relay is not stopped")
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	relay.CancelCtx = ctx
	relay.CancelFunc = cancelFunc

	if err := relay.StartRelay(); err != nil {
		return nil, errors.Wrap(err, "unable to start relay")
	}

	s.Log.Infof("Relay '%s' started", relay.Id)

	relay.Active = true

	return &protos.ResumeRelayResponse{
		Status: &common.Status{
			Code:      common.Code_OK,
			Message:   "Relay resumed",
			RequestId: uuid.NewV4().String(),
		},
	}, nil
}

func (s *Server) DeleteRelay(ctx context.Context, req *protos.DeleteRelayRequest) (*protos.DeleteRelayResponse, error) {
	if err := s.validateAuth(req.Auth); err != nil {
		return nil, CustomError(common.Code_UNAUTHENTICATED, fmt.Sprintf("invalid auth: %s", err))
	}

	relay := s.PersistentConfig.GetRelay(req.RelayId)
	if relay == nil {
		return nil, CustomError(common.Code_NOT_FOUND, "relay does not exist")
	}

	// Stop workers
	relay.CancelFunc()

	// Delete in etcd
	_, err := s.Etcd.Delete(ctx, etcd.CacheRelaysPrefix+"/"+relay.Id)
	if err != nil {
		return nil, CustomError(common.Code_INTERNAL, fmt.Sprintf("unable to delete connection: "+err.Error()))
	}

	// Delete in memory
	s.PersistentConfig.DeleteRelay(relay.Id)

	// Publish delete event
	if err := s.Etcd.PublishDeleteRelay(ctx, relay.Options); err != nil {
		s.Log.Error(err)
	}

	return &protos.DeleteRelayResponse{
		Status: &common.Status{
			Code:      common.Code_OK,
			Message:   "Relay deleted",
			RequestId: uuid.NewV4().String(),
		},
	}, nil
}
