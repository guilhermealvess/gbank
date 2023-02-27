from pydantic import AnyHttpUrl, BaseSettings


class Settings(BaseSettings):
    class Config:
        env_file = ".env"
        env_file_encoding = "utf-8"

    GBANK_API_URL: str


settings = Settings()
