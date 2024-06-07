package friend

import (
	"github.com/openimsdk/open-im-server/v3/pkg/common/storage/model"
	"github.com/openimsdk/protocol/relation"
	"github.com/openimsdk/tools/utils/datautil"
)

func friendDB2PB(db *model.Friend) *relation.FriendInfo {
	return &relation.FriendInfo{
		OwnerUserID:    db.OwnerUserID,
		FriendUserID:   db.FriendUserID,
		FriendNickname: db.FriendNickname,
		FriendFaceURL:  db.FriendFaceURL,
		Remark:         db.Remark,
		CreateTime:     db.CreateTime.UnixMilli(),
		AddSource:      db.AddSource,
		OperatorUserID: db.OperatorUserID,
		Ex:             db.Ex,
		IsPinned:       db.IsPinned,
	}
}

func friendsDB2PB(db []*model.Friend) []*relation.FriendInfo {
	return datautil.Slice(db, friendDB2PB)
}

func blackDB2PB(db *model.Black) *relation.BlackInfo {
	return &relation.BlackInfo{
		OwnerUserID:    db.OwnerUserID,
		BlackUserID:    db.BlackUserID,
		BlackNickname:  db.BlackNickname,
		BlackFaceURL:   db.BlackFaceURL,
		CreateTime:     db.CreateTime.UnixMilli(),
		AddSource:      db.AddSource,
		OperatorUserID: db.OperatorUserID,
		Ex:             db.Ex,
	}
}

func blacksDB2PB(db []*model.Black) []*relation.BlackInfo {
	return datautil.Slice(db, blackDB2PB)
}
