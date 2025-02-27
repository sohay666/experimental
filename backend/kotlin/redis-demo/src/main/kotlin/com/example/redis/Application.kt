package com.example.redis

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.cache.annotation.EnableCaching

@SpringBootApplication
@EnableCaching  // Enable caching in the application
class Application

fun main(args: Array<String>) {
    runApplication<Application>(*args)
}
