package raft

// NodeState defines the interface for a Raft state
type NodeState interface {
	RequestVote(args *RequestVoteArgs, reply *RequestVoteReply)

	AppendEntries(args *AppendEntriesReply, reply *AppendEntriesReply)

	HeartBeatHandler()
}

// FollowerState contains the Raft follower state implementation
type FollowerState struct {
	raftNode *Raft
}

func (state *FollowerState) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {

}

func (state *FollowerState) AppendEntries(args *AppendEntriesReply, reply *AppendEntriesReply) {

}

func (state *FollowerState) HeartBeatHandler() {}

// CandidateState contains the Raft candidate state implementation
type CandidateState struct {
	raftNode *Raft
}

func (state *CandidateState) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {

}

func (state *CandidateState) AppendEntries(args *AppendEntriesReply, reply *AppendEntriesReply) {

}

func (state *CandidateState) HeartBeatHandler() {}

// LeaderState contains the Raft leader state implementation
type LeaderState struct {
	raftNode *Raft
}

func (state *LeaderState) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {

}

func (state *LeaderState) AppendEntries(args *AppendEntriesReply, reply *AppendEntriesReply) {

}

func (state *LeaderState) HeartBeatHandler() {}
