// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	roomTblFieldNames          = builder.RawFieldNames(&RoomTbl{})
	roomTblRows                = strings.Join(roomTblFieldNames, ",")
	roomTblRowsExpectAutoSet   = strings.Join(stringx.Remove(roomTblFieldNames), ",")
	roomTblRowsWithPlaceHolder = strings.Join(stringx.Remove(roomTblFieldNames, "`id`"), "=?,") + "=?"
)

type (
	roomTblModel interface {
		Insert(ctx context.Context, data *RoomTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RoomTbl, error)
		Update(ctx context.Context, data *RoomTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRoomTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RoomTbl struct {
		Id       int64          `db:"id"`
		HouseId  sql.NullInt64  `db:"house_id"`
		Name     sql.NullString `db:"name"`
		Status   int64          `db:"status"`
		Capacity sql.NullInt64  `db:"capacity"`
		EIndex   sql.NullInt64  `db:"e_index"`
		WIndex   sql.NullInt64  `db:"w_index"`
	}
)

func newRoomTblModel(conn sqlx.SqlConn) *defaultRoomTblModel {
	return &defaultRoomTblModel{
		conn:  conn,
		table: "`room_tbl`",
	}
}

func (m *defaultRoomTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRoomTblModel) FindOne(ctx context.Context, id int64) (*RoomTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roomTblRows, m.table)
	var resp RoomTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoomTblModel) Insert(ctx context.Context, data *RoomTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, roomTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.HouseId, data.Name, data.Status, data.Capacity, data.EIndex, data.WIndex)
	return ret, err
}

func (m *defaultRoomTblModel) Update(ctx context.Context, data *RoomTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, roomTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.HouseId, data.Name, data.Status, data.Capacity, data.EIndex, data.WIndex, data.Id)
	return err
}

func (m *defaultRoomTblModel) tableName() string {
	return m.table
}
