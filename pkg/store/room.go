package store

import (
	"errors"

	roomv1 "github.com/pilab-dev/ion/pkg/gen/pb/room/v1"
)

// ErrRoomNotFound is returned when a room is not found.
var ErrRoomNotFound = errors.New("room not found")

// NewRoom creates a new room. The sid is the room id, the name is the name of the room,
// the password is the password of the room, the description is the description of the room,
// and the lock is the lock status of the room. By default, the max peers is 0, which means
// there is no limit to the number of peers in the room.
func NewRoom(sid, name, password, description string, lock bool) *Room {
	return &Room{
		SID:         sid,
		Name:        name,
		Password:    password,
		Description: description,
		Lock:        lock,
		MaxPeers:    0,
	}
}

// Room is a room in the store. It contains information about the room.
type Room struct {
	SID         string
	Name        string
	Password    string
	Description string
	Lock        bool
	MaxPeers    uint32
}

// ToProto converts a room to a proto room.
func (r *Room) ToProto() *roomv1.Room {
	return &roomv1.Room{
		Sid:         r.SID,
		Name:        r.Name,
		Lock:        r.Lock,
		Password:    r.Password,
		Description: r.Description,
		MaxPeers:    r.MaxPeers,
	}
}

// Peers is a list of peers.
type Peers []*Peer

// ToProto converts a list of peers to a list of proto peers.
func (p Peers) ToProto() []*roomv1.Peer {
	peers := make([]*roomv1.Peer, len(p))
	for i, peer := range p {
		peers[i] = peer.ToProto()
	}

	return peers
}

// Peer is a peer in a room. It contains information about the peer.
// The SID is the room id, the UID is the peer id, the DisplayName is the
// name of the peer, the ExtraInfo is extra information about the peer, the
// Role is the role of the peer, the Protocol is the protocol used by the peer,
// the Avatar is the avatar of the peer, the Direction is the direction of the
// peer, the Vendor is the vendor of the peer, and the Destination is the
// destination of the peer. The Role, Protocol, and Direction are enums.
type Peer struct {
	SID         string
	UID         string
	DisplayName string
	ExtraInfo   []byte
	Role        string
	Protocol    string
	Avatar      string
	Direction   string
	Vendor      string
	Destination string
}

// ToProto converts a peer to a proto peer.
func (p *Peer) ToProto() *roomv1.Peer {
	return &roomv1.Peer{
		Sid:         p.SID,
		Uid:         p.UID,
		DisplayName: p.DisplayName,
		ExtraInfo:   p.ExtraInfo,
		Role:        roomv1.Role(roomv1.Role_value[p.Role]),
		Protocol:    roomv1.Protocol(roomv1.Protocol_value[p.Protocol]),
		Avatar:      p.Avatar,
		Direction:   roomv1.Peer_Direction(roomv1.Peer_Direction_value[p.Direction]),
		Vendor:      p.Vendor,
		Destination: p.Destination,
	}
}

// RoomStore is an interface for storing rooms. It is used by the RoomManager
// to store and retrieve rooms, and to add and remove peers from rooms.
type RoomStore interface {
	// SaveRoom saves a room to the store.
	SaveRoom(room *Room) error

	// RemoveRoom removes a room from the store.
	RemoveRoom(sid string) error

	// RemovePeer removes a peer from the room. If the room is empty, it will be removed.
	// The sid is the room id and the uid is the peer id.
	RemovePeer(sid, uid string) error

	// GetRoom gets a room from the store. If the room does not exist,
	// it will return an error.
	GetRoom(sid string) (*Room, error)

	// GetRoomPeers gets all peers in the room.
	GetRoomPeers(sid string) (Peers, error)

	// AddPeer adds a peer to the room.
	// The sid is the room id and the p is the peer.
	AddPeer(sid string, p *Peer) error

	// IsPeerExists checks if a peer exists in the room.
	IsPeerExists(sid, uid string) bool
}
