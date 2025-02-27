package com.example.redis.cache

import com.example.redis.model.Product // Import the Product class
import org.springframework.data.redis.core.RedisTemplate
import org.springframework.stereotype.Component

@Component
class ProductCache(
    override val redisTemplate: RedisTemplate<String, Any> // Use override here
) : GenericCache<Product>(redisTemplate) {
    override val cacheKey: String = CacheKeys.PRODUCT_PREFIX
}
