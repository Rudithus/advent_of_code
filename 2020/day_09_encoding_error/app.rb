require 'set'
load 'preamble.rb'

def input
  File.open('input.txt').readlines.map(&:chomp).map(&:to_i)
end

def invalid_number(preamble, numbers)
  numbers.each do |n|
    return n unless preamble.include?(n)

    preamble.add(n)
  end
end

def find_sum(numbers, sum)
  upper_limit = numbers.count - 1
  (0..upper_limit).each do |i|
    (0..upper_limit).each do |j|
      slice = numbers.slice(i, j).to_a
      slice_sum = slice.sum
      return slice if slice_sum == sum
      break if slice_sum > sum
    end
  end
end

numbers = input.to_a
preamble = Preamble.new(numbers.shift(25))

invalid_number = invalid_number(preamble, numbers)

arr = find_sum(input.to_a, invalid_number)
p arr.min + arr.max
