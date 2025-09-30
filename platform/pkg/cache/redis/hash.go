package platform_redis

import "context"

type HashCache interface {
	// HSet sets the field-value pairs in the hash.
	//
	// If the key does not exist, it will be created.
	// If the field already exists, it will be overwritten.
	//
	// Returns the number of fields that were added.
	// Possible errors: [ErrWrongType], [ErrReadOnly], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed], [ErrTooManyClients].
	HSet(ctx context.Context, key string, values map[string]string) (int64, RedisError)

	// HGet returns the value of the field in the hash.
	//
	// If the key or field is not found, returns empty string and [ErrNotFound].
	// Possible errors: [ErrNotFound], [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HGet(ctx context.Context, key string, field string) (string, RedisError)

	// HGetAll returns all field-value pairs in the hash.
	//
	// If the key does not exist, returns empty map.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HGetAll(ctx context.Context, key string) (map[string]string, RedisError)

	// HMGet returns the values of the fields in the hash.
	//
	// If the field is not found, the value is empty string.
	// Returns slice with same length as fields.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HMGet(ctx context.Context, key string, fields ...string) ([]string, RedisError)

	// HDel deletes the fields from the hash.
	//
	// If the key or field does not exist, it is a noop.
	//
	// Returns the number of fields that were deleted.
	// Possible errors: [ErrWrongType], [ErrReadOnly], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HDel(ctx context.Context, key string, fields ...string) (int64, RedisError)

	// HExists checks if the field exists in the hash.
	//
	// If the key or field does not exist, returns false.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HExists(ctx context.Context, key string, field string) (bool, RedisError)

	// HLen returns the number of fields in the hash.
	//
	// If the key does not exist, returns 0.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HLen(ctx context.Context, key string) (int64, RedisError)

	// HScan iterates over the hash fields using cursor-based pagination.
	//
	// This is the preferred method for iterating over large hashes.
	// Use HKeys/HVals only for small hashes (< 1000 fields).
	//
	// Parameters:
	//   - cursor: iteration cursor. Use 0 to start iteration.
	//   - match: glob-style pattern to filter fields (e.g., "user:*", "item_?"). Use "" for no filtering.
	//   - count: hint for number of fields to return per iteration (default: 10).
	//            Redis may return more or fewer fields. This is a hint, not a limit.
	//
	// Returns:
	//   - fields: slice of alternating field names and values [field1, value1, field2, value2, ...]
	//   - nextCursor: cursor for next iteration. When 0, iteration is complete.
	//
	// Example iteration:
	//   cursor := uint64(0)
	//   for {
	//       fields, nextCursor, err := cache.HScan(ctx, "myHash", cursor, "", 100)
	//       if err != nil { return err }
	//       // Process fields (field at even indices, values at odd)
	//       for i := 0; i < len(fields); i += 2 {
	//           field, value := fields[i], fields[i+1]
	//           // process field and value
	//       }
	//       if nextCursor == 0 { break }
	//       cursor = nextCursor
	//   }
	//
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) (fields []string, nextCursor uint64, err RedisError)

	// HKeys returns all field names in the hash.
	//
	// WARNING: This loads all fields into memory at once.
	// For large hashes (> 1000 fields), use [HScan] instead.
	//
	// If the key does not exist, returns empty slice.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HKeys(ctx context.Context, key string) ([]string, RedisError)

	// HVals returns all values in the hash.
	//
	// WARNING: This loads all values into memory at once.
	// For large hashes (> 1000 fields), use [HScan] instead.
	//
	// If the key does not exist, returns empty slice.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	HVals(ctx context.Context, key string) ([]string, RedisError)
}
