package com.github.gbank.controllers

import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/ping")
class HealthCheckController {
    @GetMapping
    fun ping() = ResponseEntity.ok("PONG")
}