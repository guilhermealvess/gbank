from abc import ABC, abstractclassmethod


class ExecutorScript(ABC):
    @abstractclassmethod
    async def process(self):
        pass
    