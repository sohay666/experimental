package com.example.redis.repository

import com.example.redis.model.User // Import User
import com.example.redis.cache.CacheKeys // Import CacheKeys
import org.springframework.data.redis.core.RedisTemplate
import org.springframework.stereotype.Repository


import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.kotlin.KotlinModule
import com.fasterxml.jackson.module.kotlin.readValue


// Asume the database use Redis for the example, so you can store all data user model into key User
@Repository
class UserRepository(private val redisTemplate: RedisTemplate<String, Any>) {
    
    private val objectMapper = ObjectMapper()  // Create an instance of ObjectMapper

    fun jsonToUser(jsonString: String): User? {
        return try {
            objectMapper.registerModule(KotlinModule()).readValue<User>(jsonString)
        } catch (e: Exception) {
            // Handle exception (e.g., log it)
            null
        }
    }


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
        val cachedUser: Any? = redisTemplate.opsForHash<String, Any>().get(CacheKeys.USER_CACHE_KEY, id)
        return when (cachedUser) {
            is String -> jsonToUser(cachedUser) // Convert JSON to User
            is User -> cachedUser // If already a User, return as is
            else -> null
        }
    }

    fun delete(id: String) {
        redisTemplate.opsForHash<String, Any>().delete(CacheKeys.USER_CACHE_KEY, id)
    }

    fun findByName(name: String): List<User> {
        val users = redisTemplate.opsForHash<String, Any>().entries(CacheKeys.USER_CACHE_KEY)
        return users.values.filterIsInstance<User>().filter { it.name == name } // Filtering in application
    }
}
