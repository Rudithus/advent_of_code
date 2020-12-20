load 'integer.rb'

class Operation
  def initialize(memory, mask, address, value)
    @memory = memory
    @address = address.to_36_bit
    @mask = mask
    @value = value
  end

  def apply
    instructions = @mask.chars.map.with_index
    memory_addresses(instructions).each { |address| @memory[address.to_i(2)] = @value }
  end

  def memory_addresses(instructions)
    instructions.reject { |instruction, _| instruction == '0' }
                .each { |instruction, index| @address[index] = instruction }

    bit_array = combination_of_bits(@address.count('X'))

    bit_array.map { |bits| bits.reduce(@address) { |address, bit| address.sub('X', bit.to_s) } }
  end

  def combination_of_bits(n)
    args = []
    (n - 1).times { args.push([0, 1]) }
    [0, 1].product(*args)
  end
end