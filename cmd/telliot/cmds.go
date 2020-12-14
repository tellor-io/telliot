package main

type configPath string
type logConfigPath string
type logLevel string
type dataserverCmd string

type mineCmd struct {
	remote bool
}

type newDisputeCmd struct {
	requestID  string `arg required `
	minerIndex string `arg required `
	timestamp  string `arg required `
}

type voteCmd struct {
	disputeId string `arg required`
	support   bool   `arg required`
}

type stakeCmd struct {
	Operation string `arg required`
}

type balanceCmd struct {
	Address string `arg optional`
}

type transferCmd tokenCmd
type approveCmd tokenCmd

type tokenCmd struct {
	Address string `arg required`
	Amount  string `arg required`
}
