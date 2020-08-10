package db

//DataServerProxy interface for local interaction/abstraction/testing
type DataServerProxy interface {
	//RequestSigner
	//RequestValidator
	//local call to get a data server value by its key
	Get(key string) ([]byte, error)

	//local call to put data into the data server's store. All keys should
	//be prefixed with the calling miner's public ETH key to avoid conflicts.
	//Implementation must ensure thread safety from multiple miners attempting
	//to write at the same time.
	Put(key string, value []byte) (map[string][]byte, error)

	//put multiple keys and values on remote data server
	BatchPut(keys []string, values [][]byte) (map[string][]byte, error)

	//local call to get several data server values by their keys
	BatchGet(keys []string) (map[string][]byte, error)

	//notification that a remote miner has requested data
	IncomingRequest(data []byte) ([]byte, error)
}
