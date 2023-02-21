package com.github.gbank.controllers

import com.github.gbank.dto.CustomerDto
import com.github.gbank.serives.ICustomerService
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.ExceptionHandler
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PatchMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import java.util.UUID
import kotlin.NoSuchElementException

@RestController
@RequestMapping("/api/customers")
class CustomerController(private val service: ICustomerService) {
    @ExceptionHandler(NoSuchElementException::class)
    fun handlerNoSuchElementException(e: NoSuchElementException) =
            ResponseEntity.notFound().build<Void>()

    @PostMapping
    fun create(@RequestBody customerDto: CustomerDto): ResponseEntity<CustomerDto> {
        val customer = service.create(customerDto)

        return ResponseEntity.status(HttpStatus.CREATED).body(customer)
    }

    @GetMapping
    fun search() =
            ResponseEntity.ok(service.search())

    @GetMapping("/{id}")
    fun fetch(@PathVariable("id") id: UUID) =
            ResponseEntity.ok(service.findById(id))

    @PatchMapping("/{id}")
    fun updatePartial(@PathVariable("id") id: UUID, @RequestBody customerDto: CustomerDto): ResponseEntity<CustomerDto> {
        val customer = service.updatePartial(id, customerDto)
        return ResponseEntity.ok(customer)
    }
}