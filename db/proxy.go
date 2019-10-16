package db

//DataServerProxy interface for local interaction/abstraction/testing
type DataServerProxy interface {
	//RequestSigner
	//RequestValidator
	//local call to get a data server value by its key
	Get(key string) ([]byte, error)

	//local call to get several data server values by their keys
	BatchGet(keys []string) (map[string][]byte, error)

	//notification that a remote miner has requested data
	IncomingRequest(data []byte) ([]byte, error)
}
