package com.example.redis.cache

object CacheKeys {
    // For Hash-based storage
    const val USER_CACHE_KEY = "UserTbl"
    const val PRODUCT_CACHE_KEY = "ProductTbl"

    const val USER_PREFIX = "user:cache"  // For Value-based storage (if using opsForValue)
    const val PRODUCT_PREFIX = "product:cache"  // For Value-based storage (if using opsForValue)
}
