package com.github.gbank.services.impl

import com.github.gbank.dto.CustomerDto
import com.github.gbank.mappers.CustomerMapper
import com.github.gbank.models.Customer
import com.github.gbank.models.Customer_
import com.github.gbank.repositories.ICustomerRepository
import com.github.gbank.services.ICustomerService
import org.springframework.data.domain.Page
import org.springframework.data.domain.Pageable
import org.springframework.data.jpa.domain.Specification
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
        customer.name = customer.name?.uppercase()

        return mapper.toDto(repository.save(customer))
    }

    override fun findById(id: UUID): CustomerDto {
        val customer = repository.findById(id).orElseThrow()
        return mapper.toDto(customer)
    }

    override fun search(pageable: Pageable, name: String?, id: UUID?): Page<CustomerDto> {

        return repository.findAll(filterByName(name).and(filterById(id)), pageable).map { customer -> mapper.toDto(customer) }
    }

    override fun updatePartial(id: UUID, customerDto: CustomerDto): CustomerDto {
        var customer = repository.findById(id).orElseThrow()
        mapper.merge(customerDto, customer)
        customer.updatedAt = LocalDateTime.now()
        customer.name = customer.name?.uppercase()
        customer = repository.save(customer)

        return mapper.toDto(customer)
    }
}

fun filterByName(name: String?): Specification<Customer> {
    return Specification<Customer> {root, query, builder ->
        if(!name.isNullOrBlank()) {
            builder.and(root.get(Customer_.name).`in`(name))
        } else {
            null
        }
    }
}

fun filterById(id: UUID?): Specification<Customer> {
    return Specification<Customer> {root, query, builder ->
        if(id != null) {
            builder.and(root.get(Customer_.id).`in`(id))
        }
        return@Specification null
    }
}
