package com.example.redis.repository

import com.example.redis.model.User // Import User
import com.example.redis.cache.CacheKeys // Import CacheKeys
import org.springframework.data.redis.core.RedisTemplate
import org.springframework.stereotype.Repository

import com.example.redis.utils.JsonUtils // Import the utility function

import com.fasterxml.jackson.databind.ObjectMapper

// Assumed the database use Redis for the example, so you can store all data user model into key User
@Repository
class UserRepository(private val redisTemplate: RedisTemplate<String, Any>) {
    
    private val objectMapper = ObjectMapper()  // Create an instance of ObjectMapper

    fun save(user: User) {
        // Print the user object before storing it
        println("Storing user: $user")

        // You can also convert it to JSON string if needed
        val jsonString = objectMapper.writeValueAsString(user)
        println("User as JSON: $jsonString")

        // save as a model
        redisTemplate.opsForHash<String, Any>().put(CacheKeys.USER_CACHE_KEY, user.id, jsonString)
    }

    fun findById(id: String): User? {
        val actualData: Any? = redisTemplate.opsForHash<String, Any>().get(CacheKeys.USER_CACHE_KEY, id)
        return when (actualData) {
            is String -> JsonUtils.jsonToObject<User>(actualData)
            else -> null
        }
    }

    fun delete(id: String) {
        redisTemplate.opsForHash<String, Any>().delete(CacheKeys.USER_CACHE_KEY, id)
    }

    fun findAll(): List<User> {
        val jsonList: List<String> = redisTemplate.opsForHash<String, String>().values(CacheKeys.USER_CACHE_KEY)
        return jsonList.mapNotNull { json ->
            JsonUtils.jsonToObject<User>(json) // Convert JSON string to User
        }
    }
}
