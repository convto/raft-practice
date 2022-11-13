// raftpractice is the implementation of [raft](https://raft.github.io/raft.pdf) consensus algolithm.
// The purpose of this project is to understand it better by doing a simple implementation of raft.
// This project is a simple on-memory implementation.
package raftpractice

type server struct {
	*persistentState // This is a state that should be persisted. But this project is for learning purposes. For simplicity, the state is implemented on-memory.
	*volatileState
	*volatileStateOnLeader
}

type persistentState struct {
	currentTerm uint64 // latest term server has seen (initialized to 0 on first boot, increases monotonically)
	votedFor    uint64 // candidateId that received vote in current term (or null if none)
	logs        []logEntry
}

type logEntry struct {
	command string // command for state machine
	term    uint64 // term when entry was received by leader
}

type volatileState struct {
	commitIndex int // index of highest log entry known to be committed (initialized to 0, increases monotonically)
	lastApplied int // index of highest log entry applied to state machine (initialized to 0, increases monotonically)
}

type volatileStateOnLeader struct {
	nextIndex  []int // index of the next log entry to send to each server (initialized to leader last log index + 1)
	matchIndex []int // index of highest log entry known to be replicated on server (initialized to 0, increases monotonically)
}
