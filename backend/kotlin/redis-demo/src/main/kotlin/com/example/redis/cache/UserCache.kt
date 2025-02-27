package com.example.redis.cache

import com.example.redis.model.User // Import the User class
import org.springframework.data.redis.core.RedisTemplate
import org.springframework.stereotype.Component

@Component
class UserCache(
    override val redisTemplate: RedisTemplate<String, Any> // Use override here
) : GenericCache<User>(redisTemplate) {
    override val cacheKey: String = CacheKeys.USER_PREFIX
}
