package com.example.redis.service


import com.example.redis.model.User  // Import the User class
import com.example.redis.repository.UserRepository  // Import the UserRepository class
import com.example.redis.cache.UserCache // Import the UserCache class

import com.example.redis.utils.JsonUtils // Import the utility function

import org.springframework.stereotype.Service

@Service
class UserService(
    private val userRepository: UserRepository,
    private val userCache: UserCache
) {

    fun saveUser(user: User) {
        userRepository.save(user)
    }

    fun getUser(id: String): User? {
        
        val cachedData: Any? = userCache.read(id) { userRepository.findById(id) }

        return when (cachedData) {
            is String -> JsonUtils.jsonToObject<User>(cachedData)
            is User -> cachedData // If already a User, return as is
            else -> null
        }
    }

    fun deleteUser(id: String) {
        userRepository.delete(id)
    }

    fun getAllUsers(): List<User> {
        return userRepository.findAll()
    }    
}
