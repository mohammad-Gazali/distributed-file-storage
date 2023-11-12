package peertopeer


// HandeshakeFunc .... TODO
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }