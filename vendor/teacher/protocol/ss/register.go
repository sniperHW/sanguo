package ss

import (
	"github.com/sniperHW/sanguo/codec/pb"
	"github.com/sniperHW/sanguo/protocol/ss/rpc"
	"github.com/sniperHW/sanguo/protocol/ss/ssmessage"
)

func init() {
	//普通消息
	pb.Register("ss", &ssmessage.Echo{}, 1001)
	pb.Register("ss", &ssmessage.KickGameUser{}, 1002)
	pb.Register("ss", &ssmessage.KickGateUser{}, 1003)
	pb.Register("ss", &ssmessage.GateUserDisconnect{}, 1004)
	pb.Register("ss", &ssmessage.SsToGate{}, 1011)
	pb.Register("ss", &ssmessage.ReportGate{}, 1012)
	pb.Register("ss", &ssmessage.SsToGateError{}, 1013)
	pb.Register("ss", &ssmessage.MapHeartbeat{}, 1014)
	pb.Register("ss", &ssmessage.UserLoginToDir{}, 1021)
	pb.Register("ss", &ssmessage.Relay{}, 1022)
	pb.Register("ss", &ssmessage.ReportStatus{}, 1023)
	pb.Register("ss", &ssmessage.StartAoi{}, 1024)
	pb.Register("ss", &ssmessage.MapToWorld{}, 1031)
	pb.Register("ss", &ssmessage.WorldBroadcastToMap{}, 1032)
	pb.Register("ss", &ssmessage.FunctionSwitchReload{}, 1033)
	pb.Register("ss", &ssmessage.FirewallUpdate{}, 1041)
	pb.Register("ss", &ssmessage.TimeUpdate{}, 1042)
	pb.Register("ss", &ssmessage.RankEnd{}, 1043)
	pb.Register("ss", &ssmessage.ScarsIngrainUpdate{}, 1044)

	//rpc请求
	pb.Register("rpc_req", &rpc.EchoReq{}, 1001)
	pb.Register("rpc_req", &rpc.SynctokenReq{}, 1002)
	pb.Register("rpc_req", &rpc.GateUserLoginReq{}, 1003)
	pb.Register("rpc_req", &rpc.GateUserReconnectReq{}, 1004)
	pb.Register("rpc_req", &rpc.ForwardUserMsgReq{}, 1005)
	pb.Register("rpc_req", &rpc.EnterWorldReq{}, 1011)
	pb.Register("rpc_req", &rpc.EnterMapReq{}, 1012)
	pb.Register("rpc_req", &rpc.LeaveMapReq{}, 1013)
	pb.Register("rpc_req", &rpc.MoveReq{}, 1014)
	pb.Register("rpc_req", &rpc.WorldObjPushReq{}, 1021)
	pb.Register("rpc_req", &rpc.WorldObjDoReq{}, 1022)
	pb.Register("rpc_req", &rpc.GetLogicTimeReq{}, 1031)
	pb.Register("rpc_req", &rpc.GetNextMailIDReq{}, 1040)
	pb.Register("rpc_req", &rpc.SendMailReq{}, 1041)
	pb.Register("rpc_req", &rpc.DelMailReq{}, 1042)
	pb.Register("rpc_req", &rpc.UpdateMailReq{}, 1043)
	pb.Register("rpc_req", &rpc.GetMailListReq{}, 1044)
	pb.Register("rpc_req", &rpc.GetMailByIDReq{}, 1045)
	pb.Register("rpc_req", &rpc.TeamCreateReq{}, 1051)
	pb.Register("rpc_req", &rpc.TeamGetNearTeamReq{}, 1052)
	pb.Register("rpc_req", &rpc.TeamGetNearPlayerReq{}, 1053)
	pb.Register("rpc_req", &rpc.TeamPlayerLeaveReq{}, 1054)
	pb.Register("rpc_req", &rpc.TeamKickPlayerReq{}, 1055)
	pb.Register("rpc_req", &rpc.TeamJoinApplyReq{}, 1056)
	pb.Register("rpc_req", &rpc.TeamJoinReplyReq{}, 1057)
	pb.Register("rpc_req", &rpc.TeamPlayerGetFromGameReq{}, 1058)
	pb.Register("rpc_req", &rpc.TeamDismissReq{}, 1059)
	pb.Register("rpc_req", &rpc.RankGetTopListReq{}, 1071)
	pb.Register("rpc_req", &rpc.RankSetScoreReq{}, 1072)
	pb.Register("rpc_req", &rpc.RankCheckIDReq{}, 1073)
	pb.Register("rpc_req", &rpc.RankCreateReq{}, 1074)
	pb.Register("rpc_req", &rpc.RankDeleteScoreReq{}, 1075)
	pb.Register("rpc_req", &rpc.RankGetRankReq{}, 1076)
	pb.Register("rpc_req", &rpc.ConflictZoneReqGroupReq{}, 1081)

	//rpc响应
	pb.Register("rpc_resp", &rpc.EchoResp{}, 1001)
	pb.Register("rpc_resp", &rpc.SynctokenResp{}, 1002)
	pb.Register("rpc_resp", &rpc.GateUserLoginResp{}, 1003)
	pb.Register("rpc_resp", &rpc.GateUserReconnectResp{}, 1004)
	pb.Register("rpc_resp", &rpc.ForwardUserMsgResp{}, 1005)
	pb.Register("rpc_resp", &rpc.EnterWorldResp{}, 1011)
	pb.Register("rpc_resp", &rpc.EnterMapResp{}, 1012)
	pb.Register("rpc_resp", &rpc.LeaveMapResp{}, 1013)
	pb.Register("rpc_resp", &rpc.MoveResp{}, 1014)
	pb.Register("rpc_resp", &rpc.WorldObjPushResp{}, 1021)
	pb.Register("rpc_resp", &rpc.WorldObjDoResp{}, 1022)
	pb.Register("rpc_resp", &rpc.GetLogicTimeResp{}, 1031)
	pb.Register("rpc_resp", &rpc.GetNextMailIDResp{}, 1040)
	pb.Register("rpc_resp", &rpc.SendMailResp{}, 1041)
	pb.Register("rpc_resp", &rpc.DelMailResp{}, 1042)
	pb.Register("rpc_resp", &rpc.UpdateMailResp{}, 1043)
	pb.Register("rpc_resp", &rpc.GetMailListResp{}, 1044)
	pb.Register("rpc_resp", &rpc.GetMailByIDResp{}, 1045)
	pb.Register("rpc_resp", &rpc.TeamCreateResp{}, 1051)
	pb.Register("rpc_resp", &rpc.TeamGetNearTeamResp{}, 1052)
	pb.Register("rpc_resp", &rpc.TeamGetNearPlayerResp{}, 1053)
	pb.Register("rpc_resp", &rpc.TeamPlayerLeaveResp{}, 1054)
	pb.Register("rpc_resp", &rpc.TeamKickPlayerResp{}, 1055)
	pb.Register("rpc_resp", &rpc.TeamJoinApplyResp{}, 1056)
	pb.Register("rpc_resp", &rpc.TeamJoinReplyResp{}, 1057)
	pb.Register("rpc_resp", &rpc.TeamPlayerGetFromGameResp{}, 1058)
	pb.Register("rpc_resp", &rpc.TeamDismissResp{}, 1059)
	pb.Register("rpc_resp", &rpc.RankGetTopListResp{}, 1071)
	pb.Register("rpc_resp", &rpc.RankSetScoreResp{}, 1072)
	pb.Register("rpc_resp", &rpc.RankCheckIDResp{}, 1073)
	pb.Register("rpc_resp", &rpc.RankCreateResp{}, 1074)
	pb.Register("rpc_resp", &rpc.RankDeleteScoreResp{}, 1075)
	pb.Register("rpc_resp", &rpc.RankGetRankResp{}, 1076)
	pb.Register("rpc_resp", &rpc.ConflictZoneReqGroupResp{}, 1081)

}
