package x16r

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestX16R(t *testing.T) {
	data, _ := hex.DecodeString("000000302aefb1655612597af8b82943fbe79efe672d44a226ae9a548aab00000000000002aeda49909cbedc98b1101bce8beae89a342ddf2946705a528fb1bc95e18bfce7eae15bfbc6001b43f2c36a")
	hash, _ := hex.DecodeString("01b800b258ed908d8ed4dffd87b87105e0b6e11a5e7f465b741d000000000000")
	res := Sum(data)
	if !bytes.Equal(res, hash) {
		t.Error("hash", hex.EncodeToString(res))
	}
}
