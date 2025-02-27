package com.example.redis.utils

import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.kotlin.KotlinModule
import com.fasterxml.jackson.module.kotlin.readValue

object JsonUtils {

    val objectMapper = ObjectMapper().registerModule(KotlinModule())  // Create an instance of ObjectMapper

    /**
     * Converts a JSON string to an object of the specified type.
     *
     * @param jsonString The JSON string to convert.
     * @param objectMapper The ObjectMapper instance to use for conversion.
     * @return An object of type T, or null if conversion fails.
     */
    inline fun <reified T> jsonToObject(jsonString: String, objectMapper: ObjectMapper = JsonUtils.objectMapper): T? {
        return try {
            objectMapper.readValue<T>(jsonString)
        } catch (e: Exception) {
            println("[ctl] Error converting JSON to ${T::class.simpleName}: ${e.message}")
            null
        }
    }
}