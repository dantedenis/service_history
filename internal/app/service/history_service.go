package service

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"service_history/internal/app/proto"
	"time"
)

var query = "select deal_time, coast from history_deal where pair_id=(select id from currency_pair where pair=$1) and deal_time between $2 and $3 order by 1;"

type HistoryServer struct {
	conn *sql.DB
}

func NewHistoryServer(db *sql.DB) *HistoryServer {
	return &HistoryServer{
		conn: db,
	}
}

func (h HistoryServer) GetHistory(ctx context.Context, message *proto.RequestMessage) (*proto.ResponseMessage, error) {
	rows, err := h.conn.Query(query, message.Subject, message.From.AsTime(), message.To.AsTime())
	if err != nil {
		return nil, status.New(codes.Internal, "error query to DB").Err()
	}

	log.Println("Query for", message.Subject, "from", message.From.String(), "to", message.To.String())

	var res []*proto.Pair
	for rows.Next() {
		var temp proto.Pair
		var tempTime time.Time
		err = rows.Scan(&tempTime, &temp.Value)
		temp.Time = timestamppb.New(tempTime)
		if err != nil {
			return nil, status.New(codes.Canceled, "error scan rows from exec").Err()
		}
		res = append(res, &temp)
	}

	return &proto.ResponseMessage{P: res}, status.New(codes.OK, "").Err()
}
