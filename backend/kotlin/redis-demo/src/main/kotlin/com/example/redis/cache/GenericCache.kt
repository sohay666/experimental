package com.example.redis.cache

import org.springframework.data.redis.core.RedisTemplate
import java.time.Duration

import com.fasterxml.jackson.databind.ObjectMapper


abstract class GenericCache<T>(
    open val redisTemplate: RedisTemplate<String, Any> // Make it open for overriding
) {
    abstract val cacheKey: String // Key defined in subclass
    private val objectMapper = ObjectMapper()  // Create an instance of ObjectMapper

    fun read(cacheKeySuffix: String, expireSeconds: Long = 600, functionToGetActualData: () -> T?): T? {
        val cacheKey = "${this.cacheKey}:$cacheKeySuffix"

        // Check if data exists in cache
        val cachedData: Any? = redisTemplate.opsForValue().get(cacheKey)

        if (cachedData != null) {
            println("[data-redis]cachedData=$cachedData")
            // Safe cast to T
            @Suppress("UNCHECKED_CAST") // Suppress warning for casting
            return cachedData as? T
        }

        // Fetch actual data
        val actualData: T? = functionToGetActualData()
        // Store in Redis with expiration if actualData is not null
        if (actualData != null) {
            println("[data-db]actualData=$actualData")
            val jsonData = objectMapper.writeValueAsString(actualData)
            redisTemplate.opsForValue().set(cacheKey, jsonData, Duration.ofSeconds(expireSeconds))
        }

        return actualData
    }
}
