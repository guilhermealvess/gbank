package com.github.gbank.controllers

import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.get

@WebMvcTest
class HealthCheckControllerSpec(@Autowired private val mockMvc: MockMvc) {

    @Test
    fun `should ruturn 'PONG'`() {
        mockMvc.get("/ping")
                .andExpect {
                    status { isOk() }
                    content { "PONG" }
                }
    }
}