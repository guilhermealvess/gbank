import random as rd
import asyncio

from faker import Faker

from script_executors.abstract_executor import ExecutorScript
from clients import gbank_client


class CreateCustomers(ExecutorScript):

    def __init__(self) -> None:
        super().__init__()
        self.__faker = Faker()

    def generate_cpf(self):                                                        
        cpf = [rd.randint(0, 9) for x in range(9)]                              
        for _ in range(2):                                                          
            val = sum([(len(cpf) + 1 - i) * v for i, v in enumerate(cpf)]) % 11      
            cpf.append(11 - val if val > 1 else 0)                                  
        return '%s%s%s%s%s%s%s%s%s%s%s' % tuple(cpf)

    def make_fake_customer(self):
        return {
            "name": self.__faker.name(),
            "birtDate": self.__faker.date_of_birth().isoformat(),
            "documentNumber": self.generate_cpf(),
            "email": self.__faker.email(),
            "phoneNumber": self.__faker.phone_number(),
            "incomeCentsBRL": rd.randint(2000*100, 17000*100),
        }
    
    async def create_customer(self, customer):
        response = await gbank_client.create_customer(customer)
        print(f"Create customer: {response['id']}")

    
    async def process(self):
        fake_customers = [self.make_fake_customer() for _ in range(1000)]
        tasks = [self.create_customer(fake_customer) for fake_customer in fake_customers]
        await asyncio.gather(*tasks)

        return True
