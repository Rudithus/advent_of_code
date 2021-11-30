def input
  File.read('input2.txt').split(',').map(&:to_i)
end
@hash = Hash.new(0)

numbers = input.to_a
last_number = numbers.pop

numbers.each_with_index.each { |number, index| @hash[number] = index + 1 }

final_round = 30_000_000
starting_round = numbers.count + 1

(starting_round..final_round - 1).each do |turn|
  last_index = @hash[last_number]
  @hash[last_number] = turn
  last_number = last_index.zero? ? 0 : turn - last_index
end

p last_number
