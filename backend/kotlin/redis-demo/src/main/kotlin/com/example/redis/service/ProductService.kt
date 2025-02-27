package com.example.redis.service


import com.example.redis.model.Product  // Import the User class
import com.example.redis.repository.ProductRepository  // Import the UserRepository class
import com.example.redis.cache.ProductCache // Import the ProductCache class

import com.example.redis.utils.JsonUtils // Import the utility function

import org.springframework.stereotype.Service

@Service
class ProductService(
    private val productRepository: ProductRepository,
    private val productCache: ProductCache
) {
    fun saveProduct(product: Product) {
        productRepository.save(product)
    }

    fun getProduct(id: String): Product? {

        val cachedData: Any? = productCache.read(id) { productRepository.findById(id) }

        return when (cachedData) {
            is String -> JsonUtils.jsonToObject<Product>(cachedData)
            is Product -> cachedData // If already a Product, return as is
            else -> null
        }
    }

    fun deleteProduct(id: String) {
        productRepository.delete(id)
    }

    fun getAllProducts(): List<Product> {
        return productRepository.findAll()
    }
}
