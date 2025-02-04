package sync

import (
	"encoding/binary"
	"sync"
	"time"
)

type ObjSync struct {
	mu           sync.Mutex
	prvTimestamp int64
	instanceId   int
}

func NewObjSync(instanceId int) *ObjSync {
	return &ObjSync{
		mu:           sync.Mutex{},
		prvTimestamp: 0,
		instanceId:   instanceId,
	}
}

// GenServiceObjID generates a unique ID for a service object.
// The ID is composed of a timestamp, instance ID, and object type.
// The timestamp is in milliseconds since the Unix epoch.
// The instance ID is a 7-bit integer representing the instance of the service.
// The object type is a 5-bit integer representing the type of the object.
// The function ensures that the generated ID is unique by using a mutex lock.
//
// Parameters:
// objId (int): The ID of the object.
//
// Returns:
// int64: The generated unique ID for the service object.
func (oSync *ObjSync) GenServiceObjID() int64 {
	var ret int64 = 0

	binsID := make([]byte, 8)
	baseB := make([]byte, 8)
	instanceB := make([]byte, 4)
	objB := make([]byte, 4)

	var instanceMod = oSync.instanceId % 128 // max 128 instance
	var objMod = 1 % 32                      // max 32 type; 11111 -> type other
	oSync.mu.Lock()
	defer oSync.mu.Unlock()

	t := time.Now().UnixMilli()
	if t <= oSync.prvTimestamp {
		ret = oSync.prvTimestamp + 1
	} else {
		ret = t
	}
	oSync.prvTimestamp = ret

	binary.BigEndian.PutUint64(baseB, uint64(ret))
	binary.BigEndian.PutUint32(instanceB, uint32(instanceMod))
	binary.BigEndian.PutUint32(objB, uint32(objMod))

	// set last 7 bit
	binsID[7] = instanceB[3]
	// set next 41 bit time
	binsID[7] = baseB[7]<<7 | binsID[7]
	binsID[6] = baseB[7]>>1 | baseB[6]<<7
	binsID[5] = baseB[6]>>1 | baseB[5]<<7
	binsID[4] = baseB[5]>>1 | baseB[4]<<7
	binsID[3] = baseB[4]>>1 | baseB[3]<<7
	binsID[2] = baseB[3]>>1 | baseB[2]<<7
	// set 5 byte obj type
	binsID[1] = objB[3]

	ret = int64(binary.BigEndian.Uint64(binsID))

	return ret
}
