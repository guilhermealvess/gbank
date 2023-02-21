package com.github.gbank.serives

import com.github.gbank.dto.CustomerDto
import java.util.*

interface ICustomerService {
    fun create(customerDto: CustomerDto): CustomerDto

    fun search(): List<CustomerDto>

    fun findById(id: UUID): CustomerDto

    fun updatePartial(id: UUID, customerDto: CustomerDto): CustomerDto
}