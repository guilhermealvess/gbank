import requests_async as requests
from config import settings


class ApiClient:
    __endpoints = {
        "customers": "/customers"
    }    

    def __init__(self, base_url: str) -> None:
        self.__base_url = base_url

    async def create_customer(self, data: dict):
        endpoint = self.__endpoints["customers"]
        url = f"{self.__base_url}/{endpoint}"
        response = await requests.post(url)

        status_code = response.status_code

        if not 200 <= status_code <= 299:
            raise

        return response.json()


gbank_client = ApiClient(base_url=settings.GBANK_API_URL)
