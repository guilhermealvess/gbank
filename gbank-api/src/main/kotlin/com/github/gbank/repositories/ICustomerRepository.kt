package com.github.gbank.repositories

import com.github.gbank.models.Customer
import org.springframework.data.jpa.domain.Specification
import org.springframework.data.jpa.repository.JpaSpecificationExecutor
import org.springframework.data.repository.CrudRepository
import org.springframework.data.repository.PagingAndSortingRepository
import org.springframework.stereotype.Repository
import java.util.UUID

@Repository
interface ICustomerRepository: PagingAndSortingRepository<Customer, UUID>, CrudRepository<Customer, UUID>, JpaSpecificationExecutor<Customer> {
}

fun findByName(name: String) {
    return
}