import asyncio

import asyncclick as click

from script_executors.customers import CreateCustomers


SCRIPTS_EXECUTORS = {
    "CREATE_CUSTOMERS": CreateCustomers
}


@click.command
@click.option("--script")
async def execute_command(script):
    Executor = SCRIPTS_EXECUTORS[script]

    await Executor().process()


if __name__ == "__main__":
    asyncio.run(execute_command())
