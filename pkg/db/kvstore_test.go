package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const dbPath = "/tmp/test"

func deleteDBAfterTest(t *testing.T) {
	t.Cleanup(func() {
		// delete file
		err := os.RemoveAll(dbPath)
		require.NoError(t, err)
		fmt.Println("Deleted db " + dbPath)
	})
}

func TestSet(t *testing.T) {
	deleteDBAfterTest(t)
	kv, err := NewKVStore(dbPath)
	require.NoError(t, err)
	require.NotNil(t, kv)
	defer kv.Close()

	require.NoError(t, kv.Set([]byte("key"), []byte("value")))
}

func TestGet(t *testing.T) {
	deleteDBAfterTest(t)
	kv, err := NewKVStore(dbPath)
	require.NoError(t, err)
	require.NotNil(t, kv)
	defer kv.Close()

	require.NoError(t, kv.Set([]byte("key"), []byte("value")))

	val, err := kv.Get([]byte("key"))
	require.NoError(t, err)
	require.Equal(t, "value", val)
}

func TestDelete(t *testing.T) {
	deleteDBAfterTest(t)
	kv, err := NewKVStore(dbPath)
	require.NoError(t, err)
	require.NotNil(t, kv)
	defer kv.Close()

	require.NoError(t, kv.Set([]byte("key"), []byte("value")))

	require.NoError(t, kv.Delete([]byte("key")))

	_, err = kv.Get([]byte("key"))
	require.Error(t, err)
}

func TestList(t *testing.T) {
	deleteDBAfterTest(t)
	kv, err := NewKVStore(dbPath)
	require.NoError(t, err)
	require.NotNil(t, kv)
	defer kv.Close()

	require.NoError(t, kv.Set([]byte("key1"), []byte("value1")))
	require.NoError(t, kv.Set([]byte("key2"), []byte("value2")))

	list, err := kv.List()
	require.NoError(t, err)
	require.Equal(t, map[string][]byte{"key1": []byte("value1"), "key2": []byte("value2")}, list)

	fmt.Println(list)
}
