// +build go1.14

package utils

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestComplexClassDecryption tests the complex struct decryption.
// Decrypted string should match Bc846Ri5HK1ixqP/dzAyZq23Z/NBlcPn2UX8h38xTGINs72yF5gtU0t9fFEMxjY+DmezWt0nG7eN7RABrj697tK1nooVHYIxgDLMsjMTw5N0K+rUM823n7LcHfEoXaX8oH2E6zkg6iK5pmT8nlh6LF6Bw1G5zkluT8oTjnbFJcpEvTyT2ZKzcqptgYsE9XZiqgBMEfYqwphDzmOv+TjHkJai+paV0rzFxIfVK8KHCA14z+1kKDMPghlmzx2tUmmbQb04hjhvgDvvi3tknytYVqJo1L5jZkAZTVXRfed7wq+L+1V824c9AwVsG9iCv15/Jemjjfzk07MXawk+hjmQvjQDWLS/ww3vwkNXiuJITbVCPOBADwJhBnFqkkb/Hd8LaKwyFhWeXwoZWbqugDoYufUzJApf4Nl/4RthYoisqJIokmxiWvYeD1TuH+C457kDaEu3aJd+KdLf8k9QkmaDNqkZo9Z/BRkZ63oMna1aEBy7bSE3l/lw40dnhsMaYfYk
func TestComplexClassDecryption(t *testing.T) {
	assert := assert.New(t)

	message := "Bc846Ri5HK1ixqP/dzAyZq23Z/NBlcPn2UX8h38xTGINs72yF5gtU0t9fFEMxjY+DmezWt0nG7eN7RABrj697tK1nooVHYIxgDLMsjMTw5N0K+rUM823n7LcHfEoXaX8oH2E6zkg6iK5pmT8nlh6LF6Bw1G5zkluT8oTjnbFJcpEvTyT2ZKzcqptgYsE9XZiRs1+hbUOxVQW8p7EIsErOmzEB8OKrFemADTysTDKflPcOoAyCLLEB1+uV/DMpXdUUXW0aXpU1/PZt3ggVMJI9AYG/lEzbXtqLZKtaMUcHOrADP9TH2ePxn6OHzNYvgJCs0KF/HAfIiv4tl6sx2jP/PBQ+IIC3/R5WBLMLcIgC7igWyLeAmltKFvda2c5MUqnYzqYiJIAdK6SoQWoaV6lBl0MMKGXG6UB/iB4YFQQN22qdfQ0a9RdYVRDb5iUap2aqhuAAmQZnCJxGbptuyn2MV1Y6fczSstwCUrlXQE5+1E="

	decrypted, decErr := DecryptString("enigma", message)
	assert.NoError(decErr)

	customComplexMessage := initComplexMessage()
	b, err := json.Marshal(customComplexMessage)

	assert.NoError(err)
	assert.Equal(string(b), decrypted)
}

// TestComplexClassEncryption tests the complex struct encryption.
// Encrypted string should match Bc846Ri5HK1ixqP/dzAyZq23Z/NBlcPn2UX8h38xTGINs72yF5gtU0t9fFEMxjY+DmezWt0nG7eN7RABrj697tK1nooVHYIxgDLMsjMTw5N0K+rUM823n7LcHfEoXaX8oH2E6zkg6iK5pmT8nlh6LF6Bw1G5zkluT8oTjnbFJcpEvTyT2ZKzcqptgYsE9XZiqgBMEfYqwphDzmOv+TjHkJai+paV0rzFxIfVK8KHCA14z+1kKDMPghlmzx2tUmmbQb04hjhvgDvvi3tknytYVqJo1L5jZkAZTVXRfed7wq+L+1V824c9AwVsG9iCv15/Jemjjfzk07MXawk+hjmQvjQDWLS/ww3vwkNXiuJITbVCPOBADwJhBnFqkkb/Hd8LaKwyFhWeXwoZWbqugDoYufUzJApf4Nl/4RthYoisqJIokmxiWvYeD1TuH+C457kDaEu3aJd+KdLf8k9QkmaDNqkZo9Z/BRkZ63oMna1aEBy7bSE3l/lw40dnhsMaYfYk
func TestComplexClassEncryption(t *testing.T) {
	assert := assert.New(t)

	customComplexMessage := initComplexMessage()

	b1, err := json.Marshal(customComplexMessage)
	assert.NoError(err)

	encrypted := EncryptString("enigma", string(b1))
	assert.Equal("Bc846Ri5HK1ixqP/dzAyZq23Z/NBlcPn2UX8h38xTGINs72yF5gtU0t9fFEMxjY+DmezWt0nG7eN7RABrj697tK1nooVHYIxgDLMsjMTw5N0K+rUM823n7LcHfEoXaX8oH2E6zkg6iK5pmT8nlh6LF6Bw1G5zkluT8oTjnbFJcpEvTyT2ZKzcqptgYsE9XZiRs1+hbUOxVQW8p7EIsErOmzEB8OKrFemADTysTDKflPcOoAyCLLEB1+uV/DMpXdUUXW0aXpU1/PZt3ggVMJI9AYG/lEzbXtqLZKtaMUcHOrADP9TH2ePxn6OHzNYvgJCs0KF/HAfIiv4tl6sx2jP/PBQ+IIC3/R5WBLMLcIgC7igWyLeAmltKFvda2c5MUqnYzqYiJIAdK6SoQWoaV6lBl0MMKGXG6UB/iB4YFQQN22qdfQ0a9RdYVRDb5iUap2aqhuAAmQZnCJxGbptuyn2MV1Y6fczSstwCUrlXQE5+1E=", encrypted)
}
