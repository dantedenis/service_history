package service

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"service_history/internal/app/proto"
)

var query = "select deal_time, coast from history_deal where pair_id=(select id from currency_pair where pair=$1) and deal_time between $2 and $3"

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
		err = rows.Scan(&temp.Time, &temp.Value)
		if err != nil {
			return nil, status.New(codes.Canceled, "error scan rows from exec").Err()
		}
		res = append(res, &temp)
	}

	return &proto.ResponseMessage{P: res}, status.New(codes.OK, "").Err()
}
