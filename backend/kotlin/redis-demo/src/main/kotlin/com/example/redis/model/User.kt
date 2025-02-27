package com.example.redis.model
import java.io.Serializable

data class User(
    val id: String,
    val name: String,
    val gender: String,
    val age: Int
) : Serializable