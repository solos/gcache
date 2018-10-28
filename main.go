package main

import (
	"fmt"
	"gcache/config"
	pb "gcache/model"
	"github.com/dgraph-io/badger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var Std *Service

type Service struct {
	*ServOpt
	db     *badger.DB
	ctx    context.Context
	cancel context.CancelFunc
	mutex  sync.RWMutex
}

type ServOpt struct {
	bucket uint32
	total  uint32
}

func Init() {
	Std = NewService(&ServOpt{})
}

func NewService(opt *ServOpt) *Service {
	s := &Service{}
	s.Init(opt)
	return s
}

func (s *Service) Init(opt *ServOpt) error {
	s.ServOpt = opt
	ctx, cancelFunc := context.WithCancel(context.Background())
	s.ctx = ctx
	s.cancel = cancelFunc
	return nil
}

func (s *Service) Get(ctx context.Context, in *pb.GetReq) (*pb.GetResp, error) {

	key := in.Key
	val := []byte("")
	status := false
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		status = true
		return nil
	})
	return &pb.GetResp{Status: status, Value: string(val)}, err
}

func (s *Service) Set(ctx context.Context, in *pb.SetReq) (*pb.SetResp, error) {

	k := []byte(in.Key)
	v := []byte(in.Value)
	status := false
	err := s.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(k, v)
		if err == nil {
			status = true
		}
		return err
	})
	return &pb.SetResp{Status: status}, err
}

func main() {
	_, e := config.NewConfig("config.toml")
	if e != nil {
		fmt.Println(e)
		return
	}
	Init()

	opts := badger.DefaultOptions
	opts.Dir = config.DefaultConfig.Server.Dir
	opts.ValueDir = config.DefaultConfig.Server.ValueDir
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	Std.db = db

	addr := config.DefaultConfig.Server.Addr
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	serv := grpc.NewServer(
		grpc.MaxRecvMsgSize(256<<20),
		grpc.MaxSendMsgSize(256<<20),
	)
	pb.RegisterDataServer(serv, Std)
	serv.Serve(lis)
}
