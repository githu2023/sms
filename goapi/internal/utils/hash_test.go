package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcryptHash(t *testing.T) {
	password := "testpassword123"
	hash := BcryptHash(password)

	// Hash应该不为空
	assert.NotEmpty(t, hash)
	// Hash应该与原密码不同
	assert.NotEqual(t, password, hash)
	// Hash长度应该是bcrypt标准长度（60字符）
	assert.Equal(t, 60, len(hash))
	// Hash应该以$2a$开头（bcrypt标准格式）
	assert.Contains(t, hash, "$2a$")
}

func TestBcryptCheck(t *testing.T) {
	password := "testpassword123"
	hash := BcryptHash(password)

	// 正确密码应该验证成功
	assert.True(t, BcryptCheck(password, hash))

	// 错误密码应该验证失败
	assert.False(t, BcryptCheck("wrongpassword", hash))

	// 空密码应该验证失败
	assert.False(t, BcryptCheck("", hash))

	// 空hash应该验证失败
	assert.False(t, BcryptCheck(password, ""))
}

func TestBcryptConsistency(t *testing.T) {
	password := "consistency_test_password"

	// 多次加密同一密码应该产生不同的hash（因为salt不同）
	hash1 := BcryptHash(password)
	hash2 := BcryptHash(password)
	assert.NotEqual(t, hash1, hash2)

	// 但验证时都应该成功
	assert.True(t, BcryptCheck(password, hash1))
	assert.True(t, BcryptCheck(password, hash2))
}

func TestMD5V(t *testing.T) {
	input := []byte("test string")
	hash := MD5V(input)

	// MD5 hash应该不为空
	assert.NotEmpty(t, hash)
	// MD5 hash长度应该是32字符
	assert.Equal(t, 32, len(hash))
	// 相同输入应该产生相同的MD5
	hash2 := MD5V(input)
	assert.Equal(t, hash, hash2)

	// 预期的MD5值
	expected := "6f8db599de986fab7a21625b7916589c"
	assert.Equal(t, expected, hash)
}
