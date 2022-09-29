package db

import (
	"GoIM/pkg/common/constant"
	log2 "GoIM/pkg/common/log"
	"GoIM/pkg/utils"
	"context"
	"strconv"
	"time"

	go_redis "github.com/go-redis/redis/v8"
)

const (
	accountTempCode               = "ACCOUNT_TEMP_CODE"
	resetPwdTempCode              = "RESET_PWD_TEMP_CODE"
	userIncrSeq                   = "REDIS_USER_INCR_SEQ:" // user incr seq
	appleDeviceToken              = "DEVICE_TOKEN"
	userMinSeq                    = "REDIS_USER_MIN_SEQ:"
	uidPidToken                   = "UID_PID_TOKEN_STATUS:"
	conversationReceiveMessageOpt = "CON_RECV_MSG_OPT:"
	getuiToken                    = "GETUI_TOKEN"
	messageCache                  = "MESSAGE_CACHE:"
	SignalCache                   = "SIGNAL_CACHE:"
	SignalListCache               = "SIGNAL_LIST_CACHE:"
	GlobalMsgRecvOpt              = "GLOBAL_MSG_RECV_OPT"
	FcmToken                      = "FCM_TOKEN:"
	groupUserMinSeq               = "GROUP_USER_MIN_SEQ:"
	groupMaxSeq                   = "GROUP_MAX_SEQ:"
	groupMinSeq                   = "GROUP_MIN_SEQ:"
	sendMsgFailedFlag             = "SEND_MSG_FAILED_FLAG:"
	userBadgeUnreadCountSum       = "USER_BADGE_UNREAD_COUNT_SUM:"
)

func (d *DataBases) JudgeAccountEXISTS(account string) (bool, error) {
	key := accountTempCode + account
	n, err := d.RDB.Exists(context.Background(), key).Result()
	if n > 0 {
		return true, err
	} else {
		return false, err
	}
}
func (d *DataBases) SetAccountCode(account string, code, ttl int) (err error) {
	key := accountTempCode + account
	return d.RDB.Set(context.Background(), key, code, time.Duration(ttl)*time.Second).Err()
}
func (d *DataBases) GetAccountCode(account string) (string, error) {
	key := accountTempCode + account
	return d.RDB.Get(context.Background(), key).Result()
}

// Perform seq auto-increment operation of user messages
func (d *DataBases) IncrUserSeq(uid string) (uint64, error) {
	key := userIncrSeq + uid
	seq, err := d.RDB.Incr(context.Background(), key).Result()
	return uint64(seq), err
}

// Get the largest Seq
func (d *DataBases) GetUserMaxSeq(uid string) (uint64, error) {
	key := userIncrSeq + uid
	seq, err := d.RDB.Get(context.Background(), key).Result()
	return uint64(utils.StringToInt(seq)), err
}

// set the largest Seq
func (d *DataBases) SetUserMaxSeq(uid string, maxSeq uint64) error {
	key := userIncrSeq + uid
	return d.RDB.Set(context.Background(), key, maxSeq, 0).Err()
}

// Set the user's minimum seq
func (d *DataBases) SetUserMinSeq(uid string, minSeq uint32) (err error) {
	key := userMinSeq + uid
	return d.RDB.Set(context.Background(), key, minSeq, 0).Err()
}

// Get the smallest Seq
func (d *DataBases) GetUserMinSeq(uid string) (uint64, error) {
	key := userMinSeq + uid
	seq, err := d.RDB.Get(context.Background(), key).Result()
	return uint64(utils.StringToInt(seq)), err
}

func (d *DataBases) SetGroupUserMinSeq(groupID, userID string, minSeq uint64) (err error) {
	key := groupUserMinSeq + "g:" + groupID + "u:" + userID
	return d.RDB.Set(context.Background(), key, minSeq, 0).Err()
}
func (d *DataBases) GetGroupUserMinSeq(groupID, userID string) (uint64, error) {
	key := groupUserMinSeq + "g:" + groupID + "u:" + userID
	seq, err := d.RDB.Get(context.Background(), key).Result()
	return uint64(utils.StringToInt(seq)), err
}

func (d *DataBases) GetGroupMaxSeq(groupID string) (uint64, error) {
	key := groupMaxSeq + groupID
	seq, err := d.RDB.Get(context.Background(), key).Result()
	return uint64(utils.StringToInt(seq)), err
}

func (d *DataBases) IncrGroupMaxSeq(groupID string) (uint64, error) {
	key := groupMaxSeq + groupID
	seq, err := d.RDB.Incr(context.Background(), key).Result()
	return uint64(seq), err
}

func (d *DataBases) SetGroupMaxSeq(groupID string, maxSeq uint64) error {
	key := groupMaxSeq + groupID
	return d.RDB.Set(context.Background(), key, maxSeq, 0).Err()
}

func (d *DataBases) SetGroupMinSeq(groupID string, minSeq uint32) error {
	key := groupMinSeq + groupID
	return d.RDB.Set(context.Background(), key, minSeq, 0).Err()
}

// Store userid and platform class to redis
func (d *DataBases) AddTokenFlag(userID string, platformID int, token string, flag int) error {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	log2.NewDebug("", "add token key is ", key)
	return d.RDB.HSet(context.Background(), key, token, flag).Err()
}

func (d *DataBases) GetTokenMapByUidPid(userID, platformID string) (map[string]int, error) {
	key := uidPidToken + userID + ":" + platformID
	log2.NewDebug("", "get token key is ", key)
	m, err := d.RDB.HGetAll(context.Background(), key).Result()
	mm := make(map[string]int)
	for k, v := range m {
		mm[k] = utils.StringToInt(v)
	}
	return mm, err
}
func (d *DataBases) SetTokenMapByUidPid(userID string, platformID int, m map[string]int) error {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	mm := make(map[string]interface{})
	for k, v := range m {
		mm[k] = v
	}
	return d.RDB.HSet(context.Background(), key, mm).Err()
}
func (d *DataBases) DeleteTokenByUidPid(userID string, platformID int, fields []string) error {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	return d.RDB.HDel(context.Background(), key, fields...).Err()
}
func (d *DataBases) SetSingleConversationRecvMsgOpt(userID, conversationID string, opt int32) error {
	key := conversationReceiveMessageOpt + userID
	return d.RDB.HSet(context.Background(), key, conversationID, opt).Err()
}

func (d *DataBases) GetSingleConversationRecvMsgOpt(userID, conversationID string) (int, error) {
	key := conversationReceiveMessageOpt + userID
	result, err := d.RDB.HGet(context.Background(), key, conversationID).Result()
	return utils.StringToInt(result), err
}
func (d *DataBases) SetUserGlobalMsgRecvOpt(userID string, opt int32) error {
	key := conversationReceiveMessageOpt + userID
	return d.RDB.HSet(context.Background(), key, GlobalMsgRecvOpt, opt).Err()
}
func (d *DataBases) GetUserGlobalMsgRecvOpt(userID string) (int, error) {
	key := conversationReceiveMessageOpt + userID
	result, err := d.RDB.HGet(context.Background(), key, GlobalMsgRecvOpt).Result()
	if err != nil {
		if err == go_redis.Nil {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return utils.StringToInt(result), err
}
func (d *DataBases) CleanUpOneUserAllMsgFromRedis(userID string, operationID string) error {
	ctx := context.Background()
	key := messageCache + userID + "_" + "*"
	vals, err := d.RDB.Keys(ctx, key).Result()
	log2.Debug(operationID, "vals: ", vals)
	if err == go_redis.Nil {
		return nil
	}
	if err != nil {
		return utils.Wrap(err, "")
	}
	for _, v := range vals {
		err = d.RDB.Del(ctx, v).Err()
	}
	return nil
}

func (d *DataBases) DelUserSignalList(userID string) error {
	keyList := SignalListCache + userID
	err := d.RDB.Del(context.Background(), keyList).Err()
	return err
}

func (d *DataBases) SetGetuiToken(token string, expireTime int64) error {
	return d.RDB.Set(context.Background(), getuiToken, token, time.Duration(expireTime)*time.Second).Err()
}

func (d *DataBases) GetGetuiToken() (string, error) {
	result, err := d.RDB.Get(context.Background(), getuiToken).Result()
	return result, err
}

func (d *DataBases) SetSendMsgStatus(status int32, operationID string) error {
	return d.RDB.Set(context.Background(), sendMsgFailedFlag+operationID, status, time.Hour*24).Err()
}

func (d *DataBases) GetSendMsgStatus(operationID string) (int, error) {
	result, err := d.RDB.Get(context.Background(), sendMsgFailedFlag+operationID).Result()
	if err != nil {
		return 0, err
	}
	status, err := strconv.Atoi(result)
	return status, err
}

func (d *DataBases) SetFcmToken(account string, platformid int, fcmToken string, expireTime int64) (err error) {
	key := FcmToken + account + ":" + strconv.Itoa(platformid)
	return d.RDB.Set(context.Background(), key, fcmToken, time.Duration(expireTime)*time.Second).Err()
}

func (d *DataBases) GetFcmToken(account string, platformid int) (string, error) {
	key := FcmToken + account + ":" + strconv.Itoa(platformid)
	return d.RDB.Get(context.Background(), key).Result()
}
func (d *DataBases) IncrUserBadgeUnreadCountSum(uid string) (int, error) {
	key := userBadgeUnreadCountSum + uid
	seq, err := d.RDB.Incr(context.Background(), key).Result()
	return int(seq), err
}
func (d *DataBases) SetUserBadgeUnreadCountSum(uid string, value int) error {
	key := userBadgeUnreadCountSum + uid
	return d.RDB.Set(context.Background(), key, value, 0).Err()
}
func (d *DataBases) GetUserBadgeUnreadCountSum(uid string) (int, error) {
	key := userBadgeUnreadCountSum + uid
	seq, err := d.RDB.Get(context.Background(), key).Result()
	return utils.StringToInt(seq), err
}
