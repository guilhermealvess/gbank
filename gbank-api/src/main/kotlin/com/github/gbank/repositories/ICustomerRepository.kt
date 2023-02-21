package com.github.gbank.repositories

import com.github.gbank.models.Customer
import org.springframework.data.repository.CrudRepository
import org.springframework.stereotype.Repository
import java.util.UUID

@Repository
interface ICustomerRepository: CrudRepository<Customer, UUID> {
}