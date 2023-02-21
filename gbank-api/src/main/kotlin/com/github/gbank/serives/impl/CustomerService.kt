package com.github.gbank.serives.impl

import com.github.gbank.dto.CustomerDto
import com.github.gbank.mappers.CustomerMapper
import com.github.gbank.repositories.ICustomerRepository
import com.github.gbank.serives.ICustomerService
import org.springframework.data.domain.Page
import org.springframework.data.domain.Pageable
import org.springframework.stereotype.Service
import java.time.LocalDateTime
import java.util.*

@Service
class CustomerService(private val repository: ICustomerRepository, private val mapper: CustomerMapper): ICustomerService {
    override fun create(customerDto: CustomerDto): CustomerDto {
        val customer = mapper.toModel(customerDto)
        val now = LocalDateTime.now()
        customer.createdAt = now
        customer.updatedAt = now
        customer.name = customer.name?.toUpperCase()

        return mapper.toDto(repository.save(customer))
    }

    override fun findById(id: UUID): CustomerDto {
        val customer = repository.findById(id).orElseThrow()
        return mapper.toDto(customer)
    }

    override fun search(pageable: Pageable): Page<CustomerDto> {
        return repository.findAll(pageable).map { customer -> mapper.toDto(customer) }
    }

    override fun updatePartial(id: UUID, customerDto: CustomerDto): CustomerDto {
        var customer = repository.findById(id).orElseThrow()
        mapper.merge(customerDto, customer)
        customer.updatedAt = LocalDateTime.now()
        customer.name = customer.name?.toUpperCase()
        customer = repository.save(customer)

        return mapper.toDto(customer)
    }
}