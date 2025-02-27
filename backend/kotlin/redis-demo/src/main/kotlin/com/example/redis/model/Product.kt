package com.example.redis.model
import java.io.Serializable

data class Product (
    val id: String,
    val name: String,
    val category: String,
    val qty: Int
) : Serializable
