package pool

import "errors"

type op string

const (
	Hash   op = "encrypt"
	Cmpare    = "decrypt"
)

type WorkRequest struct {
	OP      op
	Text    []byte
	Compare []byte
}

type WorkResponse struct {
	Wr      WorkRequest
	Result  []byte
	Matched bool
	Err     error
}

func Process(wr WorkRequest) WorkResponse {
	switch wr.OP {

	case Hash:
		return hashWork(wr)
	case Cmpare:
		return compareWork(wr)

	default:
		return WorkResponse{Err: errors.New("unsupported operation")}
	}
}
