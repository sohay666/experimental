package com.example.redis.repository

import com.example.redis.model.Product // Import Product
import com.example.redis.cache.CacheKeys // Import CacheKeys
import org.springframework.data.redis.core.RedisTemplate
import org.springframework.stereotype.Repository

import com.example.redis.utils.JsonUtils // Import the utility function

import com.fasterxml.jackson.databind.ObjectMapper

// Assumed the database use Redis for the example, so you can store all data product model into key Product
@Repository
class ProductRepository(private val redisTemplate: RedisTemplate<String, Any>) {
    
    private val objectMapper = ObjectMapper()  // Create an instance of ObjectMapper

    fun save(product: Product) {
        // Print the product object before storing it
        println("Storing product: $product")

        // You can also convert it to JSON string if needed
        val jsonString = objectMapper.writeValueAsString(product)
        println("Product as JSON: $jsonString")

        // save as a model
        redisTemplate.opsForHash<String, Any>().put(CacheKeys.PRODUCT_CACHE_KEY, product.id, jsonString)
    }

    fun findById(id: String): Product? {
        val actualData: Any? = redisTemplate.opsForHash<String, Any>().get(CacheKeys.PRODUCT_CACHE_KEY, id)
        return when (actualData) {
            is String -> JsonUtils.jsonToObject<Product>(actualData)
            else -> null
        }
    }

    fun findAll(): List<Product> {
        val jsonList: List<String> = redisTemplate.opsForHash<String, String>().values(CacheKeys.PRODUCT_CACHE_KEY)
        return jsonList.mapNotNull { json ->
            JsonUtils.jsonToObject<Product>(json) // Convert JSON string to Product
        }
    }

    fun delete(id: String) {
        redisTemplate.opsForHash<String, Any>().delete(CacheKeys.PRODUCT_CACHE_KEY, id)
    }
}
