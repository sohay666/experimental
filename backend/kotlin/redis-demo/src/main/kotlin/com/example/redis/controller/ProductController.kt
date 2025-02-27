package com.example.redis.controller

import com.example.redis.model.Product
import com.example.redis.service.ProductService  

import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/products")
class ProductController(private val productService: ProductService) {

    @PostMapping
    fun saveProduct(@RequestBody product: Product) {
        productService.saveProduct(product)
    }

    @GetMapping("/{id}")
    fun getProduct(@PathVariable id: String): Product? {
        return productService.getProduct(id)
    }

    @GetMapping
    fun getAllProducts(): List<Product> {
        return productService.getAllProducts()
    }

    @DeleteMapping("/{id}")
    fun deleteProduct(@PathVariable id: String) {
        productService.deleteProduct(id)
    }
}
