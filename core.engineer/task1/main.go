package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"io"
	"log"
	"sort"
	"strings"
)

type TLV struct {
	Tag    byte
	Length byte
	Value  []byte
}

func ReadTLV(r io.Reader) (*TLV, error) {
	var record TLV

	err := binary.Read(r, binary.BigEndian, &record.Tag)
	if err == io.EOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil,
			fmt.Errorf("error reading TLV: %s", err)
	}
	err = binary.Read(r, binary.BigEndian, &record.Length)
	if err != nil {
		fmt.Errorf("error reading TLV: %s", err)
	}
	record.Value = make([]byte, record.Length)
	_, err = io.ReadFull(r, record.Value)
	if err != nil {
		return nil, fmt.Errorf("error reading TLV: %s", err)
	}

	return &record, nil
}

func main() {
	var example = "0812561ae6bd42cfb7bee1311a29893508e71c340912e4614f3d9e6c61dccdd5af228051eb4373170A058394b17ac301086ef72ac5367180eb051040d4241e426bbde12bca805f665227e8060234ec07042859ffbc0202b4ed0304bdadb4e00402de60"
	rawMessage, err := hex.DecodeString(example)
	if err != nil {
		fmt.Println("couldn't decode message:", err.Error())
		return
	}
	message := bytes.NewBuffer(rawMessage)

	tlvs, err := parseTLVs(message)
	if err != nil {
		log.Fatal(err)
	}

	// sort by tag
	sort.Slice(tlvs, func(i, j int) bool {
		return tlvs[i].Tag < tlvs[j].Tag
	})

	// print the tlvs sorted by tag
	for i, tlv := range tlvs {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%#x", tlv.Tag)
	}
	fmt.Println()

	//1. Extract and order the tag values for the following tag ids by ascending numeric tag id.
	wantedTagValuesResult := findWantedTagValues(tlvs, []byte{0x01, 0x03, 0x04, 0x05, 0x06, 0x0a})
	//2. Create an uppercase string with a full stop between each tag value.
	upcase := strings.ToUpper(getPeriodSeparatedStringForTagValues(wantedTagValuesResult))
	//3. Convert the string to a byte array.
	upcaseByteSlice := []byte(upcase)
	//4. Hash the byte array using the Blake2b algorithm.
	// NOTE it wasn't obvious which sum to use from the readme
	hash := blake2b.Sum256(upcaseByteSlice)
	//5. Base64url encode the resulting hash. Starting from the above example you should get the following output:
	result := base64.RawURLEncoding.EncodeToString(hash[:])
	fmt.Println(result)
}

func getPeriodSeparatedStringForTagValues(tlvs []*TLV) string {
	var sb []string
	for _, tlv := range tlvs {
		// NOTE this bit wasn't obvious from the readme
		sb = append(sb, hex.EncodeToString(tlv.Value))
	}
	return strings.Join(sb, ".")
}

func findWantedTagValues(tlvs []*TLV, wantedTagValues []byte) []*TLV {
	var wantedTagValuesResult []*TLV

	for _, tlv := range tlvs {
		if bytes.Contains(wantedTagValues, []byte{tlv.Tag}) {
			wantedTagValuesResult = append(wantedTagValuesResult, tlv)
		}
	}
	return wantedTagValuesResult
}

func parseTLVs(message *bytes.Buffer) ([]*TLV, error) {
	var result []*TLV
	for {
		tlv, err := ReadTLV(message)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		result = append(result, tlv)
	}
	return result, nil
}
