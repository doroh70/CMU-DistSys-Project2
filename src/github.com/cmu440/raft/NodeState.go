package raft

// NodeState defines the interface for a Raft state
type NodeState interface {
	NormalOperation()

	RequestVote(args *RequestVoteArgs, reply *RequestVoteReply)

	AppendEntries(args *AppendEntriesReply, reply *AppendEntriesReply)
}

// FollowerState contains the Raft follower state implementation
type FollowerState struct {
	raftNode *Raft
}

func (state *FollowerState) NormalOperation() {

}

func (state *FollowerState) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {

}

func (state *FollowerState) AppendEntries(args *AppendEntriesReply, reply *AppendEntriesReply) {

}

// CandidateState contains the Raft candidate state implementation
type CandidateState struct {
	raftNode *Raft
}

func (state *CandidateState) NormalOperation() {

}

func (state *CandidateState) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {

}

func (state *CandidateState) AppendEntries(args *AppendEntriesReply, reply *AppendEntriesReply) {

}

// LeaderState contains the Raft leader state implementation
type LeaderState struct {
	raftNode *Raft
}

func (state *LeaderState) NormalOperation() {

}

func (state *LeaderState) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {

}

func (state *LeaderState) AppendEntries(args *AppendEntriesReply, reply *AppendEntriesReply) {

}
