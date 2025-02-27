package com.example.redis.service


import com.example.redis.model.User  // Import the User class
import com.example.redis.repository.UserRepository  // Import the UserRepository class
import com.example.redis.cache.UserCache // Import the UserCache class

import org.springframework.stereotype.Service

import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.kotlin.KotlinModule
import com.fasterxml.jackson.module.kotlin.readValue

@Service
class UserService(
    private val userRepository: UserRepository,
    private val userCache: UserCache
) {

    private val objectMapper = ObjectMapper()  // Create an instance of ObjectMapper

    fun jsonToUser(jsonString: String): User? {
        return try {
            objectMapper.registerModule(KotlinModule()).readValue<User>(jsonString)
        } catch (e: Exception) {
            println("[ctl] Error converting JSON to User: ${e.message}")
            null
        }
    }


    fun saveUser(user: User) {
        userRepository.save(user)
    }

    fun getUser(id: String): User? {

        val cachedData: Any? = userCache.read(id) { userRepository.findById(id) }

        return when (cachedData) {
            is String -> jsonToUser(cachedData) // Convert JSON to User
            is User -> cachedData // If already a User, return as is
            else -> null
        }
    }

    fun deleteUser(id: String) {
        userRepository.delete(id)
    }
}
