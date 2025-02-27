package com.example.redis.cache

object CacheKeys {
    const val USER_CACHE_KEY = "UserTbl"  // For Hash-based storage
    const val USER_PREFIX = "user:"  // For Value-based storage (if using opsForValue)
}