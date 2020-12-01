from collections import defaultdict

import trio


class AsyncComputer:
    def __init__(self, program):
        self.position = 0
        self.relative_base = 0
        self.parameters_mode = None
        self.program = program.copy()

    def set_input_output(self, buffer_max=10):
        input_write, input_read = trio.open_memory_channel(buffer_max)
        output_write, output_read = trio.open_memory_channel(buffer_max)
        
        self.input, self.output = input_read, output_write
        
        return input_write, output_read
        
    @staticmethod
    def log(*args):
        if False:
            print(*args)

    def get_param_value(self, param):
        param_address = self.get_param_address(param)
        
        try:
            value = self.program[param_address]
        except IndexError:
            self.extend(param_address)
            value = self.program[param_address]
            
        return value

    def get_param_address(self, param):
        mode = self.parameters_mode[param]

        if mode == 0:  # position mode
            return self.program[self.position + param]
        elif mode == 1:  # immediate mode
            return self.position + param
        elif mode == 2:  # relative mode:
            param_position = self.position + param
            return self.relative_base + self.program[param_position]
        else:
            raise ValueError("Unknown paremeter mode")

    def read_instruction(self):
        # read instruction and split it into opcode and parameters mode
        try:
            instruction = self.program[self.position]
        except IndexError:
            # end of program
            return None

        opcode = instruction % 100
        self.parameters_mode = defaultdict(lambda: 0)

        value = instruction // 100
        i = 1
        while value > 0:
            self.parameters_mode[i] = value % 10
            value //= 10
            i += 1

        self.log(self.position, '>', instruction, opcode, dict(self.parameters_mode))
        return opcode
    
    def write(self, address, value):
        try:
            self.program[address] = value
        except IndexError:
            self.extend(address)
            self.program[address] = value
            
    def extend(self, address):
        self.program.extend([0] * (1 + address - len(self.program)))

    async def compute_close(self):
        """Run the programm, but close the output when terminated."""
        async with self.output:
            await self.compute()
            
        
    async def compute(self):
        """Compute the final state of a program, and return it.

        `program` is a list of int.
        `read_input` is a method returning a value each time it is called.
        `write_output` is a method taking a value as parameter.
        """
        # init state
        self.position = 0

        # run program
        while True:
            opcode = self.read_instruction()
            if opcode is None:
                # end of program
                log("Stop: EOF")
                return

            # execute opcode
            if opcode == 1:
                # add x y
                x = self.get_param_value(1)
                y = self.get_param_value(2)
                result_addr = self.get_param_address(3)

                result = x + y
                self.write(result_addr, result)
                self.log("add", x, y, result_addr, result)

                step = 4
            elif opcode == 2:
                # mult x y
                x = self.get_param_value(1)
                y = self.get_param_value(2)
                result_addr = self.get_param_address(3)

                result = x * y
                self.write(result_addr, result)
                self.log("mult", x, y, result_addr, result)

                step = 4
            elif opcode == 3:
                # read x
                result_addr = self.get_param_address(1)

                value = await self.input.receive()
                self.write(result_addr, value)
                self.log("read", value, result_addr)

                step = 2
            elif opcode == 4:
                # write x
                value = self.get_param_value(1)
                await self.output.send(value)
                self.log("write", value)

                step = 2
            elif opcode == 5:
                # jump-if-true x y
                value = self.get_param_value(1)
                jump_addr = self.get_param_value(2)

                self.log("jump-if-true", value, jump_addr)

                if value:
                    self.position = jump_addr
                    step = 0
                else:
                    step = 3
            elif opcode == 6:
                # jump-if-false x y
                value = self.get_param_value(1)
                jump_addr = self.get_param_value(2)

                self.log("jump-if-false", value, jump_addr)

                if not value:
                    self.position = jump_addr
                    step = 0
                else:
                    step = 3
            elif opcode == 7:
                # lt x y z
                x = self.get_param_value(1)
                y = self.get_param_value(2)
                result_addr = self.get_param_address(3)

                self.log("lt", x, y, result_addr)

                if x < y:
                    self.write(result_addr, 1)
                else:
                    self.write(result_addr, 0)

                step = 4
            elif opcode == 8:
                # eq x y z
                x = self.get_param_value(1)
                y = self.get_param_value(2)
                result_addr = self.get_param_address(3)

                self.log("eq", x, y, result_addr)

                if x == y:
                    self.write(result_addr, 1)
                else:
                    self.write(result_addr, 0)

                step = 4
            elif opcode == 9:
                # relative_base_add x
                x = self.get_param_value(1)
                self.relative_base += x
                
                self.log("relative_base_add", x)
                
                step = 2
            elif opcode == 99:
                # end of program
                self.log("stop")
                break
            else:
                # unknown instruction
                self.log("unknown opcode")
                break

            self.position += step

        return self.program


class Computer:
    def __init__(self, program, io):
        self.io = io
        self.parameters_mode = None
        self.position = 0
        self.program = program.copy()
        self.relative_base = 0

    @staticmethod
    def log(*args):
        if False:
            print(*args)

    def get_param_value(self, param):
        param_address = self.get_param_address(param)
        
        try:
            value = self.program[param_address]
        except IndexError:
            self.extend(param_address)
            value = self.program[param_address]
            
        return value

    def get_param_address(self, param):
        mode = self.parameters_mode[param]

        if mode == 0:  # position mode
            return self.program[self.position + param]
        elif mode == 1:  # immediate mode
            return self.position + param
        elif mode == 2:  # relative mode:
            param_position = self.position + param
            return self.relative_base + self.program[param_position]
        else:
            raise ValueError("Unknown paremeter mode")

    def read_instruction(self):
        # read instruction and split it into opcode and parameters mode
        try:
            instruction = self.program[self.position]
        except IndexError:
            # end of program
            return None

        opcode = instruction % 100
        self.parameters_mode = defaultdict(lambda: 0)

        value = instruction // 100
        i = 1
        while value > 0:
            self.parameters_mode[i] = value % 10
            value //= 10
            i += 1

        self.log(self.position, '>', instruction, opcode, dict(self.parameters_mode))
        return opcode
    
    def write(self, address, value):
        try:
            self.program[address] = value
        except IndexError:
            self.extend(address)
            self.program[address] = value
            
    def extend(self, address):
        self.program.extend([0] * (1 + address - len(self.program)))

    def compute(self):
        """Compute the final state of a program, and return it.

        `program` is a list of int.
        `read_input` is a method returning a value each time it is called.
        `write_output` is a method taking a value as parameter.
        """
        # init state
        self.position = 0

        # run program
        while True:
            opcode = self.read_instruction()
            if opcode is None:
                # end of program
                log("Stop: EOF")
                return

            # execute opcode
            if opcode == 1:
                # add x y
                x = self.get_param_value(1)
                y = self.get_param_value(2)
                result_addr = self.get_param_address(3)

                result = x + y
                self.write(result_addr, result)
                self.log("add", x, y, result_addr, result)

                step = 4
            elif opcode == 2:
                # mult x y
                x = self.get_param_value(1)
                y = self.get_param_value(2)
                result_addr = self.get_param_address(3)

                result = x * y
                self.write(result_addr, result)
                self.log("mult", x, y, result_addr, result)

                step = 4
            elif opcode == 3:
                # read x
                result_addr = self.get_param_address(1)

                value = self.io.input()
                self.write(result_addr, value)
                self.log("read", value, result_addr)

                step = 2
            elif opcode == 4:
                # write x
                value = self.get_param_value(1)
                self.io.output(value)
                self.log("write", value)

                step = 2
            elif opcode == 5:
                # jump-if-true x y
                value = self.get_param_value(1)
                jump_addr = self.get_param_value(2)

                self.log("jump-if-true", value, jump_addr)

                if value:
                    self.position = jump_addr
                    step = 0
                else:
                    step = 3
            elif opcode == 6:
                # jump-if-false x y
                value = self.get_param_value(1)
                jump_addr = self.get_param_value(2)

                self.log("jump-if-false", value, jump_addr)

                if not value:
                    self.position = jump_addr
                    step = 0
                else:
                    step = 3
            elif opcode == 7:
                # lt x y z
                x = self.get_param_value(1)
                y = self.get_param_value(2)
                result_addr = self.get_param_address(3)

                self.log("lt", x, y, result_addr)

                if x < y:
                    self.write(result_addr, 1)
                else:
                    self.write(result_addr, 0)

                step = 4
            elif opcode == 8:
                # eq x y z
                x = self.get_param_value(1)
                y = self.get_param_value(2)
                result_addr = self.get_param_address(3)

                self.log("eq", x, y, result_addr)

                if x == y:
                    self.write(result_addr, 1)
                else:
                    self.write(result_addr, 0)

                step = 4
            elif opcode == 9:
                # relative_base_add x
                x = self.get_param_value(1)
                self.relative_base += x
                
                self.log("relative_base_add", x)
                
                step = 2
            elif opcode == 99:
                # end of program
                self.log("stop")
                break
            else:
                # unknown instruction
                self.log("unknown opcode")
                break

            self.position += step

        return self.program
