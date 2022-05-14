package main

import (
	"math/rand"
	"fmt"
)

// This file is for the mining code.
// Note that "targetBits" for this assignment, at least initially, is 33.
// This could change during the assignment duration!  I will post if it does.

var (
	characterSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func (self *Block) AlterNonce() {
	currentNonce := self.Nonce
	
	// Random replace one char into another char from char set.
	// Are there more efficient ways to do
	nonceBytes := []byte(currentNonce)

	randomReplaceIndex := rand.Intn(len(nonceBytes))
	randomNewCharIndexInSet := rand.Intn(len(characterSet))

	nonceBytes[randomReplaceIndex] = characterSet[randomNewCharIndexInSet]

	self.Nonce = string(nonceBytes)
}

// Mine mines a block by varying the nonce until the hash has targetBits 0s in
// the beginning.  Could take forever if targetBits is too high.
// Modifies a block in place by using a pointer receiver.
func (self *Block) Mine(targetBits uint8) {
	// your mining code here
	// also feel free to get rid of this method entirely if you want to
	// organize things a different way; this is just a suggestion
	trailCount := 0

	for ! CheckWork(*self, targetBits) {
		self.AlterNonce()
		trailCount += 1

		if trailCount % 1000000000 == 0 {
			fmt.Println(trailCount)
		}
	}

	fmt.Println("Total number of trials: ", trailCount)
	return
}

// // CheckWork checks if there's enough work
// func CheckWork(bl Block, targetBits uint8) bool {
// 	// your checkwork code here
// 	// feel free to inline this or do something else.  I just did it this way
// 	// so I'm giving empty functions here.
// 	h := bl.Hash()

// 	for i := uint8(0); i < targetBits; i ++ {
// 		// Check work fails when meet 1.
// 		if (h[i / 8] >> (7 - (i % 8))) & 0x01 ==1 {
// 			return false
// 		}
// 	}
	
// 	return true
// }


// CheckWork checks if there's enough work
// Is this a better implementation?
func CheckWork(bl Block, targetBits uint8) bool {
	// your checkwork code here
	// feel free to inline this or do something else.  I just did it this way
	// so I'm giving empty functions here.
	h := bl.Hash()

	var numZeroBytes int = int(targetBits / 8)
	var bitsToCheck  uint8 = targetBits % 8
	// Check chunk of bytes to save time.
	for i := 0; i < numZeroBytes; i ++ {
		if h[i] != 0 {
			return false
		}
	}

	if bitsToCheck == 0 {
		return true
	}

	// Check remaing bits if any.
	for i := uint8(0); i < bitsToCheck; i ++ {
		// Check work fails when meet 1.
		if (h[numZeroBytes] >> (7 - (i % 8))) & 0x01 ==1 {
			return false
		}
	}
	
	return true
}