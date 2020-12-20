load 'operation.rb'
@memory = {}

def operations
  file = File.open('input.txt')
  line = file.gets.chomp!
  operations = []
  while line
    mask = line.split(' = ')[1]
    while (line = file.gets) && (match = line.match(/mem\[(\d*)\] = (\d*)/))
      address, value = match.captures
      operations.push(Operation.new(@memory, mask, address.to_i, value.to_i))
    end
  end

  operations
end

operations.each(&:apply)
puts @memory.sum(&:v)
