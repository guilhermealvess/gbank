package com.github.gbank.services

import com.github.gbank.dto.CustomerDto
import org.springframework.data.domain.Page
import org.springframework.data.domain.Pageable
import java.util.*

interface ICustomerService {
    fun create(customerDto: CustomerDto): CustomerDto

    fun search(pageable: Pageable, name: String?, id: UUID?): Page<CustomerDto>

    fun findById(id: UUID): CustomerDto

    fun updatePartial(id: UUID, customerDto: CustomerDto): CustomerDto
}