package go_tiny_cache

import pb "go_tiny_cache/tinycachepb"

// 根据传入的key值选择相应的节点，理解为客户端？
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// Get方法用于从对应的group查找缓存值。
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
