package com.example.redis.controller

import com.example.redis.model.User
import com.example.redis.service.UserService  

import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/users")
class UserController(private val userService: UserService) {

    @PostMapping
    fun saveUser(@RequestBody user: User) {
        userService.saveUser(user)
    }

    @GetMapping("/{id}")
    fun getUser(@PathVariable id: String): User? {
        return userService.getUser(id)
    }

    @DeleteMapping("/{id}")
    fun deleteUser(@PathVariable id: String) {
        userService.deleteUser(id)
    }
}
