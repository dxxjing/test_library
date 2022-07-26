package client

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"hantalk_go/pb_library/pb"
)

type TaskClient struct {
	conn *grpc.ClientConn
}

func NewTaskClient(addr string) (*TaskClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		err = errors.New("init task client err:" + err.Error())
		return nil, err
	}
	return &TaskClient{
		conn: conn,
	}, nil
}

func (c *TaskClient) WriteEsLog(req *pb.WriteEsLogReq) (rsp *pb.WriteEsLogRsp, err error) {
	defer c.conn.Close()
	rsp, err = pb.NewAsyncTaskClient(c.conn).WriteEsLog(context.Background(), req)
	if err != nil {
		err = errors.New("task clent: write es log err:" + err.Error())
		return
	}
	return
}
